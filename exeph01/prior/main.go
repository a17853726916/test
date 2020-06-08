package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"time"
)

//定义挖矿难度，该难度是全局的为了方便计算
// 实际的难度是随时改变的
const targetBit = 20
const dbName = "blockchain.db" //数据库的名字
const bkName = "blocks"        //桶的名字

// 在区块的结构中加入新的属性Nonce(随机数)，用于生成工作量证明的哈希
// Nonce 用来保存一个随机值，poW算法中Nonce和区块其他信息一起进行Hash计算，使计算符合一定的条件，
// 区块结构
type Block struct {
	Index         int64  //区块位置的索引
	TimeStamp     int64  //时间戳
	Data          []byte //存放的数据
	PrevBlockHash []byte //前一区块的哈希值
	Hash          []byte //当前区块的哈希值
	Nonce         int64  //随机数，生成工作量证明的哈希
}

// 定义poW
type ProofOfWork struct {
	block  *Block
	target *big.Int //定义为一个大整数
}

// 先使用数组保证区块的结构
type Blockchain struct {
	blocks []*Block
}

// 构建创世区块，即头区块
func NewGenesisBlock() *Block {
	return NewBlock(0, "first,block", []byte{})
}

// 将创世区块添加进去，用来引导之后的区块
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// 根据上一个区块，添加到新的区块链中去
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)                 //将 big.Int 初始化为 1
	target.Lsh(target, uint(256-targetBit)) //左移 256 - 1 位
	// fmt.Printf("%x\n", target)
	pow := &ProofOfWork{b, target}

	return pow
}

// 数据准备
// 使用IntToHex函数将int型数据转行为16进制。
func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)                         //创建一个缓冲池
	err := binary.Write(buffer, binary.BigEndian, data) //BigEndian :大端字节序
	if err != nil {
		log.Panicf("int to []byte fail err = %v\n", err)
	}
	return buffer.Bytes()
}

// poW算法中Nonce和区块其他信息一起进行Hash计算，使计算符合一定的条件。
// 用prepareData函数将Nonce与区块其他信息合并
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Index),
			IntToHex(pow.block.TimeStamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// 实现poW的核心代码
func (pow *ProofOfWork) Run() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	var nonce int64 = 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for {
		dataBytes := pow.prepareData(nonce) //获取准备数据
		hash = sha256.Sum256(dataBytes)     //对数据进行hash运算
		hashInt.SetBytes(hash[:])
		fmt.Printf("hash: \r%x", hash)
		if pow.target.Cmp(&hashInt) == 1 { //对比hash值
			break
		}
		nonce++ //充当计数器，同时在循环结束后也是符合要求的值

	}
	fmt.Printf("\n碰撞次数 %d\n", nonce)
	return int64(nonce), hash[:]
}

// 创建区块
// 改造后新区块的定义添加了Nonce ，所以我们要对Nonce也进行赋值，
// 并且我们要通过poW算法来重新生成区块。
func NewBlock(index int64, data string, prevBlockHash []byte) *Block {

	//定义一个新的区块实例
	block := &Block{
		index,
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
		0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// 验证区块
// 区块的数据有任意改动，哪怕是一个字节的改动，Hash值都会随着变化。
// 因此在对于新生成的区块，我们非常有必要对它重新计算hash，验证是否合法
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

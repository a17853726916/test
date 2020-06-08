package main

import (
	"fmt"
	"strings"
)

// 在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，
// 每一列都按照从上到下递增的顺序排序。
// 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
func Find(target int, array [3][3]int) bool {

	// 二维数组循环 暴力解决法
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if target == array[i][j] {
				return true
			}
		}
	}
	return false

}
func Find2(target int, array [3][3]int) bool {
	//右上或者左下
	// 右上的的呢个元素为改行的最大值该列的最小值  如果要查询的元素的比这个元素小则从该行进行查找
	// 如果元素比这个元素大 就从下一行最后一列进行寻找
	// 同理左下也是也个意思  时间复杂度会比暴力法低一些
	// for i := 0; i < 3; i++ {
	// 	if target == array[i][2] {
	// 		return true
	// 	} else if target < array[i][2] {
	// 		for j := 0; j < 3; j++ {
	// 			if target == array[0][j] {
	// 				return true
	// 			}
	// 		}
	// 	} else {
	// 		for j := 0; j < 3; j++ {
	// 			if target == array[0][j] {
	// 				return true
	// 			}
	// 		}
	// 	}
	// }
	// return false

	// 简洁写法
	i := 0 //行
	j := 2 // 列
	if (i < 3) && (j >= 0) {
		if array[i][j] < target {
			i++
		} else if array[i][j] > target {
			j--
		} else {
			return true
		}
	}
	return false
}

// 2.请实现一个函数，将一个字符串中的每个空格替换成“%20”。
// 例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy
// 思路：
// 		go语言中有一个操作字符串的函数strings.Fields(s string ) []string
// 		该函数的作用是按照空格切割
// strings.Fields(s string) []string
// 返回使用空格分割的字符串 s，结果为切片。
// strings.Fields("Han Zhong Kang")   利用该函数 切割
// 返回 []string, ["Han", "Zhong", "Kang"]

// strings.Join(a []string, sep string) string
// 使用分隔符 sep 连接字符串切片 a。
// ss := []string{"Go", "Hank", "Python", "PHP"}
// strings.Join(ss, "-")
// 返回 "Go-Hank-Python-PHP"   使用该函数进行拼接
func replaceSpace(s string) string {
	s1 := strings.Fields(s)
	s2 := strings.Join(s1, "%20")

	return s2
}
func main() {
	// a := [3][3]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}

	// b := 3
	// c := 9
	// fmt.Println(Find2(b, a))
	// fmt.Println(Find2(c, a))
	a := "We are Happy"
	fmt.Println(replaceSpace(a))

}

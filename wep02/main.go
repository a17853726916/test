package main

import "fmt"

// 大家都知道斐波那契数列，现在要求输入一个整数n，
// 请你输出斐波那契数列的第n项（从0开始，第0项为0）。 n<=39
func Fibonacci(n int) int {

	//普通方法
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	var s []int
	s = append(s, 1)
	s = append(s, 1)
	var i int = 2
	for i = 2; i <= n; i++ {
		s = append(s, s[i-1]+s[i-2])
	}

	return s[i-1]

	//递归
	var a int = 1
	var b int = 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// 一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法
// 思路 ：递归的思想 就是当青蛙最后一次跳的时候可能是一步也可能是两步 f（n-1） f（n-2）
// 同理执勤最后一次的前一次也是同理的，所以就分解了这样的小问题 最后将两者加起来就是一共的方法

// 数学归纳法
// 第一次跳一步 所以只有1种 第二次 1或2 两种 第三次111 12 21 三种 第四次 1111 112 121 211 22 5种
// .....可以看出就是斐波那契数列
func junmpFloor(number int) int {
	// 递归
	// if number == 1 {
	// 	return 1
	// }
	// if number == 2 {
	// 	return 2
	// }
	// if number == 3 {
	// 	return 3
	// }
	// return junmpFloor(number-1) + junmpFloor(number-2)

	// 利用斐波那契
	var a int = 1
	var b int = 1
	for i := 0; i < number; i++ {
		a, b = b, a+b
	}
	return a

}

// 一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台
// f（n) = f(n-1)+f(n-2)+f(n-3)+...+f(1)
// f(n-1) = f(n-2)+f(n-3)+...+f(1) 两式做差
// f(n) = 2f(n-1).....
func junmpFloor2(number int) int {
	var a int = 1
	for i := 2; i < number+1; i++ {
		a = 2 * a
	}
	return a
}

// 我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。
// 请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
func rectCover(number int) int {
	// 利用斐波那契
	var c int = 1
	var d int = 1
	for i := 0; i < number; i++ {
		c, d = c, c+d
	}
	return c
}
func main() {
	// var a int
	// fmt.Println("请输入青蛙需要挑的台阶数")
	// fmt.Scanln(&a)
	// fmt.Println("青蛙需要跳的步数是：", junmpFloor(a))
	// fmt.Println("青蛙需要跳的步数是：", junmpFloor2(a))
	// fmt.Println("可以用的方法有", rectCover(a))
	fmt.Println(Fibonacci(5))

}

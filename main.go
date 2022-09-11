package main

import "fmt"

func main() {

	n := 1
	fmt.Println(climbStairs(n))

}

func climbStairs(n int) int {

	pre := 0
	res := 1
	for i := 1; i <= n; i++ {
		res, pre = pre+res, res
	}

	return res
}

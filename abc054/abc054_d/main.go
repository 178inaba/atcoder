package main

import "fmt"

// TODO: incomplete.
func main() {
	var n, ma, mb int
	fmt.Scanf("%d %d %d", &n, &ma, &mb)

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &a[i], &b[i], &c[i])
	}

	var min int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			solutionA := a[i] + a[j]
			solutionB := b[i] + b[j]
			if solutionA/ma == solutionB/mb {
				money := c[i] + c[j]
				if min == 0 || money < min {
					min = money
				}
			}
		}
	}

	if min == 0 {
		min = -1
	}

	fmt.Println(min)
}

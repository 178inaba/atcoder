package main

import "fmt"

func main() {
	var n, ma, mb int
	fmt.Scanf("%d %d %d", &n, &ma, &mb)

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	var abMax, cMax int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &a[i], &b[i], &c[i])

		if a[i] > abMax {
			abMax = a[i]
		}

		if b[i] > abMax {
			abMax = b[i]
		}

		if c[i] > cMax {
			cMax = c[i]
		}
	}

	mixMax := n*abMax + 1
	priceMax := n*cMax + 1

	// Initialize.
	dp := make([][][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([][]int, mixMax)
		for j := 0; j < mixMax; j++ {
			dp[i][j] = make([]int, mixMax)
			for k := 0; k < mixMax; k++ {
				dp[i][j][k] = priceMax
			}
		}
	}

	dp[0][0][0] = 0

	for i := 0; i < n; i++ {
		for ca := 0; ca < mixMax; ca++ {
			for cb := 0; cb < mixMax; cb++ {
				if dp[i][ca][cb] == priceMax {
					continue
				}

				dp[i+1][ca][cb] = selectMin(dp[i+1][ca][cb], dp[i][ca][cb])
				dp[i+1][ca+a[i]][cb+b[i]] = selectMin(dp[i+1][ca+a[i]][cb+b[i]], dp[i][ca][cb]+c[i])
			}
		}
	}

	ans := priceMax
	for ca := 1; ca < mixMax; ca++ {
		for cb := 1; cb < mixMax; cb++ {
			if ca*mb == cb*ma {
				ans = selectMin(ans, dp[n][ca][cb])
			}
		}
	}

	if ans == priceMax {
		ans = -1
	}

	fmt.Println(ans)
}

func selectMin(x, y int) int {
	if x < y {
		return x
	}

	return y
}

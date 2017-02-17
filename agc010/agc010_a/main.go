package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Input. ----------

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	return i
}

// ---------- Input.

func main() {
	N := nextInt()
	if N == 1 {
		fmt.Println("YES")
		return
	}

	sum := 0
	for i := 0; i < N; i++ {
		sum += nextInt()
	}

	ans := ""
	if sum%2 == 0 {
		ans = "YES"
	} else {
		ans = "NO"
	}

	fmt.Println(ans)
}

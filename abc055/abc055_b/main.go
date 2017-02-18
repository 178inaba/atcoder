package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	N := uint64(nextInt())
	power := uint64(1)
	max := uint64(math.Pow10(9)) + 7
	for i := uint64(1); i <= N; i++ {
		power = power * i % max
	}

	fmt.Println(power)
}

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

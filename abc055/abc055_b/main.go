package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	power := 1
	max := int(math.Pow10(9)) + 7
	for i := 1; i <= N; i++ {
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

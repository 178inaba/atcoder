package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	_ = nextInt()
	S := nextString()
	var lCnt, rCnt int
	ans := []byte(S)
	for _, s := range S {
		if s == '(' {
			lCnt++
		} else {
			rCnt++
		}
		diffCnt := lCnt - rCnt
		if diffCnt < 0 {
			ans = append([]byte{'('}, ans...)
			lCnt++
		}
	}
	diffCnt := lCnt - rCnt
	if diffCnt > 0 {
		for i := 0; i < diffCnt; i++ {
			ans = append(ans, []byte{')'}...)
		}
	}
	fmt.Println(string(ans))
}

// Input. ----------

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
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

func nextString() string {
	sc.Scan()
	if err := sc.Err(); err != nil {
		panic(err)
	}

	return sc.Text()
}

// ---------- Input.

// Util. ----------

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// ---------- Util.

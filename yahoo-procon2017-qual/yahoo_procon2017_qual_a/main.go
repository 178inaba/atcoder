package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	str := nextString()

	base := []rune{'y', 'a', 'h', 'o', 'o'}
	for _, r := range str {
		for i, br := range base {
			if br == r {
				base = append(base[:i], base[i+1:]...)
				break
			}
		}
	}

	if len(base) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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
	if x < y {
		return x
	}

	return y
}

// ---------- Util.

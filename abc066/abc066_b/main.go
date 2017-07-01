package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	S := nextString()
	S = S[:len(S)-2]
	for {
		S1 := S[:len(S)/2]
		S2 := S[len(S)/2:]
		isSame := true
		for i := 0; i < len(S1); i++ {
			if S1[i] != S2[i] {
				isSame = false
				break
			}
		}
		if isSame {
			break
		}
		S = S[:len(S)-2]
	}
	fmt.Println(len(S))
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

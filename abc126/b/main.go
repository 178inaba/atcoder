package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	S := nextString()
	before, err := strconv.Atoi(S[:2])
	if err != nil {
		log.Fatal(err)
	}

	after, err := strconv.Atoi(S[2:])
	if err != nil {
		log.Fatal(err)
	}

	var okBefore, okAfter bool
	if 1 <= before && before <= 12 {
		okBefore = true
	}
	if 1 <= after && after <= 12 {
		okAfter = true
	}

	if okBefore && okAfter {
		fmt.Println("AMBIGUOUS")
		return
	} else if okBefore {
		fmt.Println("MMYY")
		return
	} else if okAfter {
		fmt.Println("YYMM")
		return
	}

	fmt.Println("NA")
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

func nextFloat() float64 {
	sc.Scan()
	f64, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}

	return f64
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

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// ---------- Util.

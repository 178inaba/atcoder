package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	n := nextInt()
	S := make([]string, 0, n-1)
	var base string
	lenMin := 50
	for i := 0; i < n; i++ {
		s := nextString()
		if len(s) < lenMin {
			lenMin = len(s)
			base = s
		} else if len(s) == lenMin && s < base {
			base = s
		} else {
			S = append(S, s)
		}
	}

	okRunes := []rune(base)
	for ir := 0; ir < len(okRunes); ir++ {
		for i := 0; i < len(S); i++ {
			exist := false
			for j := 0; j < len(S[i]); j++ {
				if S[i][j] == byte(okRunes[ir]) {
					exist = true
					S[i] = string(append([]byte(S[i])[:j], S[i][j+1:]...))
				}
			}

			if !exist {
				okRunes = append(okRunes[:ir], okRunes[ir+1:]...)
			}
		}
	}

	sort.Sort(runeSlice(okRunes))
	fmt.Println(string(okRunes))
}

type runeSlice []rune

func (p runeSlice) Len() int           { return len(p) }
func (p runeSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p runeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

// ---------- Util.

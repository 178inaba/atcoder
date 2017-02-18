package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	N := nextInt()
	s := nextString()

	s += string(s[0])
	starts := [][]string{{"S", "S"}, {"S", "W"}, {"W", "S"}, {"W", "W"}}
	for _, start := range starts {
		animals := make([]string, N+2)
		animals[0] = start[0]
		animals[1] = start[1]
		for i := 1; i <= N; i++ {
			if animals[i] == "S" {
				if s[i] == 'o' {
					if animals[i-1] == "S" {
						animals[i+1] = "S"
					} else {
						animals[i+1] = "W"
					}
				} else {
					if animals[i-1] == "S" {
						animals[i+1] = "W"
					} else {
						animals[i+1] = "S"
					}
				}
			} else {
				if s[i] == 'o' {
					if animals[i-1] == "S" {
						animals[i+1] = "W"
					} else {
						animals[i+1] = "S"
					}
				} else {
					if animals[i-1] == "S" {
						animals[i+1] = "S"
					} else {
						animals[i+1] = "W"
					}
				}
			}
		}

		if animals[0] == animals[N] && animals[1] == animals[N+1] {
			fmt.Println(strings.Join(animals[:N], ""))
			return
		}
	}

	fmt.Println(-1)
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

func nextString() string {
	sc.Scan()

	return sc.Text()
}

// ---------- Input.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	M := nextInt()

	ans := 0
	for {
		for {
			if N == 0 || M < 2 {
				break
			}

			ans++
			N--
			M -= 2
		}

		if M < 4 {
			break
		}

		if N == 0 {
			N++
			M -= 2
		}
	}

	fmt.Println(ans)
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

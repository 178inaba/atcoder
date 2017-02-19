package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// N is ...
const N = 50

func main() {
	blocks := make([][]byte, N+2)
	blocks[0] = make([]byte, N+2)
	blocks[N+1] = make([]byte, N+2)
	for i := 1; i <= N; i++ {
		blocks[i] = append([]byte{' '}, []byte(nextString())...)
		blocks[i] = append(blocks[i], ' ')
	}

	_, maxBlocks := edit(blocks, 2, 1)
	for _, block := range maxBlocks {
		fmt.Println(string(block))
	}
}

func copyBlocks(base [][]byte) [][]byte {
	cb := make([][]byte, N+2)
	for i := 0; i < N+2; i++ {
		cb[i] = make([]byte, N+2)
		copy(cb[i], base[i])
	}

	return cb
}

func edit(blocks [][]byte, i, j int) (int, [][]byte) {
	maxPoint := 0
	maxBlocks := copyBlocks(blocks)
	for k := i; k > 1; k-- {
		for l := j; l <= N; l++ {
			if blocks[k][l] == '.' {
				plus := copyBlocks(blocks)
				minus := copyBlocks(blocks)

				plus[k][l] = '+'
				minus[k][l] = '-'

				if p := result(copyBlocks(blocks)); p > maxPoint {
					maxPoint = p
					maxBlocks = blocks
				}

				if p := result(copyBlocks(plus)); p > maxPoint {
					maxPoint = p
					maxBlocks = plus
				}

				if p := result(copyBlocks(minus)); p > maxPoint {
					maxPoint = p
					maxBlocks = minus
				}

				if l == N {
					if k-1 == 1 {
						return maxPoint, maxBlocks
					}

					if p, b := edit(blocks, k-1, 1); p > maxPoint {
						maxPoint = p
						maxBlocks = b
					}
				} else {
					if p, b := edit(blocks, k, l+1); p > maxPoint {
						maxPoint = p
						maxBlocks = b
					}
				}

				if p, b := edit(plus, k, l+1); p > maxPoint {
					maxPoint = p
					maxBlocks = b
				}

				if p, b := edit(minus, k, l+1); p > maxPoint {
					maxPoint = p
					maxBlocks = b
				}
			}
		}
	}

	return maxPoint, maxBlocks
}

func result(blocks [][]byte) int {
	for i := N; i >= 1; i-- {
		for j, b := range blocks[i] {
			if b == '.' {
				drop(blocks, i, j)
			}
		}
	}

	var oMax, xMax int
	for i := 1; i <= N; i++ {
		for j, b := range blocks[i] {
			if b == 'o' {
				// Search up, down, right, left.
				cm := checkMap()
				p := check(blocks, i, j, cm)
				if p > oMax {
					oMax = p
				}
			} else if b == 'x' {
				cm := checkMap()
				p := check(blocks, i, j, cm)
				if p > xMax {
					xMax = p
				}
			}
		}
	}

	return oMax + xMax
}

func checkMap() [][]bool {
	cm := make([][]bool, N+2)
	for i := 0; i < N+2; i++ {
		cm[i] = make([]bool, N+2)
	}

	return cm
}

func check(blocks [][]byte, i, j int, cm [][]bool) int {
	if cm[i][j] {
		return 0
	}

	cm[i][j] = true
	r := blocks[i][j]
	res := 0

	// Top.
	if blocks[i-1][j] == r {
		res++
		res += check(blocks, i-1, j, cm)
	}

	// Bottom.
	if blocks[i+1][j] == r {
		res++
		res += check(blocks, i+1, j, cm)
	}

	// Left.
	if blocks[i][j-1] == r {
		res++
		res += check(blocks, i, j-1, cm)
	}

	// Right.
	if blocks[i][j+1] == r {
		res++
		res += check(blocks, i, j+1, cm)
	}

	return res
}

func drop(blocks [][]byte, i, j int) {
	for k := i; k >= 1; k-- {
		if k == 1 || blocks[k-1][j] == '-' {
			blocks[k][j] = ' '
			continue
		} else if blocks[k-1][j] == '.' {
			drop(blocks, k-1, j)
		}

		blocks[k][j] = blocks[k-1][j]
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

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

	fw := make([]string, N)
	fs := make([]string, N)
	fw[0], fw[1] = "", "W"
	fs[0], fs[1] = "", "S"
	okFW, okFS := true, true
	for i := 1; i < N; i++ {
		said := string(s[i-1])
		// First is wolf.
		switch said {
		case "o":
			if i+1 == N {
				if fw[i-1] == "S" {
					if fw[0] != "W" {
						okFW = false
					}
				} else {
					if fw[0] != "S" {
						okFW = false
					}
				}
			} else {
				if fw[i-1] == "" {
					fw[i-1] = "S"
					fw[i+1] = "W"
				} else if fw[i-1] == "S" {
					fw[i+1] = "W"
				} else {
					fw[i+1] = "S"
				}
			}
		case "x":
			if i+1 == N {
				if fw[i-1] == "S" {
					if fw[0] != "S" {
						okFW = false
					}
				} else {
					if fw[0] != "W" {
						okFW = false
					}
				}
			} else {
				if fw[i-1] == "" {
					fw[i-1] = "S"
					fw[i+1] = "S"
				} else if fw[i-1] == "S" {
					fw[i+1] = "S"
				} else {
					fw[i+1] = "W"
				}
			}
		}

		// First is sheep.
		switch said {
		case "o":
			if i+1 == N {
				if fs[i-1] == "S" {
					if fs[0] != "S" {
						okFS = false
					}
				} else {
					if fs[0] != "W" {
						okFS = false
					}
				}
			} else {
				if fs[i-1] == "" {
					fs[i-1] = "S"
					fs[i+1] = "S"
				} else if fs[i-1] == "S" {
					fs[i+1] = "S"
				} else {
					fs[i+1] = "W"
				}
			}
		case "x":
			if i+1 == N {
				if fs[i-1] == "S" {
					if fs[0] != "W" {
						okFS = false
					}
				} else {
					if fs[0] != "S" {
						okFS = false
					}
				}
			} else {
				if fs[i-1] == "" {
					fs[i-1] = "S"
					fs[i+1] = "W"
				} else if fs[i-1] == "S" {
					fs[i+1] = "W"
				} else {
					fs[i+1] = "S"
				}
			}
		}
	}

	if okFW {
		fmt.Println(strings.Join(fw, ""))
	} else if okFS {
		fmt.Println(strings.Join(fs, ""))
	} else {
		fmt.Println(-1)
	}
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

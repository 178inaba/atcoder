package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	N := uint64(nextInt())

	var wg sync.WaitGroup
	ch := make(chan uint64)
	for m := uint64(1); m <= N; m++ {
		wg.Add(1)
		go func(m uint64) {
			defer wg.Done()

			d := N / m
			mod := N % m
			if d == mod {
				ch <- m
			}
		}(m)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var ans uint64
	for {
		n, ok := <-ch
		if !ok {
			break
		}

		ans += n
	}

	fmt.Println(ans)
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

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
	N := nextInt()
	M := nextInt()
	ngCards := map[int]bool{}
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case index, ok := <-ch:
				if !ok {
					done <- struct{}{}
					return
				}
				ngCards[index] = false
			}
		}
	}()
	var wg sync.WaitGroup
	for i := 0; i < M; i++ {
		wg.Add(1)
		L, R := nextInt(), nextInt()
		go func() {
			defer wg.Done()
			for j := 1; j <= N; j++ {
				if j < L || j > R {
					ch <- j
				}
			}
		}()
	}
	wg.Wait()
	close(ch)
	<-done

	fmt.Println(N - len(ngCards))
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

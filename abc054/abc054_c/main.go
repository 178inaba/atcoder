package main

import "fmt"

type data struct {
	graph [][]bool
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	graph := make([][]bool, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		graph[a-1][b-1], graph[b-1][a-1] = true, true
	}

	d := data{graph: graph}

	visited := make([]bool, n)
	visited[0] = true

	fmt.Println(d.dfs(0, n, visited))
}

func (d data) dfs(v, n int, visited []bool) int {
	allVisited := true
	for i := 0; i < n; i++ {
		if !visited[i] {
			allVisited = false
		}
	}

	if allVisited {
		return 1
	}

	var ret int
	for i := 0; i < n; i++ {
		if !d.graph[v][i] {
			continue
		}

		if visited[i] {
			continue
		}

		visited[i] = true
		ret += d.dfs(i, n, visited)
		visited[i] = false
	}

	return ret
}

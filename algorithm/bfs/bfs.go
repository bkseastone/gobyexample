package main

import "fmt"

var (
	pl   = fmt.Println
	pf   = fmt.Printf
	maze = [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	begin = &point{0, 0}
	end   = &point{len(maze) - 1, len(maze[0]) - 1}
)

type point struct {
	x, y int
}

func bfs(maze [][]int, begin, end *point) ([][2]int, error) {
	width := len(maze)
	height := len(maze[0])
	var steps, queue [][2]int
	queue = append(queue, [2]int{begin.x, begin.y})
	for p := range queue {

	}
	steps = append(steps, [2]int{begin.x, begin.y})

	return nil, nil

}

// 广度优先算法
func main() {
	// steps := make([][]int, 1)
	// pl(len(maze))
	steps, err := bfs(maze, begin, end)
}

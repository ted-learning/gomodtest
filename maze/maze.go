package main

import (
	"fmt"
	"os"
)

func main() {
	maze := readMaze("maze/data/maze.txt")

	fmt.Println("Loading maze...")
	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%3d", maze[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Walking...")
	steps := walk(&maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for i := range *steps {
		for j := range (*steps)[i] {
			fmt.Printf("%3d", (*steps)[i][j])
		}
		fmt.Println()
	}
}

type point struct {
	i, j int
}

var directions = [4]point{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

func (cur point) add(direction point) point {
	return point{cur.i + direction.i, cur.j + direction.j}
}

func (cur point) accept(m *[][]int) bool {
	if cur.i < 0 || cur.i > len(*m)-1 {
		return false
	}

	if cur.j < 0 || cur.j > len((*m)[0])-1 {
		return false
	}

	if (*m)[cur.i][cur.j] > 0 {
		return false
	}

	return true
}

func walk(maze *[][]int, start, end point) *[][]int {
	q := []point{start}
	steps := make([][]int, len(*maze))
	for i := range steps {
		steps[i] = make([]int, len((*maze)[0]))
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == end {
			break
		}
		for _, direction := range directions {
			next := cur.add(direction)

			if !next.accept(maze) {
				continue
			}

			if !next.accept(&steps) {
				continue
			}

			if next == start {
				continue
			}

			q = append(q, next)
			steps[next.i][next.j] = steps[cur.i][cur.j] + 1
		}
	}
	return &steps
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, err = fmt.Fscanf(file, "%d %d", &row, &col)
	if err != nil {
		return nil
	}

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
	}

	for i := range maze {
		for j := range maze[i] {
			_, err = fmt.Fscanf(file, "%d", &maze[i][j])
			if err != nil {
				return nil
			}
		}
	}

	return maze
}

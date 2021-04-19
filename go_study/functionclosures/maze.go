package main

import "fmt"

func genMaze(n int) [][]int {
	var maze [][]int
	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < n; j++ {
			row = append(row, j+1+i)
		}
		maze = append(maze, row)
	}
	return maze
}

func printMaze(maze [][]int) {
	for i := 0; i < len(maze); i++ {
		fmt.Println(maze[i])
	}
}

type Direction [2]int

var RIGHT Direction = [2]int{0, 1}
var DOWN Direction = [2]int{1, 0}
var LEFT Direction = [2]int{0, -1}
var TOP Direction = [2]int{-1, 0}

func (d Direction) Next() Direction {
	switch d {
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return TOP
	}
	return RIGHT
}

func printRotate() func(maze [][]int, n int) int {
	direction := RIGHT
	poi := []int{0, 0}
	return func(maze [][]int, n int) int {
		for i := 1; i < n; i++ {
			for {
				x, y := poi[0]+direction[0], poi[1]+direction[1]
				if x < len(maze) && y < len(maze[0]) {
					poi[0] = x
					poi[1] = y
					break
				} else {
					direction = direction.Next()
				}
			}
		}
		return maze[poi[0]][poi[1]]
	}
}

func main() {
	maze := genMaze(10)
	printMaze(maze)

	printA := printRotate()
	fmt.Println(printA(maze, 23))
}

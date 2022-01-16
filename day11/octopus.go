package day11

import "fmt"

var R = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var C = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func FlashCountOld(grid [][]int, steps int) int {
	count := 0
	fmt.Println("Before any steps:")
	printGrid(grid)
	h, w := len(grid), len(grid[0])
	for i := 1; i <= steps; i++ {
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				grid[r][c]++
				if grid[r][c] > 9 {
					//flash
					for n := 0; n < 8; n++ {
						//flash neighbours
						row, col := r+R[n], c+C[n]
						if row >= 0 && row < h && col >= 0 && col < w {
							grid[row][col]++
						}
					}
				}
			}
		}
		// set all values > 9 to 0
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				if grid[r][c] > 9 {
					count++
					grid[r][c] = 0

				}
			}
		}
		fmt.Printf("After step %d:\n", i)
		printGrid(grid)
	}
	return count
}

func FlashCount(grid [][]int, steps int) int {
	count := 0
	fmt.Println("Before any steps:")
	printGrid(grid)
	h, w := len(grid), len(grid[0])
	for i := 1; i <= steps; i++ {
		// Step 1 : increase every octopus' energy level by 1
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				grid[r][c]++
			}
		}

		// Step 2 : Any octopus > 9 flashes. Neighbours increase by 1. Also flash if > 9
		hasFlashed := make(map[point]bool)
		flashed := queue{}
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				if grid[r][c] > 9 {
					p := point{r, c}
					flashed.Enqueue(p)
					hasFlashed[p] = true
				}
			}
		}

		for len(flashed) > 0 {
			curr := flashed.Dequeue()
			//			fmt.Println("flashing", curr)
			for n := 0; n < 8; n++ {
				row, col := curr.R+R[n], curr.C+C[n]
				if row >= 0 && row < h && col >= 0 && col < w {
					grid[row][col]++
					if grid[row][col] > 9 && !hasFlashed[point{row, col}] {
						p := point{row, col}
						flashed.Enqueue(p)
						hasFlashed[p] = true
					}
				}
			}
		}
		// Step 3: Any octopus that flashed is set to 0
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				if grid[r][c] > 9 {
					count++
					grid[r][c] = 0

				}
			}
		}
		if i%10 == 0 {
			fmt.Printf("After step %d:\n", i)
			printGrid(grid)
		}
	}
	return count
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

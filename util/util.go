package util 

import "fmt"

func Reverse[T any](t []T) {
	for i := len(t)/2 -1; i >= 0; i-- {
		opp := len(t) - 1 - i
		t[i], t[opp] = t[opp], t[i]
	}
}

func Map[T any](t []T, f func(T) T) []T {
	res := make([]T, 0)
	for _, e := range t {
		res = append(res, f(e))
	}
	return res
}

func PrintGrid(grid [][]int) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func Has[T comparable](slice []T, item T) bool {
	for _,e := range slice {
		if e == item {
			return true
		}
	}
	return false
}
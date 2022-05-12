package main

import (
	"fmt"
	"time"
)
func draw(arr []int, s int) {
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			if arr[i] == j {
				fmt.Print("♕\t")
			} else {
				fmt.Print("0\t")
			}
		}
		fmt.Println(" ")
	}
}

func check_diameter(arr []int, r int, c int) bool {
	for i := 0; i < r; i++ {
		t1, t2 := arr[i], arr[i]
		for j := i; j < r+1; j++ {
			if (t1 == c && j == r) || (t2 == c && j == r) {
				return false
			}
			t1++
			t2--
		}
	}
	return true
}

func check_col(arr []int, r int, c int) bool {
	for i := 0; i < r; i++ {
		if arr[i] == c {
			return false
		}
	}
	return true
}

func set_queen(arr []int, r int, c int, size int) bool {
	for i := c; i < size; i++ {
		if check_col(arr, r, i) && check_diameter(arr, r, i) {
			arr[r] = i
			return true
		}
	}
	return false
}

func n_queen(arr []int, r int, c int, size int) bool {
	if size < 4 {
		fmt.Println("ERROR: el número debe ser mayor a 3...")
		return false
	} else if r == size {
		return true
	} else if set_queen(arr, r, c, size) {
		n_queen(arr, r+1, 0, size)
	} else {
		return n_queen(arr, r-1, arr[r-1]+1, size)
	}
	return true
}

func main() {
	var size int
	start := time.Now()
	fmt.Print("Tamaño del tablero (N): ")
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := range arr {
		arr[i] = -1
	
	for i := 5;i<5; i++ {
		start := n_queen(arr, 0, 0, size)
		start = 
	}
}
	draw(arr, size)

}
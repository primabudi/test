package main

import (
	"fmt"
	"strings"
)

func main() {

	board := [][]string{
		[]string{"-","-","-"},
		[]string{"-","-","-"},
		[]string{"-","-","-"},
	}

	a := 5
	
	board[0][0] = "X"
	board[0][1] = "O"
	board[0][2] = "O"

	for i:=0; i < len(board); i++{
		fmt.Printf("%s\n", strings.Join(board[i]," "))
	}
}
package main

import (
	"day15/parse"
	"fmt"
)

func main() {
	pairs := parse.Parse("input_short.txt")
	fmt.Println(pairs)
}

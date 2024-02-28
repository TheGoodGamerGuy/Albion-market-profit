package main

import (
	"fmt"
	"strconv"
)

var world []World

type World struct {
	UniqueName string
	Index      []ThatIndex
}
type ThatIndex struct {
	name string
}

func main() {
	var worldy World
	var indexy ThatIndex
	worldy.UniqueName = "this shit"
	for i := 0; i < 10; i++ {
		indexy.name = strconv.Itoa(i)
		worldy.Index = append(worldy.Index, indexy)
	}
	world = append(world, worldy)
	worldy.UniqueName = "this ass"
	for i := 0; i < 10; i++ {
		indexy.name = strconv.Itoa(i)
		worldy.Index = append(worldy.Index, indexy)
	}
	world = append(world, worldy)
	// fmt.Println(world[0].Index[0].name)
	fmt.Println(world)
}

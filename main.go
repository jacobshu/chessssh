package main

import (
	"fmt"

	game "github.com/jacobshu/chessssh/internal"
)

func main() {
	g := game.Game{}
	fmt.Println(g.String(true))
}

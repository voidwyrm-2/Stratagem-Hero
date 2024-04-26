package main

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Stratagem struct {
	name string
	code string
}

var Stratagems = []Stratagem{
	{"MG-43 Machine Gun", "dldur"},
}

func (sg Stratagem) getKeys() []int {
	var out []int
	for _, dir := range strings.Split(sg.code, "") {
		switch dir {
		case "u":
			out = append(out, rl.KeyUp)
		case "d":
			out = append(out, rl.KeyDown)
		case "l":
			out = append(out, rl.KeyLeft)
		case "r":
			out = append(out, rl.KeyRight)
		}
	}
	return out
}

func main() {
	//var timer int = 100
	//var timerTick int = 10
	//tick := 0
	rl.InitWindow(800, 450, "Stratagem Hero")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		//rl.DrawText("Creeper, oh man", 316, 200, 20, rl.White)

		rl.EndDrawing()
	}
}

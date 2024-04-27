package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Stratagem struct {
	name string
	code string
	kind string
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
		default:
			fmt.Printf("error: while getting keys for Stratagem '%s', found illegal direction '%s'", sg.name, dir)
		}
	}
	return out
}

var stratagemPool = []Stratagem{
	{"MG-43 Machine Gun", "dldur", "Supply"},
	{"APW-1 Anti-Materiel Rifle", "dlrud", "Supply"},
	{"M-105 Stalwart", "dlduul", "Supply"},
	{"EAT-17 Expendable Anti-Tank", "ddlur", "Supply"},
	{"GR-8 Recoilless Rifle", "dlrrl", "Supply"},
}

func shuffleStratagems(sPool []Stratagem) []Stratagem {
	var out = make([]Stratagem, len(sPool))

	perm := rand.Perm(len(sPool))
	for i, v := range perm {
		out[v] = sPool[i]
	}
	return out
}

func DrawUpArrow(posX, posY float32, col color.RGBA) {
	rl.DrawLineEx(rl.Vector2{X: posX - 2.75, Y: posY - 15}, rl.Vector2{X: posX + 15, Y: posY}, 10, col) // right
	rl.DrawLineEx(rl.Vector2{X: posX + 2.75, Y: posY - 15}, rl.Vector2{X: posX - 15, Y: posY}, 10, col) // left
	rl.DrawLineEx(rl.Vector2{X: posX, Y: posY - 15}, rl.Vector2{X: posX, Y: posY + 15}, 10, col)        // middle
}

func DrawDownArrow(posX, posY float32, col color.RGBA) {
	rl.DrawLineEx(rl.Vector2{X: posX - 2.75, Y: posY + 15}, rl.Vector2{X: posX + 15, Y: posY}, 10, col) // left
	rl.DrawLineEx(rl.Vector2{X: posX + 2.75, Y: posY + 15}, rl.Vector2{X: posX - 15, Y: posY}, 10, col) // right
	rl.DrawLineEx(rl.Vector2{X: posX, Y: posY + 15}, rl.Vector2{X: posX, Y: posY - 15}, 10, col)        // middle
}

func DrawLeftArrow(posX, posY float32, col color.RGBA) {
	rl.DrawLineEx(rl.Vector2{X: posX - 15, Y: posY - 2.75}, rl.Vector2{X: posX, Y: posY + 15}, 10, col) // right
	rl.DrawLineEx(rl.Vector2{X: posX - 15, Y: posY + 2.75}, rl.Vector2{X: posX, Y: posY - 15}, 10, col) // left
	rl.DrawLineEx(rl.Vector2{X: posX - 15, Y: posY}, rl.Vector2{X: posX + 15, Y: posY}, 10, col)        // middle
}

func DrawRightArrow(posX, posY float32, col color.RGBA) {
	rl.DrawLineEx(rl.Vector2{X: posX + 15, Y: posY - 2.75}, rl.Vector2{X: posX, Y: posY + 15}, 10, col) // left
	rl.DrawLineEx(rl.Vector2{X: posX + 15, Y: posY + 2.75}, rl.Vector2{X: posX, Y: posY - 15}, 10, col) // right
	rl.DrawLineEx(rl.Vector2{X: posX + 15, Y: posY}, rl.Vector2{X: posX - 15, Y: posY}, 10, col)        // middle
}

var HelldiversYellow color.RGBA = color.RGBA{240, 240, 134, 255}

var lost = false

var windowX int32 = 1000
var windowY int32 = 650

var arrowsX = 400
var arrowsY = 475
var arrowsSpacing = 40

func main() {

	//imgTexture := rl.LoadTexture("./assets/Arrow.png")

	var timer int = 100
	var timerTickrate int = 8
	var timerTick int = 0
	var canTickTimer bool = true

	var redTick int = 0

	var currentStratagem int = 0

	var currentKey int = 0

	fmt.Println(stratagemPool)

	var stratagems []Stratagem = shuffleStratagems(stratagemPool)

	fmt.Println(stratagemPool)

	var cSDirs []int = stratagems[currentStratagem].getKeys()

	var score int = 0
	rl.InitWindow(windowX, windowY, "Stratagem Hero")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if lost {
			if rl.IsKeyPressed(rl.KeyR) {
				timer = 100
				score = 0
				timerTick = 0
				redTick = 0
				currentStratagem = 0
				currentKey = 0
				lost = false
			}
		} else {
			if currentStratagem+1 >= len(stratagems) {
				currentStratagem = 0
				currentKey = 0
				timer = 100
				stratagems = shuffleStratagems(stratagemPool)
				cSDirs = stratagems[currentStratagem].getKeys()
			} else {
				if currentKey >= len(cSDirs) {
					currentStratagem++
					currentKey = 0
					score += 15
					cSDirs = stratagems[currentStratagem].getKeys()
					timer += 15
					if timer > 100 {
						timer = 100
					}
				}

				gottenKey := rl.GetKeyPressed()
				if gottenKey == int32(cSDirs[currentKey]) {
					currentKey++
				} else if gottenKey != rl.KeyEscape && gottenKey != 0 {
					//fmt.Println(rl.GetKeyPressed())
					currentKey = 0
					redTick = 10
				}
			}

		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		//rl.DrawText("Creeper, oh man", 316, 200, 20, rl.White)

		//rl.DrawTextureRec(imgTexture, rl.Rectangle{X: 100, Y: 100, Width: 100, Height: 100}, rl.Vector2{X: 100, Y: 100}, color.RGBA{255, 255, 255, 255})

		if lost {
			rl.DrawText("You lost!", int32(float32(windowX)/2.5)+25, 200, 20, rl.White)
			rl.DrawText(fmt.Sprintf("Your score: %v", score), (int32(float32(windowX)/2.5)+25)-32, 230, 20, rl.White)
			rl.DrawText("Press R to restart or Escape to quit", (int32(float32(windowX)/2.5)+25)-150, 270, 20, rl.White)
		} else {

			if redTick > 0 {
				redTick--
			}

			rl.DrawText(fmt.Sprintf("%v/%v", currentStratagem+1, len(stratagems)), 20, 20, 20, rl.White)

			for i, dir := range stratagems[currentStratagem].code {
				if currentKey > i {
					if redTick > 0 {
						//rl.DrawText(string(dir), int32((50-len(stratagems[currentStratagem].code))+(i*20)), 100, 20, rl.Red)
						switch dir {
						case 'u':
							DrawUpArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.Red)
						case 'd':
							DrawDownArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.Red)
						case 'l':
							DrawLeftArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.Red)
						case 'r':
							DrawRightArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.Red)
						default:
							DrawUpArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.Magenta)
						}
					} else {
						//rl.DrawText(string(dir), int32((50-len(stratagems[currentStratagem].code))+(i*20)), 100, 20, HelldiversYellow)
						switch dir {
						case 'u':
							DrawUpArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), HelldiversYellow)
						case 'd':
							DrawDownArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), HelldiversYellow)
						case 'l':
							DrawLeftArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), HelldiversYellow)
						case 'r':
							DrawRightArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), HelldiversYellow)
						default:
							DrawUpArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.Magenta)
						}
					}
				} else {
					//rl.DrawText(string(dir), int32((50-len(stratagems[currentStratagem].code))+(i*20)), 100, 20, rl.White)
					switch dir {
					case 'u':
						DrawUpArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.White)
					case 'd':
						DrawDownArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.White)
					case 'l':
						DrawLeftArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.White)
					case 'r':
						DrawRightArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.White)
					default:
						DrawUpArrow(float32((arrowsX-len(stratagems[currentStratagem].code))+(i*arrowsSpacing)), float32(arrowsY), rl.Magenta)
					}
				}
			}

			//fmt.Println(timer)

			if timer <= 0 {
				//timer = 100
				lost = true
			}

			if timerTick >= timerTickrate && canTickTimer {
				timer--
				timerTick = 0
			} else if canTickTimer {
				timerTick++
			}

			rl.DrawRectangle(int32(windowX/2)-215, 550, int32(timer*4), 35, HelldiversYellow)
		}

		rl.EndDrawing()
	}
}

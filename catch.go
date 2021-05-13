package main

import (
	"log"
	"fmt"
	"time"
	_ "image/png"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var catcher *ebiten.Image
var ball *ebiten.Image
var randNum float64

type Game struct{
	balls [1000]*ebiten.Image
	ballNum int
	ballX float64
	ballY float64
	catcherX float64
	caught bool
	fallRate float64
	score int
}

func randomXPos() {
	rand.Seed(time.Now().UnixNano())
	randNum = float64(rand.Intn(300 - 20 + 1) + 20)
}

func init() {
	var err error
	catcher, _, err = ebitenutil.NewImageFromFile("catcher.png")
	if err != nil {
		log.Fatal(err)
	}

	ball = ebiten.NewImage(10, 10)
	ball.Fill(color.White)

	randomXPos()
}

func (g *Game) Update() error {
	g.ballX = randNum

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.catcherX += -5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.catcherX += 5
	}

	if !(g.ballY >= 240) && g.ballY == 200 && (g.ballX > g.catcherX && g.ballX < g.ballX + 40) {
		g.score += 10
		g.caught = true
	}

	if g.caught == true {
		g.caught = false
		g.ballY = -10
		randomXPos()
	}

	switch true {
		case g.catcherX == 0:
			g.catcherX = 140
			break
		case g.catcherX <= 10:
				g.catcherX = 10
			break
		case g.catcherX >= 270:
			g.catcherX = 270
			break
	}

	switch true {
		case g.score < 50:
				g.fallRate = 1
				break
			case g.score >= 50 && g.score < 150:
				g.fallRate = 2
				break
			case g.score >= 150:
				g.fallRate = 3
				break
	}

	g.ballY += g.fallRate

	return nil
}



func (g *Game) Draw(screen *ebiten.Image) {

	ballOps := &ebiten.DrawImageOptions{}
	ballOps.GeoM.Translate(g.ballX, g.ballY)

	catcherOps := &ebiten.DrawImageOptions{}
	catcherOps.GeoM.Translate(g.catcherX, 200)

	screen.DrawImage(catcher, catcherOps)
	screen.DrawImage(ball, ballOps)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v", g.score))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Catch")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

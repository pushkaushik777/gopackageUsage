package main1

import (
	"image/color"
	"math"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 300
	screenHeight = 300

	ballRadius = 15
)

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var (
	ballPositionX = float64(screenWidth) / 2
	ballPositionY = float64(screenHeight) / 2
	ballMovementX = float64(0.5)
	ballMovementY = float64(0.75)
)

func (g *Game) Update() error {
	ballPositionX += ballMovementX
	ballPositionY += ballMovementY

	if ballPositionX >= screenWidth-ballRadius || ballPositionX <= ballRadius {
		ballMovementX *= -1
	}

	if ballPositionY >= screenHeight-ballRadius || ballPositionY <= ballRadius {
		ballMovementY *= -1
	}

	return nil
}

func (g *Game) drawCircle(screen *ebiten.Image, x, y, radius int, clr color.Color) {
	radius64 := float64(radius)
	minAngle := math.Acos(1 - 1/radius64)

	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := radius64 * math.Cos(angle)
		yDelta := radius64 * math.Sin(angle)

		x1 := int(math.Round(float64(x) + xDelta))
		y1 := int(math.Round(float64(y) + yDelta))

		screen.Set(x1, y1, clr)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	purpleClr := color.RGBA{056, 0, 255, 255}
	purpleCl1 := color.RGBA{100, 50, 255, 255}
	purpleCl2 := color.RGBA{100, 0, 255, 255}

	x := int(math.Round(ballPositionX))
	y := int(math.Round(ballPositionY))

	g.drawCircle(screen, x, y, ballRadius, purpleClr)
	g.drawCircle(screen, x+10, y+10, ballRadius, purpleCl1)
	g.drawCircle(screen, x-10, y-10, ballRadius, purpleCl2)
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("The Moving Ball")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

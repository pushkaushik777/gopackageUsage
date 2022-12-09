package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const screenWidth int = 400
const screenHeight int = 500
const playerwidth float64 = 93
const playerheight float64 = 94
const enemywidth float64 = 164
const enemyheight float64 = 142

//var n []ebiten.Key

type scoreBoard struct {
	Enemyhit  int
	Playerhit int
}

type Game struct {
	player *ebiten.Image
	pPosx  float64
	pPosy  float64

	bullet *ebiten.Image
	bPosx  float64
	bPosy  float64

	enemy *ebiten.Image
	ePosx float64
	ePosy float64

	ePossMx float64

	bomb     *ebiten.Image
	bombPosx float64
	bombPosy float64

	score scoreBoard
}

func (g *Game) Update() error {

	for _, k := range inpututil.PressedKeys() {
		if k == ebiten.Key(ebiten.KeyRight) {
			g.pPosx += 5
		} else if k == ebiten.Key(ebiten.KeyLeft) {
			g.pPosx -= 5
		} else if k == ebiten.Key(ebiten.KeyUp) {
			g.pPosy -= 5
		} else if k == ebiten.Key(ebiten.KeyDown) {
			g.pPosy += 5
		}
	}

	if g.ePosx+enemywidth > float64(screenWidth) {
		g.ePossMx = -3
	}

	if g.ePosx <= 0 {
		g.ePossMx = 3
	}

	g.ePosx += g.ePossMx

	g.bPosy -= 4

	g.bombPosy += 3

	if g.pPosx+playerwidth >= float64(screenWidth) {
		g.pPosx = float64(screenWidth) - float64(playerwidth)
	}

	if g.pPosy+playerheight >= float64(screenHeight) {
		g.pPosy = float64(screenHeight) - float64(playerheight)
	}

	if g.pPosx <= 0 {
		g.pPosx = 0
	}

	if g.pPosy <= 0 {
		g.pPosy = 0
	}

	if g.bPosy <= 0 {
		g.bPosx = g.pPosx + playerwidth/2
		g.bPosy = g.pPosy - 5
	}

	if g.bPosy <= g.ePosy+float64(enemyheight) &&
		g.bPosx <= g.ePosx+float64(enemywidth) &&
		g.bPosx >= g.ePosx {
		//fmt.Println("Enemy hurt ..")
		g.bPosx = g.pPosx + playerwidth/2
		g.bPosy = g.pPosy - 5
		g.score.Enemyhit += 2

	}

	if g.bombPosy >= float64(screenHeight) {
		g.bombPosx = g.ePosx + enemywidth/2
		g.bombPosy = enemyheight
	}

	if g.bombPosy >= g.pPosy &&
		g.pPosx+float64(playerwidth) >= g.bombPosx &&
		g.pPosx <= g.bombPosx {
		g.bombPosx = g.ePosx + enemywidth/2
		g.bombPosy = enemyheight
		fmt.Println("Player hurt ..", g.bombPosy, g.pPosy, g.bombPosx, g.pPosx)
		g.score.Playerhit += 2

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Opaque)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.pPosx, g.pPosy)
	screen.DrawImage(g.player, op)

	opblt := &ebiten.DrawImageOptions{}
	opblt.GeoM.Translate(g.bPosx, g.bPosy)
	screen.DrawImage(g.bullet, opblt)

	openemy := &ebiten.DrawImageOptions{}
	openemy.GeoM.Translate(g.ePosx, g.ePosy)
	screen.DrawImage(g.enemy, openemy)

	opbomb := &ebiten.DrawImageOptions{}
	opbomb.GeoM.Translate(g.bombPosx, g.bombPosy)
	screen.DrawImage(g.bomb, opbomb)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Enemyhit :%d PlayerHit :%d", g.score.Enemyhit, g.score.Playerhit))
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("", g.score.Playerhit))
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.board.points))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 500
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("MyFirstGame")

	img, _, err := ebitenutil.NewImageFromFile("pPics/space.png")
	if err != nil {
		log.Fatalf("Image Load error : %v", err)
	}

	blt := ebiten.NewImage(5, 10)
	blt.Fill(color.Black)

	imgEnmy, _, err := ebitenutil.NewImageFromFile("pPics/enemys.png")
	if err != nil {
		log.Fatalf("Enemy Image load error : %v", err)
	}

	//imgbmb, _, err := ebitenutil.NewImageFromFile("pPics/bomb.png")
	imgbmb := ebiten.NewImage(8, 14)
	imgbmb.Fill(color.YCbCr{
		Y:  12,
		Cb: 20,
		Cr: 40,
	})
	if err != nil {
		log.Fatalf("Bomb Image load error : %v", err)
	}

	g := &Game{}
	g.player = img
	g.pPosx = float64(screenHeight/2) - playerwidth
	g.pPosy = float64(screenWidth / 2)

	g.bullet = blt
	g.bPosx = g.pPosx + playerwidth/2
	g.bPosy = g.pPosy - 5

	g.pPosy = float64(screenWidth / 2)
	fmt.Println(g.pPosx, g.pPosy)

	g.enemy = imgEnmy
	g.ePosx = 0
	g.ePosy = 0

	g.ePossMx = 3

	g.bomb = imgbmb
	g.bombPosx = g.ePosx + enemywidth/2
	g.bombPosy = enemyheight - 5

	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("Run game error : %v", err)
	}

}

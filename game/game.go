package game

import (
	"fmt"
	"github.com/fralonra/go-2048/core"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

const (
	cellSize   = 80
	cellMargin = 4
)

var (
	game      = core.NewGame()
	cellImage *ebiten.Image
)

type Game struct {
	screen *ebiten.Image
}

func (g *Game) checkKeyEvents() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		game.ToTop()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		game.ToBottom()
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		game.ToLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		game.ToRight()
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		game = core.NewGame()
	}
}

func (g *Game) renderGame(screen *ebiten.Image) {
	for idx := 0; idx < core.Size; idx++ {
		row := game.GetRow(idx)
		for col, item := range row {
			op := &ebiten.DrawImageOptions{}
			x := idx*cellSize + (col+1)*cellMargin
			y := col*cellSize + (idx+1)*cellMargin
			op.GeoM.Translate(float64(x), float64(y))
			r, g, b, a := colorToScale(getCellColor(item))
			op.ColorM.Scale(r, g, b, a)
			screen.DrawImage(cellImage, op)
		}
	}
}

func (g *Game) checkState(screen *ebiten.Image) {
	switch game.State {
	case core.StateNormal:
		{
			msg := fmt.Sprintf("Current point: %d", game.MaxNumber)
			ebitenutil.DebugPrint(screen, msg)
		}
	case core.StateWin:
		{
			msg := fmt.Sprintf("You won! Your point: %d", game.MaxNumber)
			ebitenutil.DebugPrint(screen, msg)
		}
	case core.StateLost:
		{
			ebitenutil.DebugPrint(screen, "You lost! Press 'r' to restart")
		}
	}
}

func (g *Game) Update(screen *ebiten.Image) {
	g.checkKeyEvents()
	g.renderGame(screen)
	g.checkState(screen)
}

func getCellColor(value int) color.Color {
	switch value {
	case 0:
		return color.NRGBA{0xee, 0xe4, 0xda, 0x59}
	case 2:
		return color.RGBA{0xee, 0xe4, 0xda, 0xff}
	case 4:
		return color.RGBA{0xed, 0xe0, 0xc8, 0xff}
	case 8:
		return color.RGBA{0xf2, 0xb1, 0x79, 0xff}
	case 16:
		return color.RGBA{0xf5, 0x95, 0x63, 0xff}
	case 32:
		return color.RGBA{0xf6, 0x7c, 0x5f, 0xff}
	case 64:
		return color.RGBA{0xf6, 0x5e, 0x3b, 0xff}
	case 128:
		return color.RGBA{0xed, 0xcf, 0x72, 0xff}
	case 256:
		return color.RGBA{0xed, 0xcc, 0x61, 0xff}
	case 512:
		return color.RGBA{0xed, 0xc8, 0x50, 0xff}
	case 1024:
		return color.RGBA{0xed, 0xc5, 0x3f, 0xff}
	case 2048:
		return color.RGBA{0xed, 0xc2, 0x2e, 0xff}
	}
	return color.White
}

func colorToScale(clr color.Color) (float64, float64, float64, float64) {
	r, g, b, a := clr.RGBA()
	rf := float64(r) / 0xffff
	gf := float64(g) / 0xffff
	bf := float64(b) / 0xffff
	af := float64(a) / 0xffff
	// Convert to non-premultiplied alpha components.
	if 0 < af {
		rf /= af
		gf /= af
		bf /= af
	}
	return rf, gf, bf, af
}

func init() {
	cellImage, _ = ebiten.NewImage(cellSize, cellSize, ebiten.FilterDefault)
	cellImage.Fill(color.White)
}

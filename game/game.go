package game

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/fralonra/ebiten-2048/assets/fonts"
	"github.com/fralonra/go-2048/colors"
	"github.com/fralonra/go-2048/core"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

const (
	cellSize   = 76
	cellMargin = 4
)

var (
	game       = core.NewGame()
	frameColor = color.RGBA{0xbb, 0xad, 0xa0, 0xff}

	cellImage       *ebiten.Image
	mplusSmallFont  font.Face
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

type Game struct {
	cellTable [core.Size][core.Size]*ebiten.Image
}

func (g *Game) Setup() {

}

func (g *Game) Update(screen *ebiten.Image) {
	g.checkKeyEvents()

	if ebiten.IsDrawingSkipped() {
		return
	}

	g.renderGame(screen)
	g.checkState(screen)
}

func (g *Game) checkKeyEvents() {
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		game.ToTop()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		game.ToBottom()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		game.ToLeft()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		game.ToRight()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyR) {
		game = core.NewGame()
	}
}

func (g *Game) renderGame(screen *ebiten.Image) {
	screen.Fill(frameColor)
	for i := 0; i < core.Size; i++ {
		row := game.GetRow(i)
		for j, item := range row {
			op := &ebiten.DrawImageOptions{}
			x := j*(cellSize+cellMargin) + cellMargin
			y := i*(cellSize+cellMargin) + cellMargin
			op.GeoM.Translate(float64(x), float64(y))
			r, g, b, a := colors.ColorToScale(colors.CellColor(item))
			op.ColorM.Scale(r, g, b, a)
			screen.DrawImage(cellImage, op)

			if item != 0 {
				str := strconv.Itoa(item)
				f := mplusBigFont
				switch {
				case 3 < len(str):
					f = mplusSmallFont
				case 2 < len(str):
					f = mplusNormalFont
				}
				bound, _ := font.BoundString(f, str)
				w := (bound.Max.X - bound.Min.X).Ceil()
				h := (bound.Max.Y - bound.Min.Y).Ceil()
				x = x + (cellSize-w)/2
				y = y + (cellSize-h)/2 + h
				text.Draw(screen, str, f, x, y, color.White)
			}
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

func init() {
	cellImage, _ = ebiten.NewImage(cellSize, cellSize, ebiten.FilterDefault)
	cellImage.Fill(color.White)

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusSmallFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

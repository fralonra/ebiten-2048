package main

import (
	"github.com/fralonra/ebiten-2048/game"
	"github.com/hajimehoshi/ebiten"
	"log"
)

var (
	g *game.Game
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	g.Update(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 320, 2, "2048"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/fralonra/ebiten-2048/game"
	"github.com/hajimehoshi/ebiten"
	"log"
)

var (
	g = &game.Game{}
)

func update(screen *ebiten.Image) error {
	g.Update(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, 324, 324, 2, "2048"); err != nil {
		log.Fatal(err)
	}
	g.Setup()
}

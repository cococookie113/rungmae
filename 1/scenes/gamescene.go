package scenes

import (
	"github.com/hajimehoshi/ebiten"
)

// GameScene scene of game
type GameScene struct {
	bgimg *ebiten.Image
}

//Startup initialize GameScene
func (g *GameScene) Startup() {

}

// Update GameScene
func (g *GameScene) Update(screen *ebiten.Image) error {

	return nil
}

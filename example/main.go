package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/waloader"
	"github.com/umi-l/yosui-ui"
	"github.com/umi-l/yosui-ui/gui"
	"github.com/umi-l/yosui-ui/widgets"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	Atlas map[string]waloader.Sprite
	Gui   GameUI
}

func (g *Game) Update() error {

	if g.Gui.MainMenu.PlayButton.IsPressed() && g.Gui.MainMenu.PlayButton.IsVisible() {
		fmt.Print("Button Pressed \n")
		g.Gui.MainMenu.PlayButton.Visible = false
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	w, h := screen.Size()

	g.Gui.Root.UpdateTransform(gui.Transform{X: 0, Y: 0, W: float32(w), H: float32(h)})

	g.Gui.Root.DrawTree(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

type GameUI struct {
	Root     gui.Container
	MainMenu MainMenuUI
}

type MainMenuUI struct {
	PlayButton *widgets.GuiButton
}

func main() {

	game := Game{}

	game.Atlas = waloader.LoadAtlas("assets/atlases/", "atlas.xml")

	playButtonSprite := game.Atlas["MarioPlayButton"]

	game.Gui.Root = yosui.MakeRootContainer(screenWidth, screenHeight)

	//--main menu--
	mainMenu := yosui.MakeRelativeContainer(&game.Gui.Root, gui.Transform{X: 0, Y: 0, WPercent: 1, HPercent: 1}, true)

	//play button

	trans := gui.MakeTransformWithImage(playButtonSprite.Image, gui.OriginCenter)

	trans.XPercent = 0.5
	trans.YPercent = 0.5

	playButton := widgets.NewButton(playButtonSprite.Image, trans)

	//add to main menu
	mainMenu.AddChild(&playButton)

	game.Gui.MainMenu.PlayButton = &playButton

	//add to gui
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Yosui test")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

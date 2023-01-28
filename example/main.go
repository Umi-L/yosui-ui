package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/umi-l/waloader"
	"github.com/umi-l/yosui-ui/gui"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	runnerImage *ebiten.Image
)

type Game struct {
	Atlas map[string]waloader.Sprite
	Gui   GameUI
}

func (g *Game) Update() error {

	if g.Gui.MainMenu.PlayButton.IsPressed() {
		fmt.Print("Button Pressed \n")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	w, h := screen.Size()

	g.Gui.Root.SetTransform(gui.Transform{X: 0, Y: 0, W: float32(w), H: float32(h)})

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
	PlayButton *GuiButton
}

type GuiButton struct {
	gui.Element
}

func (b GuiButton) IsPressed() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		var x, y = ebiten.CursorPosition()
		if b.Rect.Contains((float32)(x), (float32)(y)) {
			return true
		}
	}

	return false
}

func NewButton(image *ebiten.Image) GuiButton {
	return GuiButton{
		Element: gui.MakeElement(image),
	}
}

func main() {

	game := Game{}

	game.Atlas = waloader.LoadAtlas("assets/atlases/", "atlas.xml")
	playButtonSprite := game.Atlas["MarioPlayButton"]

	game.Gui.Root.Visible = true

	//--main menu--
	mainMenu := gui.NewContainer(gui.Transform{X: 0, Y: 0, WPercent: 1, HPercent: 1}, true)

	//play button
	playButton := NewButton(playButtonSprite.Image)

	trans := gui.MakeTransformWithImage(playButtonSprite.Image, gui.OriginCenter)

	trans.XPercent = 0.5
	trans.YPercent = 0.5

	//add to main menu
	mainMenu.AddChild(&playButton)

	playButton.SetTransform(trans)

	game.Gui.MainMenu.PlayButton = &playButton

	//add to gui
	game.Gui.Root.AddChild(&mainMenu)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Yosui test")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

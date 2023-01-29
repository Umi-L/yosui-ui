package widgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/umi-l/yosui-ui/gui"
	"github.com/umi-l/yosui-ui/utils"
	"log"
)

type GuiButton struct {
	gui.Container

	Image *ebiten.Image

	initialized bool

	OnPressed func()
}

func (b GuiButton) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		var x, y = ebiten.CursorPosition()
		if b.Rect.Contains((float32)(x), (float32)(y)) {
			b.OnPressed()
		}
	}
}

func (b *GuiButton) Init() {
	b.Transform = gui.MakeTransformWithImage(b.Image, gui.OriginTopLeft)
	b.calculateRect()

	b.initialized = true
}

func (b GuiButton) checkInitialized() {
	if !b.initialized {
		log.Fatal("GuiButton is not initialized; this may cause unexpected behaviour and as such is an error. Call GuiButton.Init() to fix this.")
	}
}

func (b GuiButton) drawSelf() {
	b.checkInitialized()

	log.Print("Drawself")

	if !b.Visible {
		return
	}

	call := func(screen *ebiten.Image) {
		utils.DrawImageAtRect(screen, b.Image, b.Rect, &ebiten.DrawImageOptions{})
		log.Print("Drawing button")
	}
	b.GetRoot().AddToDrawStack(b.ZIndex, call)
}

func (b GuiButton) draw() {
	gui.Defaults.Draw(&b)
}

func (b *GuiButton) calculateRect() {
	b.Rect = gui.Defaults.CalculateRect(b)
}

func (b GuiButton) GetContainer() *gui.Container {
	return &b.Container
}

func (b *GuiButton) SetParent(parent *gui.Container) {
	b.Parent = parent
}

func NewButton(image *ebiten.Image, transform gui.Transform, onClick func()) GuiButton {
	button := GuiButton{
		Image:     image,
		OnPressed: onClick,
	}

	button.Init()

	button.Transform = transform

	button.Visible = true

	return button
}

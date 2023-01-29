package widgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/umi-l/yosui-ui/gui"
)

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

func NewButton(image *ebiten.Image, transform gui.Transform) GuiButton {
	button := GuiButton{
		Element: gui.MakeElement(image),
	}

	button.Transform = transform

	return button
}

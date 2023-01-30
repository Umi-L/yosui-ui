package widgets

import (
	"github.com/umi-l/yosui-ui/gui"
	"github.com/umi-l/yosui-ui/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Panel struct { //elements are just Containers with drawables
	gui.Container

	Image *ebiten.Image

	initialized bool
}

func (e *Panel) Init() {
	e.Transform = gui.MakeTransformWithImage(e.Image, gui.OriginTopLeft)
	e.CalculateRect()

	e.initialized = true
}

func (e Panel) checkInitialized() {
	if !e.initialized {
		log.Fatal("Panel is not initialized; this may cause unexpected behaviour and as such is an error. Call element.Init() to fix this.")
	}
}

func (e Panel) DrawSelf() {
	e.checkInitialized()

	if !e.Visible {
		return
	}
	call := func(screen *ebiten.Image) {
		utils.DrawImageAtRect(screen, e.Image, e.Rect, &ebiten.DrawImageOptions{})
	}
	e.GetRoot().AddToDrawStack(e.ZIndex, call)
}

func (e Panel) Draw() {
	gui.Defaults.Draw(&e)
}

func (e *Panel) Update() {
	gui.Defaults.UpdateChildren(e)
}

func (e *Panel) CalculateRect() {
	e.Rect = gui.Defaults.CalculateRect(e)
}

func (e Panel) GetContainer() *gui.Container {
	return &e.Container
}

func (e *Panel) SetParent(parent *gui.Container) {
	e.Parent = parent
}

func MakePanel(image *ebiten.Image) Panel {
	elm := Panel{Image: image}
	elm.Init()

	elm.Visible = true

	return elm
}

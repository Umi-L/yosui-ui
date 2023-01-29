package gui

import (
	"github.com/umi-l/yosui-ui/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Element struct { //elements are just Containers with drawables
	Container

	Image *ebiten.Image

	initialized bool
}

func (e *Element) Init() {
	e.Transform = MakeTransformWithImage(e.Image, OriginTopLeft)
	e.calculateRect()

	e.initialized = true
}

func (e Element) checkInitialized() {
	if !e.initialized {
		log.Fatal("Element is not initialized; this may cause unexpected behaviour and as such is an error. Call element.Init() to fix this.")
	}
}

func (e Element) drawSelf() {
	e.checkInitialized()

	if !e.Visible {
		return
	}
	call := func(screen *ebiten.Image) {
		utils.DrawImageAtRect(screen, e.Image, e.Rect, &ebiten.DrawImageOptions{})
	}
	e.GetRoot().AddToDrawStack(e.ZIndex, call)
}

func (e Element) draw() {
	Defaults.Draw(&e)
}

func (e *Element) Update() {
	Defaults.UpdateChildren(e)
}

func (e *Element) calculateRect() {
	e.Rect = Defaults.CalculateRect(e)
}

func (e Element) GetContainer() *Container {
	return &e.Container
}

func (e *Element) SetParent(parent *Container) {
	e.Parent = parent
}

func MakeElement(image *ebiten.Image) Element {
	elm := Element{Image: image}
	elm.Init()

	elm.Visible = true

	return elm
}

type ElementInterface interface {
	drawSelf()
	draw()
	Update()
	calculateRect()
	SetParent(parent *Container)
	GetContainer() *Container
}

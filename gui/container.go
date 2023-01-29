package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/umi-l/yosui-ui/types"
)

type Container struct {
	Parent *Container
	Rect   Rect

	Transform Transform

	children []ElementInterface

	Visible bool
}

// function used to SET the transform of a container. Should ONLY be used during init or if you know what you're doing.
func (c *Container) SetTransform(t Transform) {
	c.Transform = t
}

// function used to UPDATE the transform of a container and all of its children. Should be used during runtime.
func (c *Container) UpdateTransform(t Transform) {
	c.Transform = t
	c.CalculateRect()

	for _, child := range c.children {
		child.CalculateRect()
	}
}

func (c *Container) AddChild(child ElementInterface) {
	c.children = append(c.children, child)
	child.SetParent(c)
	child.CalculateRect()
}

func (c *Container) SetParent(parent *Container) {
	c.Parent = parent
}

func (c Container) IsVisible() bool {

	if !c.Visible {
		return false
	}

	if c.Parent != nil {
		return c.Parent.IsVisible()
	}

	return true
}

func (c Container) Draw(screen *ebiten.Image) {}

func (c *Container) Update() {
	Defaults.UpdateChildren(c)
}

func (c Container) DrawTree(screen *ebiten.Image) {

	if !c.Visible {
		return
	}

	for _, child := range c.children {

		Draw(child, screen)
		child.DrawTree(screen)
	}
}

func (c Container) GetContainer() Container {
	return c
}

func (c *Container) CalculateRect() {
	c.Rect = Defaults.CalculateRect(c)
}

func NewContainer(Transform Transform, visible bool) Container {
	return Container{Transform: Transform, Visible: visible}
}

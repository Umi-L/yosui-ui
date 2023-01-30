package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/umi-l/yosui-ui/types"
	"log"
	"sort"
)

type Container struct {
	IsRoot    bool
	drawStack map[int][]DrawCall
	ZIndex    int

	Parent *Container
	Rect   Rect

	Transform Transform

	Children []ElementInterface

	Visible bool
}

// function used to SET the transform of a container. Should ONLY be used during init or if you know what you're doing.
func (c *Container) SetTransform(t Transform) {
	c.Transform = t
}

func (c *Container) AddToDrawStack(zIndex int, call DrawCall) {
	if !c.IsRoot {
		log.Fatal("Trying to add to drawstack of non root container.")
	}

	c.drawStack[zIndex] = append(c.drawStack[zIndex], call)
}

func (c *Container) GetRoot() *Container {

	if c.IsRoot {
		return c
	}

	return c.Parent.GetRoot()
}

// function used to UPDATE the transform of a container and all of its children. Should be used during runtime.
func (c *Container) UpdateTransform(t Transform) {
	c.Transform = t
	c.CalculateRect()

	for _, child := range c.Children {
		child.CalculateRect()
	}
}

func (c *Container) AddChild(child ElementInterface) {
	c.Children = append(c.Children, child)
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

func (c Container) DrawSelf() {}

func (c *Container) Update() {
	Defaults.UpdateChildren(c)
}

func (c *Container) Draw() {
	Defaults.Draw(c)
}

func (c *Container) DrawAsRoot(screen *ebiten.Image) {
	c.Draw()

	if !c.IsRoot {
		log.Fatal("Trying to Draw as root with non-root container")
	}

	//get keys in order
	keys := make([]int, len(c.drawStack))

	i := 0
	for k := range c.drawStack {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	//make Draw calls
	for j := range c.drawStack {
		calls := c.drawStack[keys[j]]
		for _, call := range calls {
			call(screen)
		}
	}

	//clear stack
	c.drawStack = make(map[int][]DrawCall)
}

func (c *Container) InitializeDrawStack() {
	if !c.IsRoot {
		log.Fatal("trying to initialize Draw stack on non-root container")
	}

	c.drawStack = make(map[int][]DrawCall)
}

func (c *Container) GetContainer() *Container {
	return c
}

func (c *Container) CalculateRect() {
	c.Rect = Defaults.CalculateRect(c)
}

func NewContainer(Transform Transform, visible bool) Container {
	return Container{Transform: Transform, Visible: visible}
}

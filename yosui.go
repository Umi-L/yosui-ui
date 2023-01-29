package yosui

import "github.com/umi-l/yosui-ui/gui"

func MakeRootContainer(width, height int) gui.Container {
	root := gui.NewContainer(gui.Transform{X: 0, Y: 0, W: float32(width), H: float32(height)}, true)

	return root
}

func MakeRelativeContainer(parent *gui.Container, transform gui.Transform, visible bool) *gui.Container {
	newContainer := gui.NewContainer(transform, visible)

	parent.AddChild(&newContainer)

	return &newContainer
}

package gui

import (
	"log"

	"github.com/umi-l/yosui-ui/types"
)

type defaultsT struct{} // unexported type

var Defaults defaultsT

func (_ defaultsT) CalculateRect(e ElementInterface) types.Rect {

	c := e.GetContainer()

	parentExists := true

	if c.Parent == nil {
		parentExists = false
		//log.Print("WARN: Container has no parent, cannot calculate rect relative to parent")
	}

	calcRect := types.Rect{}

	//X
	if c.Transform.XPercent == 0 { //check for default value, also cant be 0% as that devides by 0
		calcRect.X = c.Transform.X
	} else if parentExists {
		calcRect.X = float32(c.Parent.Rect.W) * c.Transform.XPercent
	} else {
		log.Fatal("Container has no parent, cannot calculate rect relative to parent")
	}

	//Y
	if c.Transform.YPercent == 0 { //check for default value, also cant be 0% as that devides by 0
		calcRect.Y = c.Transform.Y
	} else if parentExists {
		calcRect.Y = float32(c.Parent.Rect.H) * c.Transform.YPercent
	} else {
		log.Fatal("Container has no parent, cannot calculate rect relative to parent")
	}

	//W
	if c.Transform.WPercent == 0 { //check for default value, also cant be 0% as that devides by 0
		calcRect.W = c.Transform.W
	} else if parentExists {
		calcRect.W = float32(c.Parent.Rect.W) * c.Transform.WPercent
	} else {
		log.Fatal("Container has no parent, cannot calculate rect relative to parent")
	}

	//H
	if c.Transform.HPercent == 0 { //check for default value, also cant be 0% as that devides by 0
		calcRect.H = c.Transform.H
	} else if parentExists {
		calcRect.H = float32(c.Parent.Rect.H) * c.Transform.HPercent
	} else {
		log.Fatal("Container has no parent, cannot calculate rect relative to parent")
	}

	//Origin
	c.Transform.Origin = CalculateOriginFromRect(calcRect, c.Transform.Origin.OriginIndex)

	calcRect.X -= c.Transform.Origin.X
	calcRect.Y -= c.Transform.Origin.Y

	for _, child := range c.children {
		child.CalculateRect()
	}

	return calcRect
}

func (_ defaultsT) UpdateChildren(e ElementInterface) {

	for _, child := range e.GetContainer().children {
		child.Update()
	}
}

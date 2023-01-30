package gui

type ElementInterface interface {
	//internal ui methods
	DrawSelf()
	Draw()
	CalculateRect()

	//external methods
	Update()
	SetParent(parent *Container)
	GetContainer() *Container
}

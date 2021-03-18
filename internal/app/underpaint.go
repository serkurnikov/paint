package app

//https://github.com/golang-standards/project-layout

const (
	BasePath       = "C:\\Users\\master\\go\\src\\projects\\paint\\assets\\examples\\"
	TestPathImage0 = BasePath + "test0.jpg"
	TestPathImage1 = BasePath + "test1.jpg"

	MeanShiftPathImage = BasePath + "meanShift.jpg"
	ThresholdPathImage = BasePath + "threshold.jpg"
	WatershedPathImage = BasePath + "watershed.jpg"

	PalettePathImage = BasePath + "palette.jpg"
)

func (a App) ExternalApiTest() {}
func (a App) Scobel()          {}

func (a App) Render(ctx Ctx) {
	//_ = a.create.Publish("1")
}

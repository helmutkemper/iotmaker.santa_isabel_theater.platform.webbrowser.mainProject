// +build js

package main

import (
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.platform/abstractType/basicBox"
	"github.com/helmutkemper/iotmaker.platform/abstractType/colornames"
	"github.com/helmutkemper/iotmaker.platform/abstractType/gradient"
	"github.com/helmutkemper/iotmaker.platform/abstractType/shadow"
)

func main() {

	done := make(chan struct{}, 0)

	var density = 1.0
	var densityManager coordinateManager.IDensity = &coordinateManager.Density{}

	browserDocument := document.NewDocument()

	stage := canvas.NewStage(
		browserDocument,
		"stage",
		300,
		300,
		density,
		densityManager,
	)

	colorDarkBlue := colornames.Darkblue
	shadowFilter := shadow.NewShadowFilter(colorDarkBlue, 5, 2, 2, density, densityManager)

	colorWhite := gradient.NewColorPosition(colornames.WhiteTransparent, 0.2)
	colorBlack := gradient.NewColorPosition(colornames.Black, 1)
	colorList := gradient.NewColorList(colorBlack, colorWhite)

	coordinateP0 := gradient.NewPoint(0, 0, density, densityManager)
	coordinateP1 := gradient.NewPoint(120, 150, density, densityManager)
	gradientFilter := gradient.NewGradientLinearToStroke(coordinateP0, coordinateP1, colorList)

	bx := basicBox.NewBasicBox(&stage.Canvas, &stage.ScratchPad, "bbox_1", 20, 50, 100, 100, 5, 5, nil, nil, density, densityManager)
	basicBox.NewBasicBox(&stage.Canvas, &stage.ScratchPad, "bbox_2", 20+50, 50+50, 100, 100, 10, 8, shadowFilter, gradientFilter, density, densityManager)

	//fmt.Printf("over: %v\n", bx.GetAlphaChannel(0, 100))

	/*abstractType.NewLineTo(
		&stage.Canvas,
		"line",
		1.0,
		0,
		0,
		300,
		300,
		1,
	)*/

	mouse.AddFunctionPointer(bx.GetAlphaChannel)
	browserDocument.SetMouseMoveListener(mouse.GetDefaultFunction())

	<-done
}

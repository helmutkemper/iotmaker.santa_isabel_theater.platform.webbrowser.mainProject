// +build js

//
package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserStage"
	webBrowserMouse "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/fps"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
	"time"
)

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	stage          *canvas.Stage
)

var htmlElement iotmakerPlatformIDraw.IHtml
var browserDocument document.Document

func prepareDataBeforeRun() {
	htmlElement = &Html.Html{}

	browserDocument = factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(
		htmlElement,
		browserDocument,
		"stage",
		density,
		densityManager,
	)
	stage.SetCursor(mouse.KCursorDefault)
}

func main() {

	done := make(chan struct{}, 0)
	prepareDataBeforeRun()
	fps.Set(60)
	//htmlElement.Remove(browserDocument.SelfDocument, imgSpace.Get())

	browserDocument.AddEventListener(eventMouse.KMouseMove, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	browserDocument.AddEventListener(eventMouse.KClick, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	browserDocument.AddEventListener(eventMouse.KDblClick, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	browserDocument.AddEventListener(eventMouse.KMouseDown, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	browserDocument.AddEventListener(eventMouse.KMouseUp, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	browserDocument.AddEventListener(eventMouse.KContextMenu, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	//mouse.AddFunctionPointer("bBox2", bx2.GetCollisionBox, bateu)

	ticker := time.NewTicker(time.Second * 3)
	show := true
	go func() {
		for {
			select {
			case <-ticker.C:

				show = !show
				if show == true {
					stage.CursorShow()
				} else {
					stage.CursorHide()
				}
			}
		}
	}()

	<-done
}

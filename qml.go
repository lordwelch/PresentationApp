// PresentationApp project qml.go
package main

import (
	"fmt"
	"image"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/lordwelch/qml"
)

var (
	selCell    int        //the focused and
	rhtClkCell int        //the cell that was last right clicked
	qimg       qml.Object //file for the image object
	cellQml    qml.Object //file for the cell object
	mainQml    qml.Object //main QML file
	edtQml     qml.Object
	textEdit   qml.Object
	window     *qml.Window
	window2    *qml.Window
	engine     *qml.Engine
	quickEdit  Bool = false
	imgready   Bool = false
	QML        *qmlVar
)

func initQML() {
	window2.ObjectByName("textClrDialog").On("accepted", func() {
		window2.ObjectByName("textClrDialog").Color("color")
	})
}

func (qv *qmlVar) Changed() {
	qml.Changed(qv, qv.VerseLen)
	qml.Changed(qv, qv.OrderLen)
	qml.Changed(qv, qv.ImgLen)
	qml.Changed(qv, qv.FontLen)
}

//signals for the cell and image in qml
func (cl *cell) setSignal() {
	cl.qmlcell.ObjectByName("cellMouse").On("clicked", func(musEvent qml.Object) {
		btn := musEvent.Property("button")
		//right click
		if btn == 2 {
			//context menu
			window.ObjectByName("mnuCtx").Call("popup")
			rhtClkCell = cl.index
		} else {
			//left click
			//select and update image preview for cell
			selCell = cl.qmlcell.Int("index")
			cl.qmlcell.ObjectByName("cellMouse").Set("focus", true)
			setupScene()
		}
		//update image preview
		cl.clearcache()
	})

	cl.qmlimg.ObjectByName("cellMouse").On("clicked", func(musEvent qml.Object) {
		btn := musEvent.Property("button")
		//right click
		if btn == 2 {
			//context menu
			window.ObjectByName("mnuCtx").Call("popup")
			rhtClkCell = cl.index
		} else {
			//left click
			//select and update image preview for cell
			selCell = cl.qmlcell.Int("index")
			cl.qmlcell.ObjectByName("cellMouse").Set("focus", true)
			setupScene()
		}
		//update image preview
		cl.clearcache()
	})
	cl.qmlcell.ObjectByName("cellMouse").On("focusChanged", func(focus bool) {
		if focus {
			cl.qmlcell.ObjectByName("cellMouse").Call("selected")
		} else {
			cl.qmlcell.ObjectByName("cellMouse").Call("notSelected")
		}
	})

	cl.qmlcell.ObjectByName("cellMouse").On("doubleClicked", func() {
		if quickEdit {
			//cover the cell with the text edit
			textEdit.Set("cell", cl.index)
			textEdit.Set("x", cl.qmlcell.Int("x"))
			textEdit.Set("y", cl.qmlcell.Int("y"))
			textEdit.Set("height", cl.qmlcell.Int("height"))
			textEdit.Set("z", 100)
			textEdit.Set("visible", true)
			textEdit.ObjectByName("textEdit1").Set("focus", true)
			textEdit.Set("enabled", true)

			//set the text
			textEdit.ObjectByName("textEdit1").Set("text", cl.text)
		}
	})

}

//setSignals() for non dynamic elements
func setSignals() {
	window.ObjectByName("imgpicker").On("accepted", func() {
		//delete file://  from url
		url := filepath.Clean(strings.Replace(window.ObjectByName("imgpicker").String("fileUrl"), "file:", "", 1))

		//replace new image
		slides[rhtClkCell].img.Clear()
		slides[rhtClkCell].img.ReadImage(url)
		setupScene()
		//update image preview
		slides[rhtClkCell].clearcache()
	})

	window.ObjectByName("btnAdd").On("clicked", func() {
		slides.add()
	})

	window.ObjectByName("btnRem").On("clicked", func() {
		slides[len(slides)-1].remove()
	})

	window.ObjectByName("btnMem").On("clicked", func() {
		//run GC
		debug.FreeOSMemory()
	})

	window.On("closing", func() {
		//close glfw first
		if false == window.Property("cls") {
			win.SetShouldClose(true)
			window.Set("cls", true)
		}

	})

	window.ObjectByName("mnuDisplay").On("triggered", func() {
		win.SetShouldClose(false)
		window.Set("cls", false)
		win.Show()
		qml.ResetGLFW()
	})

	window.ObjectByName("mnuEdit").On("triggered", func() {
		(&quickEdit).Toggle()
	})

	textEdit.ObjectByName("textEdit1").On("focusChanged", func(focus bool) {
		var (
			str string
			cel *cell
		)

		if !focus {
			//set text back to the cell
			str = textEdit.ObjectByName("textEdit1").String("text")
			cel = slides[textEdit.Int("cell")]
			if textEdit.Bool("txt") {
				cel.qmlcell.ObjectByName("cellText").Set("text", str)
				cel.text = str
			}
		}
	})
}

func edtQmlShow() {
	slc := window2.ObjectByName("fontPicker").Property("model")
	fmt.Println(slc)
}

//imgProvider() for preview images in QML
func imgProvider(id string, width, height int) image.Image {
	if imgready && (len(id) > 0) {
		//fmt.Println("source (provider): ", id)
		i1 := strings.Index(id, `;`)
		i, _ := strconv.Atoi(id[:i1])
		return slides[i].getImage(width, height)

	}
	var img1 image.Image = image.NewRGBA(image.Rect(0, 0, 340, 480))
	return img1

}

//clear cache dosen't actually clear the cache
//just gives a new source so that the cache isn't used
func (cl *cell) clearcache() {
	str := cl.qmlimg.String("source")
	//fmt.Println("source (click): ", str)
	i := strings.Index(str, `;`)
	str1 := str[:i]
	//fmt.Println("ext (click): ", str1)
	i1, _ := strconv.Atoi(str[i+1:])
	str = str1 + `;` + strconv.Itoa(i1+1)
	//fmt.Println("new source (click): ", str)
	cl.qmlimg.Set("source", str)
}

// PresentationApp project qml.go
package main

import (
	"image"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"

	"gopkg.in/qml.v1"
)

var (
	cellQml        qml.Object //file for the cell object
	displayImg     qml.Object
	displayQml     qml.Object
	DisplayWindow  *qml.Window
	editQml        qml.Object
	engine         *qml.Engine
	imgQml         qml.Object //file for the image object
	imgready       Bool       = false
	MainWindow     *qml.Window
	mainQml        qml.Object //main QML file
	QML            *qmlVar    //misc var qml needs
	quickEdit      Bool       = false
	rightClickCell int        //the cell that was last right clicked
	selectedCell   int        //the focused cell
	serviceObject  qml.Object
	serviceQml     qml.Object
	songEditWindow *qml.Window
	textEdit       qml.Object
)

func qmlWindows() error {
	mainQml, err = engine.LoadFile(path + "/Main.qml")
	if err != nil {
		return err
	}

	displayQml, err = engine.LoadFile(path + "/Display.qml")
	if err != nil {
		return err
	}

	editQml, err = engine.LoadFile(path + "/SongEdit.qml")
	if err != nil {
		return err
	}

	cellQml, err = engine.LoadFile(path + "/Cell.qml")
	if err != nil {
		return err
	}

	serviceQml, err = engine.LoadFile(path + "/Service.qml")
	if err != nil {
		return err
	}

	MainWindow = mainQml.CreateWindow(engine.Context())
	songEditWindow = editQml.CreateWindow(engine.Context())
	DisplayWindow = displayQml.CreateWindow(engine.Context())

	textEdit = MainWindow.ObjectByName("textEdit")
	return nil
}

func showWindows() {
	MainWindow.Show()
	songEditWindow.Show()
	DisplayWindow.Show()
}

//imgProvider() for preview images in QML
func imgProvider(id string, width, height int) image.Image {
	var img1 image.Image
	if imgready && (len(id) > 0) {
		//fmt.Println("source (provider): ", id)
		i1 := strings.Index(id, `;`)
		i, _ := strconv.Atoi(id[i1+1:])
		img1 = slides[i].getImage(width, height)

	} else {
		img1 = image.NewRGBA(image.Rect(0, 0, 340, 480))
	}

	return img1

}

func edtQmlShow() {
	//slc := window2.ObjectByName("fontPicker").Property("model")
	//fmt.Println(slc)
}

//setSignals() for non dynamic elements
func setSignals() {
	MainWindow.ObjectByName("imgpicker").On("accepted", func() {
		//delete "file://"  from url
		url := filepath.Clean(strings.Replace(MainWindow.ObjectByName("imgpicker").String("fileUrl"), "file:", "", 1))

		//replace new image
		slides[rightClickCell].image.img.Clear()
		slides[rightClickCell].image.img.ReadImage(url)
	})

	MainWindow.ObjectByName("btnAdd").On("clicked", func() {
		slides.add()
	})

	MainWindow.ObjectByName("btnRem").On("clicked", func() {
		slides.remove(len(slides) - 1)
	})

	MainWindow.ObjectByName("btnMem").On("clicked", func() {
		//run GC
		debug.FreeOSMemory()
	})

	MainWindow.ObjectByName("mnuEdit").On("triggered", func() {
		quickEdit.Flip()
	})

	textEdit.ObjectByName("textEdit1").On("focusChanged", func(focus bool) {
		var (
			str string
			cel *Cell
		)

		if !focus {
			//set text back to the cell
			str = textEdit.ObjectByName("textEdit1").String("text")
			cel = slides[textEdit.Int("cell")]
			if textEdit.Bool("txt") {
				cel.qmlObject.ObjectByName("cellText").Set("text", str)
				cel.text = str
			}
		}
	})
}

//signals for the cell and image in qml
func (cl *Cell) setSignal() {
	cl.qmlObject.ObjectByName("cellMouse").On("clicked", func(mouseEvent qml.Object) {
		btn := mouseEvent.Property("button")
		//right click
		if btn == 2 {
			//context menu
			MainWindow.ObjectByName("mnuCtx").Call("popup")
			rightClickCell = cl.index
		} else {
			//left click
			cl.Select()

		}

	})

	cl.image.qmlImage.ObjectByName("cellMouse").On("clicked", func(mouseEvent qml.Object) {
		btn := mouseEvent.Property("button")
		//right click
		if btn == 2 {
			//context menu
			MainWindow.ObjectByName("mnuCtx").Call("popup")
			rightClickCell = cl.index
		} else {
			//left click
			cl.Select()
		}
	})

	cl.qmlObject.ObjectByName("cellMouse").On("focusChanged", func(focus bool) {
		if focus {
			cl.qmlObject.ObjectByName("cellMouse").Call("selected")
		} else {
			cl.qmlObject.ObjectByName("cellMouse").Call("notSelected")
		}
	})

	cl.qmlObject.ObjectByName("cellMouse").On("doubleClicked", func() {
		if quickEdit {
			//cover the cell with the text edit
			textEdit.Set("cell", cl.index)
			textEdit.Set("x", cl.qmlObject.Int("x"))
			textEdit.Set("y", cl.qmlObject.Int("y"))
			textEdit.Set("height", cl.qmlObject.Int("height"))
			textEdit.Set("z", 100)
			textEdit.Set("visible", true)
			textEdit.ObjectByName("textEdit1").Set("focus", true)
			textEdit.Set("enabled", true)

			//set the text
			textEdit.ObjectByName("textEdit1").Set("text", cl.text)
		}
	})

}

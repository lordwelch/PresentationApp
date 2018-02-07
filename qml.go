// PresentationApp project qml.go
package main

import (
	"image"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/limetext/qml-go"
)

var (
	selectedCell   int        //the focused cell
	rightClickCell int        //the cell that was last right clicked
	cellQml        qml.Object //file for the cell object
	mainQml        qml.Object //main QML file
	editQml        qml.Object
	textEdit       qml.Object
	displayQml     qml.Object
	displayImg     qml.Object
	DisplayWindow  *qml.Window
	MainWindow     *qml.Window
	songEditWindow *qml.Window
	serviceObject  qml.Object
	serviceQml     qml.Object
	engine         *qml.Engine
	quickEdit      bool = false
	imgready       bool = false
	QML            *qmlVar
)

func initQML() {
	/*window2.ObjectByName("textClrDialog").On("accepted", func() {
		window2.ObjectByName("textClrDialog").Color("color")
	})*/
}

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

/*func (qv *qmlVar) Changed() {
	qml.Changed(qv, qv.VerseLen)
	qml.Changed(qv, qv.OrderLen)
	qml.Changed(qv, qv.ImgLen)
	qml.Changed(qv, qv.FontLen)
}*/

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
			//select and update image preview for cell
			cl.Select()
		}
		//update image preview
		cl.clearcache()
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
			//select and update image preview for cell
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

//setSignals() for non dynamic elements
func setSignals() {
	MainWindow.ObjectByName("imgpicker").On("accepted", func() {
		//delete file://  from url
		url := filepath.Clean(strings.TrimPrefix(MainWindow.ObjectByName("imgpicker").String("fileUrl"), "file:"))

		//replace new image
		slides[rightClickCell].image.img.Clear()
		slides[rightClickCell].image.img.ReadImage(url)
	})

	MainWindow.ObjectByName("btnAdd").On("clicked", func() {
		slides.add("not")
	})

	MainWindow.ObjectByName("btnRem").On("clicked", func() {
		slides.remove(len(slides) - 1)
	})

	MainWindow.ObjectByName("btnMem").On("clicked", func() {
		//run GC
		debug.FreeOSMemory()
	})

	MainWindow.ObjectByName("mnuEdit").On("triggered", func() {
		quickEdit = !quickEdit
	})

	textEdit.ObjectByName("textEdit1").On("focusChanged", func(focus bool) {
		var (
			str  string
			cell *Cell
		)

		if !focus {
			//set text back to the cell
			str = textEdit.ObjectByName("textEdit1").String("text")
			cell = slides[textEdit.Int("cell")]
			if textEdit.Bool("txt") {
				cell.qmlObject.ObjectByName("cellText").Set("text", str)
				cell.text = str
			}
		}
	})
}

func edtQmlShow() {
	//slc := window2.ObjectByName("fontPicker").Property("model")
	//fmt.Println(slc)
}

//imgProvider() for preview images in QML
func imgProvider(id string, width, height int) image.Image {
	var img1 image.Image
	if imgready && (len(id) > 0) {
		//fmt.Println("source (provider): ", id)
		i1 := strings.Index(id, `;`)
		i, _ := strconv.Atoi(id[:i1])
		img1 = slides[i].getImage(width, height)
	} else {
		img1 = image.NewRGBA(image.Rect(0, 0, 340, 480))
	}
	return img1

}

//clear cache dosen't actually clear the cache
//just gives a new source so that the cache isn't used
func (cl *Cell) clearcache() {
	str := cl.image.qmlImage.String("source")
	i := strings.Index(str, `;`)
	str1 := str[:i]
	i1, _ := strconv.Atoi(str[i+1:])
	str = str1 + `;` + strconv.Itoa(i1+1)
	//fmt.Println("new source (click): ", str)
	cl.image.qmlImage.Set("source", str)
}

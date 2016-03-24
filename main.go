// PresentationApp project main.go
package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/kardianos/osext"
	"github.com/lordwelch/qml"
	"gopkg.in/gographics/imagick.v2/imagick"
)

//lazy wanted a toggle function
type Bool bool

type cell struct {
	text    string
	img     *imagick.MagickWand
	qmlimg  qml.Object
	qmlcell qml.Object
	index   int
}
type slide []*cell

var (
	//displayed width/height the focused and the cell that was last right clicked
	monWidth, monHeight, selCell, rhtClkCell int
	path                                     string
	qimg                                     qml.Object //file for the image object
	textEdit                                 qml.Object
	cellQml                                  qml.Object   //file for the cell object
	window                                   *qml.Window  //QML
	win                                      *glfw.Window //GLFW
	slides                                   slide
	err                                      error
	monitors                                 []*glfw.Monitor
	projMonitor                              *glfw.Monitor
	tex1                                     *uint32                //identifier for opengl texture
	texDel, quickEdit                        Bool    = false, false //if texture should be deleted
)

func main() {
	selCell = 0

	if err = qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

}

func run() error {
	var mainQml qml.Object
	imagick.Initialize()

	engine := qml.NewEngine()
	engine.AddImageProvider("images", imgProvider)
	//path for qml files TODO: change to somewhere else
	path, err = osext.ExecutableFolder()
	path = filepath.Clean(path + "/../src/github.com/lordwelch/PresentationApp/")

	mainQml, err = engine.LoadFile(path + "/main.qml")
	if err != nil {
		return err
	}

	cellQml, err = engine.LoadFile(path + "/qml/cell.qml")
	if err != nil {
		return err
	}

	qimg, err = engine.LoadFile(path + "/qml/img.qml")
	if err != nil {
		return err
	}

	window = mainQml.CreateWindow(nil)

	textEdit = window.ObjectByName("textEdit")
	slides.addCell()
	//signals for whole qml
	setSignals()
	//image is ready for imageprovider
	imgready = true

	window.Show()
	qml.RunMain(glInit)

	window.Wait()

	imagick.Terminate()
	return nil
}

func setupScene() {

	gl.ClearColor(0, 0, 0, 0)
	if texDel {
		gl.DeleteTextures(1, tex1)
	}
	tex1 = newTexture(*slides[selCell].getImage(monWidth, monHeight))

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-1, 1, -1, 1, 1.0, 10.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	texDel = true

}

func drawSlide() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0, 0, -3.0)

	gl.Begin(gl.QUADS)

	//top left
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, 1, 0)
	//top right
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, 1, 0)

	//bottom right
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, -1, 0)

	//bottom left
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, -1, 0)

	gl.End()

}

func newTexture(rgba image.RGBA) *uint32 {
	var texture1 uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture1)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return &texture1

}

func checkMon() {
	monitors = glfw.GetMonitors()
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.AutoIconify, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)

	if i := len(monitors); i < 2 {
		fmt.Println("You only have 1 monitor!!!!!!!!!!! :-P")
		monWidth = 800
		monHeight = 600

		win, err = glfw.CreateWindow(monWidth, monHeight, "Cube", nil, nil)
		if err != nil {
			panic(err)
		}
		projMonitor = monitors[0]
	} else {
		fmt.Printf("You have %d monitors\n", i)
		monWidth = monitors[1].GetVideoMode().Width
		monHeight = monitors[1].GetVideoMode().Height
		win, err = glfw.CreateWindow(monWidth, monHeight, "Cube", nil, nil)
		fmt.Println(win.GetPos())
		win.SetPos(monitors[1].GetPos())
		fmt.Println(monWidth, monHeight)
		if err != nil {
			panic(err)
		}
		projMonitor = monitors[1]

	}
	monitorInfo()

}

func monitorInfo() {
	for _, mon := range monitors {
		fmt.Printf("monitor name: %s\n", mon.GetName())
		i, t := mon.GetPos()
		fmt.Printf("position X: %d  Y: %d\n", i, t)

	}

}

func glInit() {
	if err = glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	checkMon()

	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	win.SetPos(projMonitor.GetPos())
	setupScene()

	qml.Func1 = func() int {
		if !win.ShouldClose() {
			glfw.PollEvents()
			drawSlide()
			win.SwapBuffers()
			return 0

		} else {
			win.Hide()
			//win.Destroy()
			glfw.Terminate()
			return 1

		}
	}
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
		slides.addCell()
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

//getImage() from imagick to image.RGBA
func (cl cell) getImage(width, height int) (img *image.RGBA) {
	mw := cl.img.GetImage()
	if (width == 0) || (height == 0) {
		width = int(mw.GetImageWidth())
		height = int(mw.GetImageHeight())
	}

	mw = resizeImage(mw, width, height, true, true)
	img = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	if img.Stride != img.Rect.Size().X*4 {
		panic("unsupported stride")
	}

	Tpix, _ := mw.ExportImagePixels(0, 0, uint(width), uint(height), "RGBA", imagick.PIXEL_CHAR)
	img.Pix = Tpix.([]uint8)
	mw.Destroy()
	return
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

//Adds a new cell
func (sl *slide) addCell( /*cl *cell*/ ) {
	var cl cell
	//gets the length so that the index is valid
	cl.index = len(*sl)

	//increase count on parent QML element
	window.ObjectByName("gridRect").Set("count", window.ObjectByName("gridRect").Int("count")+1)
	cl.qmlcell = cellQml.Create(nil)
	cl.qmlcell.Set("objectName", fmt.Sprintf("cellRect%d", len(*sl)))
	cl.qmlcell.Set("parent", window.ObjectByName("data1"))
	cl.qmlcell.Set("index", cl.index)

	//load image
	cl.img = imagick.NewMagickWand()
	cl.img.ReadImage("logo:")

	//give QML the text
	cl.qmlcell.ObjectByName("cellText").Set("text", cl.text)

	//keep the pointer/dereference (i'm not sure which it is)
	//problems occur otherwise
	*sl = append(*sl, &cl)
	cl.setSignal()

	//seperate image object in QML
	cl.qmlimg = qimg.Create(nil)
	fmt.Println("index", cl.index)
	fmt.Printf("objectName: %s\n", fmt.Sprintf("cellImg%d", cl.index))
	cl.qmlimg.Set("objectName", fmt.Sprintf("cellImg%d", cl.index))
	cl.qmlimg.Set("source", fmt.Sprintf("image://images/%d"+`;`+"0", cl.index))
	fmt.Println("source: ", cl.qmlimg.String("source"))
	cl.qmlimg.Set("parent", window.ObjectByName("data2"))
	cl.qmlimg.Set("index", cl.index)

}

//remove() should destroy everything
func (cl *cell) remove() {
	cl.text = ""
	cl.qmlimg.Destroy()
	cl.qmlcell.Destroy()
	cl.img.Destroy()
	window.ObjectByName("gridRect").Set("count", window.ObjectByName("gridRect").Int("count")-1)
	slides.remove(cl.index)
	cl.index = -1

}

//(slide) remove copied from gist on github
func (sl *slide) remove(i int) {
	*sl, (*sl)[len((*sl))-1] = append((*sl)[:i], (*sl)[i+1:]...), nil
}

//lazy wanted a toggle func
func (bl *Bool) Toggle() {
	if *bl == false {
		*bl = true
	} else {
		*bl = false
	}
}

//not really needed
func (cl cell) String() string {
	return fmt.Sprintf("Index: %d \nText:  %s\n", cl.index, cl.text)
}

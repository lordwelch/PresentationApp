// PresentationApp project glfw.go
package main

import (
	"fmt"
	"image"
	"log"
    "os"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/lordwelch/qml"
    "github.com/go-gl/gltext"
)

var (
	win         *glfw.Window
	monWidth    int //displayed height
	monHeight   int //displayed width
	monitors    []*glfw.Monitor
	projMonitor *glfw.Monitor
	tex1        *uint32 //identifier for opengl texture
	texDel      Bool    //if texture should be deleted
    flt *os.File
    f *gltext.Font
)

func loadfont(){

    flt,err = os.Open("Comfortaa-Regular.ttf")
    if err != nil {
        f, err = gltext.LoadTruetype(flt, 30, 32, 255, gltext.LeftToRight)
    }
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
    f.Printf(5,10,"test")

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
		win.SetPos(monitors[1].GetPos())
		fmt.Printf("Width: %d  Height: %d \n", monWidth, monHeight)
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
	window.Set("cls", false)
	if err = glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	checkMon()

	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}
    loadfont()

	win.SetPos(projMonitor.GetPos())
	setupScene()

	qml.Func1 = func() int {
		if !win.ShouldClose() {
			//glfw.PollEvents()
			drawSlide()
			win.SwapBuffers()
			return 0

		}
		win.Hide()
		//win.Destroy()
		//glfw.Terminate()
		return 1

	}
}

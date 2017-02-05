// PresentationApp project glfw.go
package main

import (
	"fmt"

	"github.com/go-gl/glfw/v3.1/glfw"
)

var (
	monitorHeight    int // displayed width
	monitors         []*glfw.Monitor
	monitorWidth     int // displayed height
	projectorMonitor *glfw.Monitor
)

func checkMon() {
	monitors = glfw.GetMonitors()

	if i := len(monitors); i < 2 {
		fmt.Println("You only have 1 monitor!!!!!!!!!!! :-P")
		monitorWidth = 800
		monitorHeight = 600

		projectorMonitor = monitors[0]
	} else {
		fmt.Printf("You have %d monitors\n", i)
		monitorWidth = monitors[1].GetVideoMode().Width
		monitorHeight = monitors[1].GetVideoMode().Height
		projectorMonitor = monitors[1]

	}
	monitorInfo()

}

func monitorInfo() {
	fmt.Println(len(monitors))
	for _, mon := range monitors {
		fmt.Printf("Monitor name: %s\n", mon.GetName())
		x, y := mon.GetPos()
		fmt.Printf("Position: %v, %v\n", x, y)
		fmt.Printf("Size: %v x %v\n", mon.GetVideoMode().Width, mon.GetVideoMode().Height)

	}

}

func glInit() {
	if err = glfw.Init(); err == nil {
		checkMon()
		DisplayWindow.Root().Set("height", monitorHeight)
		DisplayWindow.Root().Set("width", monitorWidth)
		DisplayWindow.Root().Set("x", 0)
		DisplayWindow.Root().Set("y", 0)
	}
}

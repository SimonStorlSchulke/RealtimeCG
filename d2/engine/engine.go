package engine

import (
	"fmt"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread() //in main()?
}

type App struct {
	Name                 string
	Width, Height        int
	Resizable            bool
	window               *glfw.Window
	vertexShaderSource   string
	fragmentShaderSource string
}

func Start(name string, width, height int, resizable bool) {
	win := initGlfw(name, width, height, resizable)
	for !win.ShouldClose() {
		fmt.Println("ayy")
	}
}

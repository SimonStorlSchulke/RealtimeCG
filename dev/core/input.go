package core

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func initInput(w *glfw.Window, cb glfw.KeyCallback) {
	w.SetKeyCallback(cb)
}

func cbKey() bool {
	return true
}

//CalcMousePos Returns the Mouse Position on the window on a scale of
//0 (top left) to 1 (bottom right)
func CalcMousePos(w *glfw.Window) {
	w.SetCursorPosCallback(cbCursor)
}

//CursorPositions
var MouseX, MouseY float64

//Cursor Callback
func cbCursor(window *glfw.Window, xpos, ypos float64) {
	width, height := window.GetSize()
	MouseX = xpos / float64(width)
	MouseY = ypos / float64(height)
}

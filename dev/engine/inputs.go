package engine

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

//CursorPositions
var MouseX, MouseY float32

//Wap with keys (true when pressed)
var Keys map[string]bool

//CalcMousePos Returns the Mouse Position on the window on a scale of
//0 (top left) to 1 (bottom right)
func CalcMousePos(w *glfw.Window) {
	w.SetCursorPosCallback(cbCursor)
}

func cbCursor(window *glfw.Window, xpos, ypos float64) {
	width, height := window.GetSize()
	MouseX = float32(xpos) / float32(width)
	MouseY = float32(ypos) / float32(height)
}

func ReadKeys(w *glfw.Window) {
	w.SetKeyCallback(cbKeys)
}

func cbKeys(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

}

package engine

import (
	"fmt"
	"strconv"

	"github.com/go-gl/glfw/v3.2/glfw"
)

//CursorPositions
var MouseX, MouseY float32

//Wap with keys (true when pressed)
var NumKey = 1

var Pressed bool
var TimeR bool
var Rmb bool
var Orth bool = true

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
	w.SetKeyCallback(spaceCb)
	w.SetMouseButtonCallback(lmbCb)
	w.SetCharModsCallback(charCB)
}

func spaceCb(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeySpace && action == glfw.Press {
		//ugly - refactor me!
		Pressed = true
	}
}

func lmbCb(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if button == glfw.MouseButtonRight && action == glfw.Press {
		Rmb = true
	} else {
		Rmb = false
	}
}

func charCB(w *glfw.Window, char rune, mods glfw.ModifierKey) {
	s, err := strconv.Unquote(strconv.QuoteRune(char))
	if err != nil {
		return
	}
	if s == "d" {
		Orth = !Orth
		return
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return
	}
	if mods == glfw.ModAlt {
		num += 10
	}
	NumKey = num
	fmt.Println("layer", NumKey)
	TimeR = true
}

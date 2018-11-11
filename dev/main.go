package main

import (
	"RealtimeCG/dev/core"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 500
	height = 500
)

var (
	vertexShaderSource   string
	fragmentShaderSource string
	triangle             = []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
	time float32
)

func main() {

	window := core.InitGlfw(width, height, false, "Testwindow")
	defer glfw.Terminate()
	program := initOpenGL()

	vao := core.MakeVao(triangle)
	for !window.ShouldClose() {
		draw(vao, window, program)
	}
}

func draw(vao uint32, window *glfw.Window, program uint32) {

	time += 0.1
	core.SetSingleUniformFloat(program, "time", time)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

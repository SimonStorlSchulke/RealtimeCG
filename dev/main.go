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
	triangle2 = []float32{
		0.2, -0.2, 0,
		0.2, -0.2, 0,
		-0.2, -0.2, 0,
	}
	time float32
)

func main() {
	//triangle, _ = obj.ShittyObjReader("cube.obj")

	window := core.InitGlfw(width, height, false, "Testwindow")
	defer glfw.Terminate()
	prog := initOpenGL()

	vao := core.MakeVao(triangle)
	for !window.ShouldClose() {
		draw(vao, window, prog)
	}
}

//called every frame
func draw(vao uint32, window *glfw.Window, prog uint32) {

	time += 0.03
	core.SetSingleUniformFloat(prog, "time", time)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

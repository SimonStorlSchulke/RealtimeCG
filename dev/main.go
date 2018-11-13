package main

import (
	"RealtimeCG/dev/core"
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
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

	window := core.InitGlfw(width, height, false, "Testwindow")
	defer glfw.Terminate()
	prog := initOpenGL()

	vao := core.MakeVao(triangle)
	for !window.ShouldClose() {
		draw(vao, window, prog)
	}

	//unused - Projection Stuff
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(width)/height, 0.1, 10.0)
	core.SetUniformMat4(prog, "projection", projection)

	camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	fmt.Println(camera)
}

//called every frame
func draw(vao uint32, window *glfw.Window, prog uint32) {

	time += 0.03
	core.SetUniformFloat(prog, "time", time)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

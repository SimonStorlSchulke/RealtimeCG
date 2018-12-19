package main

import (
	"RealtimeCG/dev/core"
	"RealtimeCG/dev/engine"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	width  = 1000
	height = 750
)

var (
	vertexShaderSource   string
	fragmentShaderSource string
	time                 float32

	model      mgl32.Mat4
	projection mgl32.Mat4
)

func main() {

	window := core.InitGlfw(width, height, false, "Testwindow")
	defer glfw.Terminate()
	prog := initOpenGL()
	gl.UseProgram(prog)
	projection = mgl32.Perspective(mgl32.DegToRad(45.0), float32(width)/height, 0.1, 10.0)
	core.SetUniform(prog, "projection", projection)

	camera := engine.Cam(-3, -3, -3, 0, 0, 0)
	core.SetUniform(prog, "camera", camera.Mat())

	elementBuffer := core.ElementBuffer(engine.CubeVerts, engine.CubeIndices)
	for !window.ShouldClose() {
		update(elementBuffer, window, prog)
	}
}

//called every frame
func update(elementBuffer uint32, window *glfw.Window, prog uint32) {

	model = mgl32.Ident4()
	core.SetUniform(prog, "model", model) //replace with &model later

	time += 0.01
	core.SetUniform(prog, "time", time)

	camera := engine.Cam(-engine.MouseX*4, -engine.MouseY*4, -3, 0, 0, 0)
	core.SetUniform(prog, "camera", camera.Mat())

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.DrawElements(
		gl.TRIANGLES,                     // mode
		int32(len(engine.CubeIndices)*4), // count
		gl.UNSIGNED_INT,                  // type
		gl.PtrOffset(0),                  // element array buffer offset
	)

	engine.CalcMousePos(window)

	glfw.PollEvents()
	window.SwapBuffers()
}

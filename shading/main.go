package main

import (
	"RealtimeCG/shading/core"
	"RealtimeCG/shading/engine"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	width  = 900
	height = 900
)

var (
	vertexShaderSource   string
	fragmentShaderSource string
	time                 float32

	model      mgl32.Mat4
	projection mgl32.Mat4
	camera     engine.Cam
)

func main() {

	window := core.InitGlfw(width, height, false, "Shading")
	defer glfw.Terminate()
	prog := initOpenGL()
	gl.UseProgram(prog)

	projection = mgl32.Perspective(mgl32.DegToRad(45.0), float32(width)/height, 0.1, 10.0)
	core.SetUniform(prog, "projection", projection)

	camera = engine.NewCam(0, 0, -3.41, 0, 0, 0)
	core.SetUniform(prog, "camera", camera.Mat())

	//Elementbuffer value currently unused -> what is it used for?
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
	core.SetUniform(prog, "mouse", mgl32.Vec2{engine.MouseX, engine.MouseY})
	core.SetUniform(prog, "layer", engine.NumKey) //layer
	core.SetUniform(prog, "camera", camera.Mat())

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.DrawElements(
		gl.TRIANGLES,                     // mode
		int32(len(engine.CubeIndices)*4), // count
		gl.UNSIGNED_INT,                  // type
		gl.PtrOffset(0),                  // element array buffer offset
	)

	engine.CalcMousePos(window)
	engine.ReadKeys(window)

	//ugly code! recompile shaders
	if engine.Pressed {
		time = 0
		core.UpdateShaders(prog)
		core.SetUniform(prog, "projection", projection)
	}
	engine.Pressed = false

	//pan camera with RMB, perspective/ortho with D key
	if engine.Rmb && !engine.Orth {
		camera = engine.NewCam(-engine.MouseX*4+1, -engine.MouseY*4+2, -3, 0, -0.1, 0)
	}
	if engine.Orth {
		camera = engine.NewCam(0, 0, -3.41, 0, 0, 0)
		engine.Orth = false
	}

	//reset time when switching layers
	if engine.TimeR {
		engine.TimeR = false
		time = 0
	}

	glfw.PollEvents()
	window.SwapBuffers()
}

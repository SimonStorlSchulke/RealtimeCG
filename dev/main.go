package main

import (
	"RealtimeCG/dev/core"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/rkusa/gm/math32"
)

const (
	width     = 600
	height    = 600
	toRadians = math32.Pi / 180
)

var (
	vertexShaderSource   string
	fragmentShaderSource string
	verts                = []float32{
		-1.0, -1.0, 0.0,
		0.0, -1.0, 1.0,
		1.0, -1.0, 0.0,
		0.0, 1.0, 0.0,
		-1.0, -1.0, 0.0,
		0.0, -1.0, 1.0,
		1.0, -1.0, 0.0,
		0.0, 1.0, 0.0,
	}
	indices = []uint32{
		0, 3, 1,
		1, 3, 2,
		2, 3, 0,
		0, 1, 2,
	}
	time float32

	//win95 screensaver values:
	direction            = true
	triOffset    float32 = 0
	triMaxOffset float32 = 1.0
	triIncrement float32 = 0.01
	model        mgl32.Mat4
	projection   mgl32.Mat4
	rotSpeed     = 0.01
	angle        = 0.0
)

func main() {

	window := core.InitGlfw(width, height, false, "Testwindow")
	defer glfw.Terminate()
	prog := initOpenGL()
	gl.UseProgram(prog)

	projection = mgl32.Perspective(mgl32.DegToRad(45.0), float32(width)/height, 0.1, 10.0)
	core.SetUniform(prog, "projection", projection)

	camera := mgl32.LookAtV(mgl32.Vec3{-3, -3, -3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	core.SetUniform(prog, "camera", camera)

	vao, elementBuffer := core.MakeVao(verts, indices)
	for !window.ShouldClose() {
		draw(vao, elementBuffer, window, prog)
	}

	//unused - Projection Stuff

}

//called every frame
func draw(vao, elementBuffer uint32, window *glfw.Window, prog uint32) {

	model = mgl32.Ident4()

	//Todo: Combine Model + ModelR
	core.SetUniform(prog, "model", model) //replace with &model later

	time += 0.01
	core.SetUniform(prog, "time", time)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(verts)/3))

	gl.DrawElements(
		gl.TRIANGLES,        // mode
		int32(len(indices)), // count
		gl.UNSIGNED_SHORT,   // type
		gl.Ptr(indices),     // element array buffer offset
	)

	glfw.PollEvents()
	window.SwapBuffers()
}

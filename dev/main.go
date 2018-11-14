package main

import (
	"RealtimeCG/dev/core"
	"fmt"

	"github.com/g3n/engine/math32"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	width     = 600
	height    = 600
	toRadians = math32.Pi / 180
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

	//win95 screensaver values:
	direction            = true
	triOffset    float32 = 0
	triMaxOffset float32 = 1.0
	triIncrement float32 = 0.01
	model        mgl32.Mat4
	rotSpeed     = 0.01
	angle        = 0.0
)

func main() {

	window := core.InitGlfw(width, height, false, "Testwindow")
	defer glfw.Terminate()
	prog := initOpenGL()

	//Transform
	//tMat := mgl32.Translate3D(0.4, 0.2, 0.3)

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

	//BOUNCE!
	if direction {
		triOffset += triIncrement
	} else {
		triOffset -= triIncrement
	}
	if math32.Abs(triOffset) >= triMaxOffset {
		direction = !direction
	}
	model = mgl32.Translate3D(triOffset, 0, 0)
	angle += rotSpeed
	modelR := mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0.4, 1, 0.2})

	//Todo: Combine Model + ModelR

	core.SetUniformFloat(prog, "tri_Offset", float32(triOffset))
	core.SetUniformMat4(prog, "model", model)   //replace with &model later
	core.SetUniformMat4(prog, "modelR", modelR) //replace with &model later

	time += 0.01
	core.SetUniformFloat(prog, "time", time)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

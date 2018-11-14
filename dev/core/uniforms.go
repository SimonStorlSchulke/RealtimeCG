package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

//must be called in gl main loop. Adds a uniform float32 to a program
func SetUniformFloat(prog uint32, name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(prog, gl.Str(name+"\x00")), value)
}

//Adds a uniform Matrix to a program
func SetUniformMat4(prog uint32, name string, mat mgl32.Mat4) {
	matUniform := gl.GetUniformLocation(prog, gl.Str(name+"\x00"))
	gl.UniformMatrix4fv(matUniform, 1, false, &mat[0])
}

//Set uniform of any type
func SetUniform(prog uint32, name string, value interface{}) {
	//Todo
	/*
		switch v := value.(type) {
		case int:
		case float32:
		case mgl32.Mat4:
		}
	*/
}

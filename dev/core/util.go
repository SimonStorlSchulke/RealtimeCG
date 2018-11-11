package core

import (
	"io/ioutil"

	"github.com/go-gl/gl/v4.1-core/gl"
)

//must be called in gl main loop. Adds a uniform float32 to a program
func SetSingleUniformFloat(program uint32, name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(program, gl.Str(name+"\x00")), value)
}

// makeVao initializes and returns a vertex array from the points provided.
func MakeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

//Reads a file and returns as string, usable as shader
func ReadShader(path string) (string, error) {
	byteArr, err := ioutil.ReadFile(path)
	str := string(byteArr[:])
	return str + "\x00", err
}

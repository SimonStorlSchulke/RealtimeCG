package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// makeVao initializes and returns a vertex array from the points provided.
func ElementBuffer(points []float32, indices []uint32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	//Todo add multiple vao to an vbo
	var vao uint32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	var elementbuffer uint32
	gl.GenBuffers(1, &elementbuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementbuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(&indices[0]), gl.STATIC_DRAW)

	return elementbuffer
}

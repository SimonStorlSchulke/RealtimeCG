package core

import (
	"RealtimeCG/dev/obj"
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// makeVao initializes and returns a vertex array from the points provided.
func ElementBuffer(model obj.Model) uint32 {

	var points []float32
	var indices []uint32

	//ugly
	for i, _ := range model.VecIndices {
		indices = append(indices, uint32(model.VecIndices[i]))
	}

	for i, v := range model.Vecs {
		points = append(points, v.X(), v.Y(), v.Z(), 1, 0, 0, model.Uvs[i][0], model.Uvs[i+1][1]) //todo color & uv from model
	}

	fmt.Println(indices)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	//Todo add multiple vao to an vbo
	var vao uint32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	//vert position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, nil)
	gl.EnableVertexAttribArray(0)

	//vert color
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	//vert UV
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.EnableVertexAttribArray(2)

	var elementbuffer uint32
	gl.GenBuffers(1, &elementbuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementbuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(&indices[0]), gl.STATIC_DRAW)

	return elementbuffer
}

package engine

var CubeVerts []float32
var CubeIndices []uint32
var CubeColors []float32

type Mesh struct {
	Verts   []float32
	Indices []uint32
} //unused

func init() {

	CubeVerts = []float32{
		//3x verts 3x vert color 2x uv
		// front
		-1, -1, 1, 1, 0, 0, 1, 1,
		1, -1, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 0, 1, 0, 0,
		-1, 1, 1, 1, 0, 0, 0, 1,
		// back
		-1, -1, -1, 1, 0, 0, 1, 1,
		1, -1, -1, 1, 0, 0, 1, 0,
		1, 1, -1, 1, 0, 0, 0, 0,
		-1, 1, -1, 1, 0, 0, 0, 1,
	}

	CubeIndices = []uint32{
		//front
		2, 1, 0,
		2, 3, 0,
		// right
		6, 5, 1,
		1, 2, 6,
		// back
		7, 6, 5,
		5, 4, 7,
		// left
		3, 0, 4,
		3, 7, 4,
		// bottom
		1, 5, 4,
		1, 0, 4,
	}
}

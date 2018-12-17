package engine

var CubeColors []float32
var CubeVerts []float32
var CubeIndices []uint32

type Mesh struct {
	Verts   []float32
	Indices []uint32
} //unused

func init() {

	CubeVerts = []float32{
		// front
		-1.0, -1.0, 1.0,
		1.0, -1.0, 1.0,
		1.0, 1.0, 1.0,
		-1.0, 1.0, 1.0,
		// back
		-1.0, -1.0, -1.0,
		1.0, -1.0, -1.0,
		1.0, 1.0, -1.0,
		-1.0, 1.0, -1.0,
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
		// top
		3, 2, 6,
		3, 7, 6,
	}

	CubeColors = []float32{
		// front colors
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,
		1.0, 1.0, 1.0,
		// back colors
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,
		1.0, 1.0, 1.0,
	}
}

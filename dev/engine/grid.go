package engine

func generateGrid(rows, columns int, size float32) Mesh {

	verts := []float32{}
	indices := []uint32{}

	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			verts = append(verts, float32(x)*size, 0, float32(y)*size)
		}
	}
	return Mesh{verts, indices}

}

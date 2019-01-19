package engine

import (
	"fmt"
	"io/ioutil"
)

//Load .obj from filepath and return as Mesh
func LoadObj(path string) (Mesh, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return Mesh{}, err
	}
	fStr := string(f[:])
	fmt.Println(fStr)
	return Mesh{}, err
}

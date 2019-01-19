package core

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

//SetUniform of any type
func SetUniform(prog uint32, name string, value interface{}) error {

	switch v := value.(type) {
	case float32:
		gl.Uniform1f(gl.GetUniformLocation(prog, gl.Str(name+"\x00")), v)
		return nil
	case int:
		gl.Uniform1i(gl.GetUniformLocation(prog, gl.Str(name+"\x00")), int32(v))
		return nil
	case mgl32.Mat4:
		matUniform := gl.GetUniformLocation(prog, gl.Str(name+"\x00"))
		gl.UniformMatrix4fv(matUniform, 1, false, &v[0])
		return nil
	default:
		return fmt.Errorf("type %v not supported", v)
	}
}

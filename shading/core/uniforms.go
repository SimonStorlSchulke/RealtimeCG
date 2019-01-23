package core

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

//SetUniform of any type
func SetUniform(prog uint32, name string, value interface{}) error {

	loc := gl.GetUniformLocation(prog, gl.Str(name+"\x00"))

	switch v := value.(type) {
	case float32:
		gl.Uniform1f(loc, v)
		return nil
	case int:
		gl.Uniform1i(loc, int32(v))
		return nil
	case mgl32.Mat4:
		gl.UniformMatrix4fv(loc, 1, false, &v[0])
		return nil
	case mgl32.Vec2:
		gl.Uniform2f(loc, v.X(), v.Y())
		return nil
	case mgl32.Vec3:
		gl.Uniform3f(loc, v.X(), v.Y(), v.Z())
		return nil
	default:
		return fmt.Errorf("type %v not supported", v)
	}
}

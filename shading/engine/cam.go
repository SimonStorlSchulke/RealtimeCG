package engine

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

type Cam struct {
	origin mgl32.Vec3
	target mgl32.Vec3
}

func (c *Cam) Origin() mgl32.Vec3 {
	return c.origin
}

func NewCam(originX, originY, originZ, targetX, targetY, targetZ float32) Cam {
	return Cam{mgl32.Vec3{originX, originY, originZ}, mgl32.Vec3{targetX, targetY, targetZ}}
}

func (c *Cam) Mat() mgl32.Mat4 {
	return mgl32.LookAtV(c.origin, c.target, mgl32.Vec3{0, 1, 0})
}

func (c Cam) Info() string {
	return fmt.Sprintf("Camera Info:\nOrigin: x%v y%v z%v\nTarget: x%v y%v z%v", c.origin[0], c.origin[1], c.origin[2], c.target[0], c.target[1], c.target[2])
}

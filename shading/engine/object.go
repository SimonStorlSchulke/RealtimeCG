package engine

import "github.com/go-gl/mathgl/mgl32"

type object interface {
	Origin() mgl32.Vec3
}

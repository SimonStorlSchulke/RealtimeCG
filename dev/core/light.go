package core

import (
	"github.com/go-gl/mathgl/mgl32"
)

//Light Represents a 3D light source
type Light struct {
	Color     mgl32.Vec3
	Direction mgl32.Vec3
	Intensity float32
	LightType int
}

//LightTypes
const (
	LIGHTPOINT = iota + 1
	LIGHTDIRECTIONAL
	LIGHTSPOT
	LIGHTAMBIENT
)

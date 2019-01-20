package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

//InitGlfw initializes glfw and returns a Window to use.
func InitGlfw(width, height int, resizable bool, title string) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	if resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

//Returns gl program by given vertex- and fragmentshader
func NewProgram(vShaderSource, fShaderSource string) (uint32, error) {
	vShader, err := CompileShader(vShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	fShader, err := CompileShader(fShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}
	prog := gl.CreateProgram()
	gl.AttachShader(prog, vShader)
	gl.AttachShader(prog, fShader)
	gl.LinkProgram(prog)

	var status int32
	gl.GetProgramiv(prog, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		return 0, err
	}

	return prog, nil
}

package engine

import (
	"RealtimeCG/dev/core"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {

	var err error
	vertexShaderSource, err = ReadShader("shader/test.vert")
	if err != nil {
		panic(err)
	}
	fragmentShaderSource, err = ReadShader("shader/test.frag")
	if err != nil {
		panic(err)
	}
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw(title string, width, height int, resizable bool) *glfw.Window {
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

//Startup Opengl and attach shaders
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
	prog, err := core.NewProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		panic(err)
	}
	gl.Enable(gl.DEPTH_TEST)

	return prog
}

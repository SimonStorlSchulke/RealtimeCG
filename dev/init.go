package main

import (
	"RealtimeCG/dev/core"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func init() {
	runtime.LockOSThread()

	var err error
	vertexShaderSource, err = core.ReadShader("shader/test.vert")
	if err != nil {
		panic(err)
	}
	fragmentShaderSource, err = core.ReadShader("shader/test.frag")
	if err != nil {
		panic(err)
	}
}

//Startup Opengl and attach shaders
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
	vertexShader, err := core.CompileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := core.CompileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

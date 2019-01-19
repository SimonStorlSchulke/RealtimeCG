package core

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

//ReadShader reads a file and returns as string, usable as shader
func ReadShader(path string) (string, error) {
	byteArr, err := ioutil.ReadFile(path)
	str := string(byteArr[:])
	return str + "\x00", err
}

//CompileShader returns a compiled Shader from a string and an error
func CompileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	//errorcheck
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}
	return shader, nil
}

//CreateProgram Returns gl program from given vertex- and fragmentshader
func CreateProgram(vShaderSource, fShaderSource string) (uint32, error) {
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

func UpdateShaders(progOld uint32) uint32 {
	vertexShaderSource, err := ReadShader("shader/test.vert")
	if err != nil {
		panic(err)
	}
	fragmentShaderSource, err := ReadShader("shader/test.frag")
	if err != nil {
		panic(err)
	}

	progNew, err := CreateProgram(vertexShaderSource, fragmentShaderSource)
	if err == nil {
		gl.UseProgram(progNew)
		fmt.Println("recompiled shaders")
		return progNew
	} else {
		fmt.Println(err)
		return progOld
	}
}

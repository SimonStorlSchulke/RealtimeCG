package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	//Vertex Shader
	myVShader = `
		#version 330 layout (location=0) in vec3 pos 
		void main(){
			gl_Position = vec4(0.25*pos.x, 0.25*pos.y, pos.z, 1.0);
		}
	`

	//Fragment Shader
	myFShader = `
		#version 330 out vec4 color
		void main(){
			color = vec4(1.0, 1.0, 0.0, 1.0);
		}
	`
)

var (
	triVAO    uint32
	triVBO    uint32
	triShader uint32
)

//Create VAO
func drawTriangle() {
	//Verts Position bestimmen
	vertices := []float32{
		-1, -1, 0,
		1, -1, 0,
		0, 1, 0,
	}
	fmt.Println(vertices) //temp
	gl.GenVertexArrays(1, &triVAO)
	gl.BindVertexArray(triVAO)

	gl.GenBuffers(1, &triVBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, triVBO)

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW) //Für unsafe *Ptr - gl.Ptr verwenden

	//Wirklich GL_Float?
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, gl.Ptr(0))

	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)
}

func main() {

	//in Go nötig um GL Prozesse auf einen Thread zu locken.
	runtime.LockOSThread()

	//01 Erste Amtshandlung, Definition der Dimensionen des ersten Fensters
	const WIDTH int = 800
	const HEIGHT int = 600

	//02 Initialisieren von GLFW, damit auch ein Fenster ereugt werden kann. if ist vorhanden, um auf Fehler zu prüfen.
	if glfw.Init() != nil {
		//Wurde nichts initialisiert, dann bitte beenden
		panic("GLFW Initialisierung fehlgeschlagen!")
	}

	//defer wird in Go immer zuletzt vor dem Beenden des Programms ausgeführt
	defer glfw.Terminate()

	//03 Einstellungen des GLFW Fensters festlegen
	//OpenGL Version des zu anzeigenden Fenster ist Version 4. Die Funktion setzt Eigenschaften über die Auswalh von Enums.
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	//Als nächstes muss das Profil gesetzt werden. Wie wird der auszuführende Code behandelt? Hauptprofil =
	//No backwards compatability. Hier tritt ein Fehler auf bei der Findung von altem Code.
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	//Zulassen von neuen Versionen von OpenGL.
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// GL Fenster seitens der Dimensionen und Einstellungen bereit.

	// 04 Jetzt muss das GL Fenster erzeugt werden.
	window, err := glfw.CreateWindow(WIDTH, HEIGHT, "Test_Window", nil, nil)

	// Ist das Fenster korrekt aufgebaut worden?
	if err != nil {
		panic("GLFW Fenster Aufbau fehlgeschlagen!")
	}

	//04.01, Tricky, Wir brauchen die Buffergrösse bzw. Informationen zur Buffergrösse.
	//Die Dimensionen sind gesetzt, aber der symbolische Raum darin muss definiert sein.
	bufferWidth, bufferHeight := window.GetFramebufferSize()

	//04.02, Alles was geschrieben werden soll, muss in das zuvor angelegte Fenster. GLEW Bescheid geben,
	//dass das Fenster mit dem OpenGL Kontext verknüpft ist. Können auch mehrere Fenster angelegt werden.
	window.MakeContextCurrent()

	//04.04, GL initialisieren. Initialisierung prüfen und bei Fehler das Fenster schließen, da es nun vorhanden ist.
	err = gl.Init()
	if err != nil {
		window.Destroy()
		panic("GL Init Fehlgeschlagen")
	}

	//05 Einstellen der GL Viewport Größe. Quasi die Größe eines Bereichs im Fenster in den wir schreiben. Nullpunkt und die Dimensionen.
	gl.Viewport(0, 0, int32(bufferWidth), int32(bufferHeight)) //nomal gucken
	//06 Ticking/Loop bis das Fenster geschlossen wird. Stop Loop, Programm geschlossen.
	for !window.ShouldClose() {
		// Get + Handle User Input Events. Z.B. um das Fenster zu schließen oder andere Events zu starten.
		glfw.PollEvents()

		//Säubern des Fensters, ganz gleich was im Frame Buffer hinterlegt ist. Die Funktion macht alles symbolisch platt. Sauberer Frame.
		gl.ClearColor(0, 0.8, 0.3, 1)

		//Wir müssen klarstellen welche Buffer wir löschen wollen. Gibt ja mehr als nur die Farbe. Stencil, Z-Depth.
		gl.Clear(gl.COLOR_BUFFER_BIT)

		//2 Szenen vorhanden und es wird in die Szene geschrieben die noch nicht sichtbar ist.
		window.SwapBuffers()
	}
}

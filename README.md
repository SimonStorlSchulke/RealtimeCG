# RealtimeCG
Veranstaltung Echtzeit Computergrafik an der HS Furtwangen

1. Init OpenGL + GLFW
```go
package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

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
		gl.ClearColor(1, 0, 0, 1)

		//Wir müssen klarstellen welche Buffer wir löschen wollen. Gibt ja mehr als nur die Farbe. Stencil, Z-Depth.
		gl.Clear(gl.COLOR_BUFFER_BIT)

		//2 Szenen vorhanden und es wird in die Szene geschrieben die noch nicht sichtbar ist.
		window.SwapBuffers()
	}
}
```

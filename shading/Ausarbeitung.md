# Ausarbeitung Shading


Die für Shader in OpenGL verwendete Sprache ist GLSL (OpenGL Shading Language)
GLSL Shader werden grundsätzlich unterteilt in Vertex Shader, die die Darstellung und Position der tatsächlichen Geometrie im Raum steuern, sowie Fragment shader, die für Farbgebung, Lichteffekte etc. zuständig sind. Diese Ausarbeitung beschäft sich mit im Fragment Shader definierten Prozedurale Texturen.

## 1. Vektoren
Mit die wichtigsten Datenstrukturen in GLSL sind VeKtoren. Diese stellen einen Verbund aus je 2, 3 oder 4 Fließkommazahlen dar. Für Farben wird zum Beispiel meist entweder vec3 für Rot Grün und Blau, oder vec4 für Rot, Grün, Blau und Alpha verwendet.
Das Umwandeln von Vektortypen untereinander ist meist Problemlos, so kann z.b. mit
`vec4(MeinVec3, 1.0)` MeinVec3 zu Vec4 umgewandelt werden, wobei als vierter Wert 1 übergeben wird. Mit `vec3(MeinFloat)` wird MeinFloat an alle drei Kanäle des Vectors übergeben.


## 1. Texturkoordinaten
Um die Generierung von Prozeduralen Texturen zu ermöglichen sind Daten nötig, die vom Hauptprogramm an den Shader und somit die Grafikkarte übergeben werden. Abgesehen von der reinen Geometrie, die im Vertex Shader verarbeitet wird, ist Mindestvoraussetzung für Texturen auf Geometrie die Übergabe von Texturkoordinaten. Dabei wird ein vec(2) Wert übergeben - eine Datenstruktur aus zwei float Werten (x und y). Dieser stellt pro Pixel die Texturkoordinaten auf der Geometrie dar. Auf einem Würfel mit X als Rot und Y als Grün angezeigt, sehen die Texturkoordinaten zum Beispiel so aus:

![SCREEN 1 texcoords]()

## 2. Farbverläufe
Die wohl einfachste Form der Textur sind einfache Farbverläufe, die entweder die X oder Y Komponente der Texturkoordinaten als Verlauf auf derm Objekt darstellen.

```glsl
c = TexCoord.y;
```
> `c` steht hier immer für die entgültig als `vec4` (RGBA) augegebene Farbe - auch wenn das Beispiel `float` oder `vec3` ergibt.

![SCREEN 1 verlauf1]()

<br>
mithilfe des Modulo Operators, können auch sich wiederhohlende Farbverläufe dargestellt werden:

```glsl
c = skalierung * mod(TexCoord.y, 1/skalierung);
```
Hirbei wird das Ergebniss zusätzlich noch mit der Skalierung multipliziert, um die Reichweite wieder auf 0 - 1 zu mappen.

<br>

Um den Effekt auf und absteigender Verläufe zu erzielen, kann zusätzlich noch eine Kondition eingefügt werden, die den Wert >0.5 umkehrt und im Anschluss verdoppelt.
```glsl
c = scale * mod(val, 1/scale);
if (c > 0.5) {
    c = 1-f;
}
c *= 2;
```
<br>
Weichere Verläufe sind mit Sinus oder Cosinus moglich:

```glsl
c = (sin(TexCoord.y * skalierung * PI) + 1) / 2;
```
Zusätzlich muss der Wert um 1 erhöht und halbiert werden, um die Reichweite wieder auf 0 - 1 zu mappen.

<br>

Durch mischen dieser einfachen Funkionen in den verschiedenen Farbkanälen, lassen sich schon komplexere Muster erzeugen, wie zum Beispiel eine Schuppentextur.
![Sc]()

<br>

## 3. Animierte Texturen:
Über sogenannte Uniforms, können zusätzliche Werte aus dem Hauptprogramm an den Shader gegeben werden. Um animierte Texturen zu ermöglichen, kann zum Beispiel die Zeit übergeben werden.
Hier wird die Zeit zur Y-Koordinate hinzugefügt, um einen bewegten Farbverlauf zu erzeugen:

```glsl
c = (sin(TexCoord.y + time * skalierung * PI) + 1) / 2;
```
<br>

Interessante Effekte können auch mit der Übergabe der mittles GLFW ausgelesenen Mausposition erzielt werden.
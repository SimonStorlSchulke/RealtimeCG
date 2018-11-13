#version 410
in vec3 vp;

uniform float time;

void main() {
    gl_Position = vec4(vp, 2 * abs(sin(time)));
}
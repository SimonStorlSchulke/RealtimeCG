#version 410
in vec3 vp;

uniform float time;
uniform mat4 model;
uniform mat4 modelR;

void main() {

    float v = 2 - abs(sin(time));

    //gl_Position = vec4(vp, 2 * abs(sin(time)));
    gl_Position =  modelR * model * vec4(vp, v);
}
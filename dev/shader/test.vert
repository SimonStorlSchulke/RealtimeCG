#version 410
in vec3 vp;

uniform float time;
uniform mat4 model;
uniform mat4 projection;
uniform mat4 camera;

void main() {

    gl_Position =  projection * camera * model * vec4(vp, 1);
}
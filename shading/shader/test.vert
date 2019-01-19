#version 410

//load vert positions, -color and -uv
layout (location = 0) in vec3 aPos;
layout (location = 1) in vec3 aColor;
layout (location = 2) in vec2 aTexCoord;

out vec3 vertCol;
out vec2 TexCoord;

uniform float time;
uniform mat4 model;
uniform mat4 projection;
uniform mat4 camera;

void main() {
    vertCol = aColor;
    TexCoord  =aTexCoord;
    gl_Position =  projection * camera * model * vec4(aPos, 1);
}
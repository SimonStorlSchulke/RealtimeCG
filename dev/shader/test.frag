#version 410

out vec4 frag_colour;
uniform float time;

void main() {
    float v = abs(sin(time));
    frag_colour = vec4(1-v, 1-v, 0.2, 1.0);
}
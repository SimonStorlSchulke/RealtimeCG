#version 410

uniform float time;
float near = 0.1; 
float far  = 100.0;

out vec4 frag_colour;
in vec3 vertCol;
in vec2 TexCoord;

float LinearizeDepth(float depth) 
{
    float z = depth * 2.0 - 1.0;
    return (2.0 * near * far) / (far + near - z * (far - near));	
}

void main() {
    float v = abs(sin(time));
    float depth = LinearizeDepth(gl_FragCoord.z) / far * 2;

    //Vertex Color
    //frag_colour = vec4(vertCol, 1.0);
    frag_colour = vec4(TexCoord, v, 1.0);
}
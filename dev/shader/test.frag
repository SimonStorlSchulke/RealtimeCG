#version 410

out vec4 frag_colour;
uniform float time;
float near = 0.1; 
float far  = 100.0; 

float LinearizeDepth(float depth) 
{
    float z = depth * 2.0 - 1.0;
    return (2.0 * near * far) / (far + near - z * (far - near));	
}

void main() {
    float v = abs(sin(time));
    float depth = LinearizeDepth(gl_FragCoord.z) / far * 2;
    frag_colour = vec4(vec3(depth), 1.0);
}
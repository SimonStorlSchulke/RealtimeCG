#version 410

//uniforms
uniform float time;
uniform int layer;
uniform vec2 mouse;
uniform mat4 camera;

//depth vars
float near = 0.1; 
float far  = 100.0;

out vec4 frag_colour;
in vec3 vertCol;
in vec2 TexCoord;

const float PI = 3.1416;

float x = TexCoord.x;
float y = TexCoord.y;
float v = abs(sin(time));
vec3 c;

//blendmodes
vec3 screen(vec3 base, vec3 top, float fac) {
    return mix(base, 1 - (1-base) * (1-top), fac);
}

vec3 mul(vec3 base, vec3 top, float fac) {
    return mix(base, base * top, fac);
}

// 2D Random
float random (in vec2 st) {
    return fract(sin(dot(st.xy,
        vec2(12.9898,78.233)))
            * 43758.5453123);
}

// GRADIENTS
float LinearizeDepth(float depth) {
    float z = depth * 2.0 - 1.0;
    return (2.0 * near * far) / (far + near - z * (far - near));	
}

float HardGradients(float val, float num) {
    return num * mod(val, 1/num);
}

float LinearGradients(float val, float num) {
    float f = num * mod(val, 1/num);
    if (f > 0.5) {
        f = 1-f;
    }
    return f;
}

float SinGradients(float val, float num) {
    return (sin(val * num * 2 * PI) + 1) / 2;
}

float depth() {
    return LinearizeDepth(gl_FragCoord.z) / far * 2;
}

vec3 scales() {
    vec3 baseCol = vec3(0.055, 0.063, 0.024);
    vec3 scaleCol = vec3(.4, .75, .15);
    vec3 dotsCol = vec3(0.16, 0.16, 0.025);

    float distortScales = mod(x+1.5*y, 0.4);
    float scales = SinGradients(distortScales, 5) * SinGradients(y, 12);

    float distortDots = mod(x+1*y, 0.4);
    float dots = SinGradients(distortDots, 21) * SinGradients(y, 21);

    float gradient = (1-HardGradients(x*y, 1));

    scales = pow(scales, .5);

    vec3 c = mix(baseCol, scaleCol, scales);
    c = mix(c, dotsCol, dots);
    c = 1.3*mul(c,  vec3(1,.4,.2), gradient * .8);
    return c;
}

// Aus Book Of Shaders -
// 2D Noise based on Morgan McGuire @morgan3d
// https://www.shadertoy.com/view/4dS3Wd
float noise(in vec2 st) {
    vec2 i = floor(st);
    vec2 f = fract(st);

    // Four corners in 2D of a tile
    float a = random(i);
    float b = random(i + vec2(1.0, 0.0));
    float c = random(i + vec2(0.0, 1.0));
    float d = random(i + vec2(1.0, 1.0));

    // Cubic Hermine Curve.  Same as SmoothStep()
    vec2 u = f*f*(3.0-2.0*f);
    // u = smoothstep(0.,1.,f);

    // Mix 4 coorners percentages
    return mix(a, b, u.x) +
            (c - a)* u.y * (1.0 - u.x) +
            (d - b) * u.x * u.y;
}

float fbm(in vec2 st, int OCTAVES, vec2 offset) {
    // Initial values
    int maxOctaves = 16;
    OCTAVES = clamp(OCTAVES, 1, maxOctaves);
    float value = 0.0;
    float amplitude = .5;
    //
    // Loop of octaves
    for (int i = 0; i < OCTAVES; i++) {
        value += amplitude * noise(st+offset*i);
        st *= 2.;
        amplitude *= .5;
    }
    return value;
}

float waves(float scale, float distortion, int octaves, int banding) {
    float noise = fbm(TexCoord*scale, octaves, vec2(0));
    float grad = SinGradients(mix(x, noise, distortion), banding);
    return grad;
}

vec2 distortCoords(in vec2 st, in float strength, in float map) {
    map -= 0.5;
    vec2 ou = st + map*strength;
    return ou;
}

vec2 transformCoords(in vec2 st, float x, float y) {
    return vec2(st.x + x, st.y + y);
}

float fire() {
    //move up coords
    vec2 cTransp = transformCoords(TexCoord, 0, -time);
    
    //distort coords
    float distNoise = fbm(TexCoord*12, 5, vec2(0));
    cTransp = distortCoords(cTransp, .1, distNoise);

    //generate noise
    float noise = fbm(cTransp*4, 5, vec2(0));
    noise = 1.3*pow(noise, y*6);
    return noise;
}

vec3 sparks() {
    float distNoise = fbm(TexCoord*4, 2, vec2(0));
    float noise = fbm(TexCoord*vec2(50, 20) + vec2(distNoise*20, -time * 12), 2, vec2(0));
    noise = pow(noise, 5);
    if (noise > 0.17) {
        vec3 c1 = vec3(1.000000, 0.980136, 0.749383);
        vec3 c2 = vec3(.800000, 0.146404, 0.014866);
        vec3 col = mix(c1, c2, y);
        return col * (1-pow(y, 2));
    }
    return vec3(0);
}

vec3 cartoonFire() {
    float exponent = 3-(1-2*mouse.y)*2.7;
    float noise = pow(fire(), exponent);
    vec3 c;

    //gradient
    if(noise > 1) {
        c = vec3(1.000000, 0.980136, 0.749383);
    } else if(noise > 0.7) {
        c =  vec3(1.000000, 0.980136, 0.2);
    } else if(noise > 0.56) {
        c =  vec3(1.000000, 0.451227, 0.025711);
    } else if(noise > 0.35) {
        c =  vec3(1.000000, 0.286404, 0.024866);
    } else if(noise > 0.24) {
        c =  vec3(.800000, 0.146404, 0.014866);
    } else if(noise > 0.14) {
        c =  vec3(0.216065, 0.030013, 0.007945);
    } else if(noise > 0.05) {
        c =  vec3(0.05);
    }
    c += sparks();
    
    return c;
}


vec3 sci() {
    float n0 = pow(fbm(TexCoord*50+ 10 - vec2(time*0.3, 0), 3, vec2(0)), 2);
    float n1 = pow(fbm(TexCoord+camera[0].y*5 + time*0.1, 3, vec2(0)), 3);
    
    float res = (n0 + n1);
    if(res > 0.65) {
        return vec3(res) * vec3(0,1.5,.4);
    }
    return vec3(.06,n1*.3 + .1, n0*.3 + .13);
}


void main() {
    switch(layer) {
        case 1: //Texturkoordinaten
            c = vec3(TexCoord, 0);
            break;
        case 2: //Separieren
            c = vec3(TexCoord.y);
            break;
        case 3: //Gradients
            c = vec3(HardGradients(y, 4));            
            break;
        case 4: //Sin Gradients
            c = vec3(SinGradients(x, 5));
            break;
        case 5: //scales
            c = scales();
            break;
        case 6: //Animation
            c = vec3(SinGradients(x + time, 5) * SinGradients(y + time, 5));            
            break;
        case 7: //fract
            c = vec3(fract(TexCoord*2), 0);
            break;
        case 8: //noise
            c = vec3(noise(TexCoord*20));
            break;
        case 9: //fbm
            c = vec3(fbm(TexCoord*12, int(time)+1, vec2(0)));
            break;
        case 0: // distortion,
            c = vec3(fbm(distortCoords(TexCoord*10, .7, time*fbm(TexCoord*12, 5, vec2(0))), 5, vec2(0)));            
            break;
        case 11: //waves, camera
            c = vec3(waves(5, camera[0].y * 2, 5, 5));
            break;
        case 12: //fire
            c = vec3(fire());
            break;
        case 13: //cartoon fire
            c = sparks();            
            break;
        case 14: //cartoon fire
            c = cartoonFire();            
            break;
        case 15: //depth
            c = vec3(depth());
            break;
        case 16: //mouse
            c = vec3(mouse, 0);
            break;
        case 17: //sci
            c = sci();
            break;
        default:
            c = vec3(1,0,0);
    }
    frag_colour = vec4(c,1);
}
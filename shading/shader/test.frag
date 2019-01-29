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

// GRADIENTS
float LinearizeDepth(float depth) {
    float z = depth * 2.0 - 1.0;
    return (2.0 * near * far) / (far + near - z * (far - near));	
}

float HardGradients(float val, float scale) {
    return scale * mod(val, 1/scale);
}

float LinearGradients(float val, float scale) {
    float f = scale * mod(val, 1/scale);
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

// 2D Random
float random (in vec2 st) {
    return fract(sin(dot(st.xy,
        vec2(12.9898,78.233)))
            * 43758.5453123);
}

// Aus Book Of Shaders -
// 2D Noise based on Morgan McGuire @morgan3d
// https://www.shadertoy.com/view/4dS3Wd
float noise(in vec2 st) {
    vec2 i = floor(st);

    return i.x / 12;
    return i.y / 12;

    // Four corners in 2D of a tile
    float a = random(i);
    float b = random(i + vec2(1.0, 0.0));
    float c = random(i + vec2(0.0, 1.0));
    float d = random(i + vec2(1.0, 1.0));

    return a;
    
    vec2 f = fract(st);
    return f.x;

    // Cubic Hermine Curve
    vec2 u = f*f*(3.0-2.0*f);

    //quintic interpolation curve
    u = f*f*f*(f*(f*6.-15.)+10.);

    return u.x;

    // Mix 4 coorners percentages
    return mix(a, b, u.x) +
            (c - a)* u.y * (1.0 - u.x) +
            (d - b) * u.x * u.y;
}

//Improved SIMPLEX NOISE
vec3 mod289(vec3 x) { return x - floor(x * (1.0 / 289.0)) * 289.0; }vec2 mod289(vec2 x) { return x - floor(x * (1.0 / 289.0)) * 289.0; }vec3 permute(vec3 x) { return mod289(((x*34.0)+1.0)*x); }
float simplexNoise(vec2 v){v *= 0.5;const vec4 C = vec4(0.211324865405187,0.366025403784439,-0.577350269189626,0.024390243902439); vec2 i  = floor(v + dot(v, C.yy));vec2 x0 = v - i + dot(i, C.xx);vec2 i1 = vec2(0.0);i1 = (x0.x > x0.y)? vec2(1.0, 0.0):vec2(0.0, 1.0);vec2 x1 = x0.xy + C.xx - i1;vec2 x2 = x0.xy + C.zz;i = mod289(i);vec3 p = permute(permute( i.y + vec3(0.0, i1.y, 1.0)) + i.x + vec3(0.0, i1.x, 1.0 ));vec3 m = max(0.5 - vec3(dot(x0,x0),dot(x1,x1),dot(x2,x2)), 0.0);m = m*m ; m = m*m ;vec3 x = 2.0 * fract(p * C.www) - 1.0;vec3 h = abs(x) - 0.5;vec3 ox = floor(x + 0.5);vec3 a0 = x - ox;m *= 1.79284291400159 - 0.85373472095314 * (a0*a0+h*h);vec3 g = vec3(0.0);g.x  = a0.x  * x0.x  + h.x  * x0.y;g.yz = a0.yz * vec2(x1.x,x2.x) + h.yz * vec2(x1.y,x2.y);return 130.0 * dot(m, g);}

float fbm(in vec2 st, int OCTAVES) {
    // Initial values
    int maxOctaves = 16;
    OCTAVES = clamp(OCTAVES, 1, maxOctaves);
    float value = 0.0;
    float amplitude = .5;
    //
    // Loop of octaves
    for (int i = 0; i < OCTAVES; i++) {
        value += amplitude * noise(st);                   //default
        //value += amplitude * (simplexNoise(st)*0.5+0.5);    //simplex
        st *= 2.;
        amplitude *= .5;
    }
    return value;
}

float perlin(in vec2 st, int OCTAVES) {int maxOctaves = 16;OCTAVES = clamp(OCTAVES, 1, maxOctaves);float value = 0.0;float amplitude = .5;for (int i = 0; i < OCTAVES; i++) {value += amplitude * (simplexNoise(st)*0.5+0.5);st *= 2.;amplitude *= .5;}return value;}


float waves(float scale, float distortion, int octaves, int banding) {
    float noise = perlin(TexCoord*scale, octaves);
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

float fire(vec2 st) {
    //move up coords
    vec2 cTransp = transformCoords(st, 0, -time);
    
    //distort coords
    float distNoise = perlin(st*12, 5);
    cTransp = distortCoords(cTransp, .1, distNoise);

    //generate noise
    float noise = perlin(cTransp*4, 5);
    noise = 1.3*pow(noise, y*6);
    return noise;
}

vec3 sparks() {
    float distNoise = perlin(TexCoord*4, 2);
    float noise = perlin(TexCoord*vec2(50, 20) + vec2(distNoise*20, -time * 12), 2);
    noise = pow(noise, 5);
    if (noise > 0.15) {
        vec3 c1 = vec3(1.000000, 0.980136, 0.749383);
        vec3 c2 = vec3(.800000, 0.146404, 0.014866);
        vec3 col = mix(c1, c2, y);
        return col * (1-pow(y, 2));
    }
    return vec3(0);
}

vec3 cartoonFire(vec2 st) {
    float exponent = 3-(1-2*mouse.y)*2.7;
    float noise = pow(fire(st), exponent);
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
    float n0 = pow(perlin(TexCoord*50+ 10 - vec2(time*0.3, 0), 3), 2);
    float n1 = pow(perlin(TexCoord+camera[0].y*5 + time*0.1, 3), 3);
    
    float res = (n0 + n1);
    if(res > 0.65) {
        return vec3(res) * vec3(0,1.5,.4);
    }
    return vec3(.06,n1*.3 + .1, n0*.3 + .13);
}

//http://nuclear.mutantstargoat.com/articles/sdr_fract/
vec2 julia(vec2 seed, int iterations) {
    
    //Mapping
    vec2 z;
    z.x = 3.0 * (TexCoord.x - 0.5);
    z.y = 2.0 * (TexCoord.y - 0.5);

    //iterate
    int i;
    for(i=0; i<iterations; i++) {
        float x = (z.x * z.x - z.y * z.y) + seed.x;
        float y = (z.y * z.x + z.x * z.y) + seed.y;

        if((x * x + y * y) > 4.0) break;
        z.x = x;
        z.y = y;
    }
    return z;
}

//http://nuclear.mutantstargoat.com/articles/sdr_fract/
vec3 julia1(vec2 seed, int iterations) {
    
    //Mapping
    vec3 z;
    z.x = 3.0 * (TexCoord.x - 0.5);
    z.y = 2.0 * (TexCoord.y - 0.5);

    //iterate
    int i;
    vec3 col = vec3(0);
    for(i=0; i<iterations; i += 1) {
        float x = (z.x * z.x - z.y * z.y) + seed.x;
        float y = (z.y * z.x + z.x * z.y) + seed.y;

        if((x * x + y * y) > 4.0) break;
        vec3 z1 = vec3(x,y,i);
        z.x = x;
        z.y = y;
        col = vec3((1.0/iterations) * i);
    }
   // col *= z * z* 1;
    return col;
}

//http://nuclear.mutantstargoat.com/articles/sdr_fract/
vec3 julia2(vec2 seed, int iterations) {
    
    //Mapping
    vec2 z;
    z.x = 3.0 * (TexCoord.x - 0.5);
    z.y = 2.0 * (TexCoord.y - 0.5);

    //iterate
    int i;
    for(i=0; i<iterations; i++) {
        float x = (z.x * z.x - z.y * z.y) + seed.x;
        float y = (z.y * z.x + z.x * z.y) + seed.y;

        if((x * x + y * i) > 4.0) break;
        z.x = x;
        z.y = y;
    }
    
    float r,g,b;
    r = z.y;
    g = z.x * z.y;
    b = abs(z.y) * 0.7;
    float stars = pow(perlin(z*20+time, 1), 8) * 150;
    if (stars < 0.4) stars = pow(stars, 3);
    
    return (vec3(r,g,b) *  perlin(z+time*0.5, 5)) * (SinGradients(z.x, 0.6) + 0.2) + stars;
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
            c = vec3(SinGradients(x, 3));
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
            c = vec3(noise(TexCoord*12));
            //c = vec3(simplexNoise(TexCoord*8)*0.5 + 0.5);
            break;
        case 9: //fbm
            c = vec3(fbm(TexCoord*8 + 12, int(time)+1));
            break;
        case 0: // distortion,
            c = vec3(perlin(distortCoords(TexCoord*8 + 12, .7, time*perlin(TexCoord*12 + 12, 5)), 5));            
            break;
        case 11: //waves, camera
            c = vec3(waves(5, camera[0].y * 2, 5, 5));
            break;
        case 12: //fire
            c = vec3(fire(TexCoord));
            break;
        case 13: //cartoon fire
            c = sparks();            
            break;
        case 14: //cartoon fire
            c = cartoonFire(TexCoord);            
            break;
        case 15: //depth
            c = vec3(depth());
            break;
        case 16: //sci
            c = vec3(depth());
            break;
        case 17: //mouse
            c = vec3(julia(mouse*2-0.5, 100), 0);
            break;
        case 18: //julia fractal
            c = vec3(julia2((mouse*4-0.5)*mix(vec2(perlin(TexCoord*12+time*0.4, 7)), TexCoord, 0.9), 12));
            break;
        default:
            c = julia1(mouse*2-0.5, 100);
    }
    frag_colour = vec4(c,1);
}
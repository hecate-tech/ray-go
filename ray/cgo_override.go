package ray

// DO NOT DELETE THIS FILE FOR AUTO-GENERATION. THIS IS A MANUAL FILE!

// cgo_override is used for any functions that the generator cannot handle on
// its own.

// #define LIGHTGRAY  CLITERAL(Color){ 200, 200, 200, 255 }   // Light Gray
// #define GRAY       CLITERAL(Color){ 130, 130, 130, 255 }   // Gray
// #define DARKGRAY   CLITERAL(Color){ 80, 80, 80, 255 }      // Dark Gray
// #define YELLOW     CLITERAL(Color){ 253, 249, 0, 255 }     // Yellow
// #define GOLD       CLITERAL(Color){ 255, 203, 0, 255 }     // Gold
// #define ORANGE     CLITERAL(Color){ 255, 161, 0, 255 }     // Orange
// #define PINK       CLITERAL(Color){ 255, 109, 194, 255 }   // Pink
// #define RED        CLITERAL(Color){ 230, 41, 55, 255 }     // Red
// #define MAROON     CLITERAL(Color){ 190, 33, 55, 255 }     // Maroon
// #define GREEN      CLITERAL(Color){ 0, 228, 48, 255 }      // Green
// #define LIME       CLITERAL(Color){ 0, 158, 47, 255 }      // Lime
// #define DARKGREEN  CLITERAL(Color){ 0, 117, 44, 255 }      // Dark Green
// #define SKYBLUE    CLITERAL(Color){ 102, 191, 255, 255 }   // Sky Blue
// #define BLUE       CLITERAL(Color){ 0, 121, 241, 255 }     // Blue
// #define DARKBLUE   CLITERAL(Color){ 0, 82, 172, 255 }      // Dark Blue
// #define PURPLE     CLITERAL(Color){ 200, 122, 255, 255 }   // Purple
// #define VIOLET     CLITERAL(Color){ 135, 60, 190, 255 }    // Violet
// #define DARKPURPLE CLITERAL(Color){ 112, 31, 126, 255 }    // Dark Purple
// #define BEIGE      CLITERAL(Color){ 211, 176, 131, 255 }   // Beige
// #define BROWN      CLITERAL(Color){ 127, 106, 79, 255 }    // Brown
// #define DARKBROWN  CLITERAL(Color){ 76, 63, 47, 255 }      // Dark Brown

// Colors must be done separately, because the auto generator can't handle the
// CLITERAL Macro and then translate them into an enum.

var (
// TODO add rest of the colors.
)

var (
	White    = Color{R: 255, G: 255, B: 255, A: 255}
	Black    = Color{R: 0, G: 0, B: 0, A: 255}
	Blank    = Color{R: 0, G: 0, B: 0, A: 0}
	Magenta  = Color{R: 255, G: 0, B: 255, A: 255}
	RayWhite = Color{R: 245, G: 245, B: 245, A: 255}
)

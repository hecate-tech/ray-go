// +build android

package ray

// Credit to gen2brain for the base of these compatibility .go files.

/*
#include "external/android/native_app_glue/android_native_app_glue.c"
#cgo android LDFLAGS: -llog -landroid -lEGL -lGLESv2 -lOpenSLES -lm
#cgo android CFLAGS: -DPLATFORM_ANDROID -DGRAPHICS_API_OPENGL_ES2 -DMAL_NO_SDL -Iexternal/android/native_app_glue
*/
import "C"

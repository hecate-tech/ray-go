// +build darwin

package ray

// Credit to gen2brain for the base of these compatibility .go files.

/*
#cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation
#cgo darwin CFLAGS: -x objective-c -Iexternal/glfw/include -D_GLFW_COCOA -D_GLFW_USE_CHDIR -D_GLFW_USE_MENUBAR -D_GLFW_USE_RETINA -Wno-deprecated-declarations -DPLATFORM_DESKTOP -DMAL_NO_COREAUDIO
#cgo darwin,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo darwin,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo darwin,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33
*/
import "C"

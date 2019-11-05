// +build windows

package ray

// Credit to gen2brain for the base of these compatibility .go files.

/*
#cgo windows LDFLAGS: -lopengl32 -lgdi32 -lwinmm -lole32
#cgo windows CFLAGS: -D_GLFW_WIN32 -Iexternal -Iexternal/glfw/include -Iexternal/glfw/deps/mingw -DPLATFORM_DESKTOP
#cgo windows,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo windows,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo windows,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33
*/
import "C"

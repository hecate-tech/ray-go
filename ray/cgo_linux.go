// +build linux,!arm,!arm64

package ray

// Credit to gen2brain for the base of these compatibility .go files.

/*
#cgo linux CFLAGS: -Iexternal/glfw/include -DPLATFORM_DESKTOP -Wno-stringop-overflow
#cgo linux,!wayland LDFLAGS: -lGL -lm -pthread -ldl -lrt -lX11
#cgo linux,wayland LDFLAGS: -lGL -lm -pthread -ldl -lrt -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon
#cgo linux,!wayland CFLAGS: -D_GLFW_X11
#cgo linux,wayland CFLAGS: -D_GLFW_WAYLAND
#cgo linux,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo linux,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo linux,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33
*/
import "C"
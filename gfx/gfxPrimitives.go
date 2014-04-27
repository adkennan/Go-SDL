package gfx

// #cgo pkg-config: sdl
// #cgo LDFLAGS: -lSDL_gfx
// #include "SDL.h"
// #include "SDL_gfxPrimitives.h"
import "C"
import "github.com/adkennan/Go-SDL/sdl"
import "unsafe"

type cast unsafe.Pointer

// Pixel draw with blending enabled if a<255.
func PixelColor(dst *sdl.Surface, x, y int16, color uint32) int {
	return int(
		C.pixelColor((*C.SDL_Surface)(cast(dst)),
			C.Sint16(x),
			C.Sint16(y),
			C.Uint32(color)))
}

// Pixel draw with blending enabled if a<255.
func PixelRGBA(dst *sdl.Surface, x, y int16, r, g, b, a uint8) int {
	return int(
		C.pixelRGBA((*C.SDL_Surface)(cast(dst)),
			C.Sint16(x),
			C.Sint16(y),
			C.Uint8(r),
			C.Uint8(g),
			C.Uint8(b),
			C.Uint8(a)))
}

// Draw horizontal line with blending.
func HLineColor(dst *sdl.Surface, x1, x2, y int16, color uint32) int {
	return int(
		C.hlineColor((*C.SDL_Surface)(cast(dst)),
			C.Sint16(x1),
			C.Sint16(x2),
			C.Sint16(y),
			C.Uint32(color)))
}

// Draw horizontal line with blending.
func HLineRGBA(dst *sdl.Surface, x1, x2, y int16, r, g, b, a uint8) int {
	return int(
		C.hlineRGBA((*C.SDL_Surface)(cast(dst)),
			C.Sint16(x1),
			C.Sint16(x2),
			C.Sint16(y),
			C.Uint8(r),
			C.Uint8(g),
			C.Uint8(b),
			C.Uint8(a)))
}

// Draw vertical line with blending.
func VLineColor(dst *sdl.Surface, x, y1, y2 int16, color uint32) int {
	return int(
		C.vlineColor((*C.SDL_Surface)(cast(dst)),
			C.Sint16(x),
			C.Sint16(y1),
			C.Sint16(y2),
			C.Uint32(color)))
}

//int vlineRGBA (SDL_Surface *dst, Sint16 x, Sint16 y1, Sint16 y2, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw vertical line with blending.
func VLineRGBA(dst *sdl.Surface, x, y1, y2 int16, r, g, b, a uint8) int {
	return int(
		C.vlineRGBA((*C.SDL_Surface)(cast(dst)),
			C.Sint16(x),
			C.Sint16(y1),
			C.Sint16(y2),
			C.Uint8(r),
			C.Uint8(g),
			C.Uint8(b),
			C.Uint8(a)))
}

//int rectangleColor (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Uint32 color)
// Draw rectangle with blending.
//
//int rectangleRGBA (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw rectangle with blending.
//
func RectangleRGBA(dst *sdl.Surface, x1, y1, x2, y2 int16, r, g, b, a uint8) int {
	return int(
		C.rectangleRGBA((*C.SDL_Surface)(cast(dst)),
			C.Sint16(x1),
			C.Sint16(y1),
			C.Sint16(x2),
			C.Sint16(y2),
			C.Uint8(r),
			C.Uint8(g),
			C.Uint8(b),
			C.Uint8(a)))
}

//int roundedRectangleColor (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Sint16 rad, Uint32 color)
// Draw rounded-corner rectangle with blending.
//
//int roundedRectangleRGBA (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Sint16 rad, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw rounded-corner rectangle with blending.
//
//int boxColor (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Uint32 color)
// Draw box (filled rectangle) with blending.
//
//int boxRGBA (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw box (filled rectangle) with blending.
func BoxRGBA(dst *sdl.Surface, x1, y1, x2, y2 int16, r, g, b, a uint8) int {
	return int(
		C.boxRGBA((*C.SDL_Surface)(cast(dst)),
			C.Sint16(x1),
			C.Sint16(y1),
			C.Sint16(x2),
			C.Sint16(y2),
			C.Uint8(r),
			C.Uint8(g),
			C.Uint8(b),
			C.Uint8(a)))
}

//int roundedBoxColor (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Sint16 rad, Uint32 color)
// Draw rounded-corner box (filled rectangle) with blending.
//
//int roundedBoxRGBA (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Sint16 rad, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw rounded-corner box (filled rectangle) with blending.
func RoundedBoxRGBA(dst *sdl.Surface, x1, y1, x2, y2, rad int16, r, g, b, a uint8) int {
	return int(
		C.roundedBoxRGBA((*C.SDL_Surface)(cast(dst)),
			C.Sint16(x1),
			C.Sint16(y1),
			C.Sint16(x2),
			C.Sint16(y2),
			C.Sint16(rad),
			C.Uint8(r),
			C.Uint8(g),
			C.Uint8(b),
			C.Uint8(a)))
}

//int lineColor (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Uint32 color)
// Draw line with alpha blending.
//
func LineColor(dst *sdl.Surface, x1, y1, x2, y2 int16, color uint32) int {
	return int(
		C.lineColor((*C.SDL_Surface)(cast(dst)),
			C.Sint16(x1),
			C.Sint16(y1),
			C.Sint16(x2),
			C.Sint16(y2),
			C.Uint32(color)))
}

//int lineRGBA (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw line with alpha blending.
//
//int aalineColor (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Uint32 color)
// Ddraw anti-aliased line with alpha blending.
//
//int aalineRGBA (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw anti-aliased line with alpha blending.
//
//int thickLineColor (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Uint8 width, Uint32 color)
// Draw a thick line with alpha blending.
//
//int thickLineRGBA (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Uint8 width, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw a thick line with alpha blending.
//
//int circleColor (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Uint32 color)
// Draw circle with blending.
//
//int circleRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw circle with blending.
//
//int arcColor (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Sint16 start, Sint16 end, Uint32 color)
// Arc with blending.
//
//int arcRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Sint16 start, Sint16 end, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Arc with blending.
//
//int aacircleColor (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Uint32 color)
// Draw anti-aliased circle with blending.
//
//int aacircleRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw anti-aliased circle with blending.
//
//int filledCircleColor (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 r, Uint32 color)
// Draw filled circle with blending.
//
//int filledCircleRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw filled circle with blending.
//
//int ellipseColor (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rx, Sint16 ry, Uint32 color)
// Draw ellipse with blending.
func EllipseColor(dst *sdl.Surface, x, y, rx, ry int16, color uint32) int {
	return int(C.ellipseColor((*C.SDL_Surface)(cast(dst)),
		C.Sint16(x),
		C.Sint16(y),
		C.Sint16(rx),
		C.Sint16(ry),
		C.Uint32(color)))
}

//
//int ellipseRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rx, Sint16 ry, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw ellipse with blending.
func EllipseRGBA(dst *sdl.Surface, x, y, rx, ry int16, r, g, b, a uint8) int {
	return int(C.ellipseRGBA((*C.SDL_Surface)(cast(dst)),
		C.Sint16(x),
		C.Sint16(y),
		C.Sint16(rx),
		C.Sint16(ry),
		C.Uint8(r),
		C.Uint8(g),
		C.Uint8(b),
		C.Uint8(a)))
}

//int aaellipseColor (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rx, Sint16 ry, Uint32 color)
// Draw anti-aliased ellipse with blending.
//
//int aaellipseRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rx, Sint16 ry, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw anti-aliased ellipse with blending.
//
//int filledEllipseColor (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rx, Sint16 ry, Uint32 color)
// Draw filled ellipse with blending.
//
//int filledEllipseRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rx, Sint16 ry, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw filled ellipse with blending.
func FilledEllipseRGBA(dst *sdl.Surface, x, y, rx, ry int16, r, g, b, a uint8) int {
	return int(C.filledEllipseRGBA((*C.SDL_Surface)(cast(dst)),
		C.Sint16(x),
		C.Sint16(y),
		C.Sint16(rx),
		C.Sint16(ry),
		C.Uint8(r),
		C.Uint8(g),
		C.Uint8(b),
		C.Uint8(a)))
}

//
//int pieColor (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Sint16 start, Sint16 end, Uint32 color)
// Draw pie (outline) with alpha blending.
//
//int pieRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Sint16 start, Sint16 end, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw pie (outline) with alpha blending.
//
//int filledPieColor (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Sint16 start, Sint16 end, Uint32 color)
// Draw filled pie with alpha blending.
//
//int filledPieRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, Sint16 rad, Sint16 start, Sint16 end, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw filled pie with alpha blending.
//
//int trigonColor (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Sint16 x3, Sint16 y3, Uint32 color)
// Draw trigon (triangle outline) with alpha blending.
//
//int trigonRGBA (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Sint16 x3, Sint16 y3, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw trigon (triangle outline) with alpha blending.
//
//int aatrigonColor (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Sint16 x3, Sint16 y3, Uint32 color)
// Draw anti-aliased trigon (triangle outline) with alpha blending.
//
//int aatrigonRGBA (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Sint16 x3, Sint16 y3, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw anti-aliased trigon (triangle outline) with alpha blending.
//
//int filledTrigonColor (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Sint16 x3, Sint16 y3, Uint32 color)
// Draw filled trigon (triangle) with alpha blending.
//
//int filledTrigonRGBA (SDL_Surface *dst, Sint16 x1, Sint16 y1, Sint16 x2, Sint16 y2, Sint16 x3, Sint16 y3, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw filled trigon (triangle) with alpha blending.
//
//int polygonColor (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, Uint32 color)
// Draw polygon with alpha blending.
//
//int polygonRGBA (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw polygon with alpha blending.
func PolygonRGBA(dst *sdl.Surface, vx []int16, vy []int16, r, g, b, a uint8) int {
	return int(C.polygonRGBA((*C.SDL_Surface)(cast(dst)),
		(*C.Sint16)(unsafe.Pointer(&vx[0])),
		(*C.Sint16)(unsafe.Pointer(&vy[0])),
		C.int(len(vx)),
		C.Uint8(r),
		C.Uint8(g),
		C.Uint8(b),
		C.Uint8(a)))
}

//int aapolygonColor (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, Uint32 color)
// Draw anti-aliased polygon with alpha blending.
//
//int aapolygonRGBA (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw anti-aliased polygon with alpha blending.
//
//int filledPolygonColor (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, Uint32 color)
// Draw filled polygon with alpha blending.
//
//int filledPolygonRGBA (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw filled polygon with alpha blending.
func FilledPolygonRGBA(dst *sdl.Surface, vx []int16, vy []int16, r, g, b, a uint8) int {
	return int(C.filledPolygonRGBA((*C.SDL_Surface)(cast(dst)),
		(*C.Sint16)(cast(&vx[0])),
		(*C.Sint16)(cast(&vy[0])),
		C.int(len(vx)),
		C.Uint8(r),
		C.Uint8(g),
		C.Uint8(b),
		C.Uint8(a)))
}

//int texturedPolygon (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, SDL_Surface *texture, int texture_dx, int texture_dy)
// Draws a polygon filled with the given texture.
//
//int filledPolygonColorMT (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, Uint32 color, int **polyInts, int *polyAllocated)
// Draw filled polygon with alpha blending (multi-threaded capable).
//
//int filledPolygonRGBAMT (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, Uint8 r, Uint8 g, Uint8 b, Uint8 a, int **polyInts, int *polyAllocated)
// Draw filled polygon with alpha blending (multi-threaded capable).
//
//int texturedPolygonMT (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, SDL_Surface *texture, int texture_dx, int texture_dy, int **polyInts, int *polyAllocated)
// Draws a polygon filled with the given texture (Multi-Threading Capable).
//
//int bezierColor (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, int s, Uint32 color)
// Draw a bezier curve with alpha blending.
//
//int bezierRGBA (SDL_Surface *dst, const Sint16 *vx, const Sint16 *vy, int n, int s, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw a bezier curve with alpha blending.
//
//void gfxPrimitivesSetFont (const void *fontdata, Uint32 cw, Uint32 ch)
// Sets or resets the current global font data.
//
//void gfxPrimitivesSetFontRotation (Uint32 rotation)
// Sets current global font character rotation steps.
//
//int characterColor (SDL_Surface *dst, Sint16 x, Sint16 y, char c, Uint32 color)
// Draw a character of the currently set font.
//
//int characterRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, char c, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw a character of the currently set font.
//
//int stringColor (SDL_Surface *dst, Sint16 x, Sint16 y, const char *s, Uint32 color)
// Draw a string in the currently set font.
//
//int stringRGBA (SDL_Surface *dst, Sint16 x, Sint16 y, const char *s, Uint8 r, Uint8 g, Uint8 b, Uint8 a)
// Draw a string in the currently set font.

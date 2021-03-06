package isp

import "github.com/veandco/go-sdl2/sdl"

var pixelformats = map[string]uint32{
	"UNKNOWN":     sdl.PIXELFORMAT_UNKNOWN,
	"INDEX1LSB":   sdl.PIXELFORMAT_INDEX1LSB,
	"INDEX1MSB":   sdl.PIXELFORMAT_INDEX1MSB,
	"INDEX4LSB":   sdl.PIXELFORMAT_INDEX4LSB,
	"INDEX4MSB":   sdl.PIXELFORMAT_INDEX4MSB,
	"INDEX8":      sdl.PIXELFORMAT_INDEX8,
	"RGB332":      sdl.PIXELFORMAT_RGB332,
	"RGB444":      sdl.PIXELFORMAT_RGB444,
	"RGB555":      sdl.PIXELFORMAT_RGB555,
	"BGR555":      sdl.PIXELFORMAT_BGR555,
	"ARGB4444":    sdl.PIXELFORMAT_ARGB4444,
	"RGBA4444":    sdl.PIXELFORMAT_RGBA4444,
	"ABGR4444":    sdl.PIXELFORMAT_ABGR4444,
	"BGRA4444":    sdl.PIXELFORMAT_BGRA4444,
	"ARGB1555":    sdl.PIXELFORMAT_ARGB1555,
	"RGBA5551":    sdl.PIXELFORMAT_RGBA5551,
	"ABGR1555":    sdl.PIXELFORMAT_ABGR1555,
	"BGRA5551":    sdl.PIXELFORMAT_BGRA5551,
	"RGB565":      sdl.PIXELFORMAT_RGB565,
	"BGR565":      sdl.PIXELFORMAT_BGR565,
	"RGB24":       sdl.PIXELFORMAT_RGB24,
	"BGR24":       sdl.PIXELFORMAT_BGR24,
	"RGB888":      sdl.PIXELFORMAT_RGB888,
	"RGBX8888":    sdl.PIXELFORMAT_RGBX8888,
	"BGR888":      sdl.PIXELFORMAT_BGR888,
	"BGRX8888":    sdl.PIXELFORMAT_BGRX8888,
	"ARGB8888":    sdl.PIXELFORMAT_ARGB8888,
	"RGBA8888":    sdl.PIXELFORMAT_RGBA8888,
	"ABGR8888":    sdl.PIXELFORMAT_ABGR8888,
	"BGRA8888":    sdl.PIXELFORMAT_BGRA8888,
	"ARGB2101010": sdl.PIXELFORMAT_ARGB2101010,
	"YV12":        sdl.PIXELFORMAT_YV12,
	"IYUV":        sdl.PIXELFORMAT_IYUV,
	"YUY2":        sdl.PIXELFORMAT_YUY2,
	"UYVY":        sdl.PIXELFORMAT_UYVY,
	"YVYU":        sdl.PIXELFORMAT_YVYU,
}

func SDLPixelToString(format uint32) string {
	for k, v := range pixelformats {
		if v == format {
			return k
		}
	}
	return ""
}
func SDLPixelFromString(format string) uint32 {
	return pixelformats[format]
}

func SDLPixelList(format string) []string {
	var temp = make([]string, len(pixelformats))
	var i = 0
	for k := range pixelformats{
		temp[i] = k

	}
	return temp
}

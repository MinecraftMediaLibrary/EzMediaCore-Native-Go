package dither

import (
	"C"
	"github.com/MinecraftMediaLibrary/EzMediaCore-Native-Go/dither/utils"
	"unsafe"
)

func randomDither(
	colors []int,
	fullColors []byte,
	buffer []int,
	width int) unsafe.Pointer {
	var length = len(buffer)
	var height = length / width
	data := make([]byte, length)
	for y := 0; y < height; y++ {
		var yIndex = y * width
		for x := 0; x < width; x++ {
			var index = yIndex + x
			var color = buffer[index]
			data[index] =
				utils.GetBestColorSeparateRGB(
					fullColors,
					((color>>16)&0xFF)+utils.GenerateRandom(-64, 64),
					((color>>8)&0xFF)+utils.GenerateRandom(-64, 64),
					((color)&0xFF)+utils.GenerateRandom(-64, 64),
				)
		}
	}
	return C.CBytes(data)
}

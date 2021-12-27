package dither

import (
	"C"
	"ezmediacore-native-go/ezmediacore-native-go/dither/utils"
	"unsafe"
)

func simpleDither(
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
			data[yIndex+x] = utils.GetBestColorRGB(fullColors, buffer[yIndex+x])
		}
	}
	return C.CBytes(data)
}

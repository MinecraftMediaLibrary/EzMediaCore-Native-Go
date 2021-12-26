package dither

import (
	"C"
	"ezmediacore-native-go/main/dither/utils"
	"unsafe"
)

func filterLiteDither(
	colors []int,
	fullColors []byte,
	buffer []int,
	width int) unsafe.Pointer {
	var length = len(buffer)
	var height = length / width
	var widthMinus = width - 1
	var heightMinus = height - 1
	var arrayWidth = width + width<<1
	data := make([]byte, length)
	ditherBuffer := make([][]int, 2)
	for i := 0; i < 2; i++ {
		ditherBuffer[i] = make([]int, arrayWidth)
	}
	for y := 0; y < height; y++ {
		var hasNextY = y < heightMinus
		var yIndex = y * width
		if y&0x1 == 0 {
			var bufferIndex = 0
			var buf1 = ditherBuffer[0]
			var buf2 = ditherBuffer[1]
			for x := 0; x < width; x++ {
				var index = yIndex + x
				var rgb = buffer[index]
				var red = rgb >> 16 & 0xFF
				var green = rgb >> 8 & 0xFF
				var blue = rgb & 0xFF

				bufferIndex, red = utils.CalculateIncrementalPixel(bufferIndex, red, buf1)
				bufferIndex, green = utils.CalculateIncrementalPixel(bufferIndex, green, buf1)
				bufferIndex, blue = utils.CalculateIncrementalPixel(bufferIndex, blue, buf1)

				var closest = utils.GetBestFullColor(colors, red, green, blue)
				var deltaR = red - (closest >> 16 & 0xFF)
				var deltaG = green - (closest >> 8 & 0xFF)
				var deltaB = blue - (closest & 0xFF)

				if x < widthMinus {
					buf1[bufferIndex] = deltaR >> 1
					buf1[bufferIndex+1] = deltaG >> 1
					buf1[bufferIndex+2] = deltaB >> 1
				}

				if hasNextY {
					if x > 0 {
						buf2[bufferIndex-6] = deltaR >> 2
						buf2[bufferIndex-5] = deltaG >> 2
						buf2[bufferIndex-4] = deltaB >> 2
					}
					buf2[bufferIndex-3] = deltaR >> 2
					buf2[bufferIndex-2] = deltaG >> 2
					buf2[bufferIndex-1] = deltaB >> 2
				}

				data[index] = utils.GetBestColorRGB(fullColors, closest)
			}
		} else {
			var bufferIndex = width + (width << 1) - 1
			var buf1 = ditherBuffer[1]
			var buf2 = ditherBuffer[0]
			for x := width - 1; x >= 0; x-- {
				var index = yIndex + x
				var rgb = buffer[index]
				var red = rgb >> 16 & 0xFF
				var green = rgb >> 8 & 0xFF
				var blue = rgb & 0xFF

				bufferIndex, blue = utils.CalculateDecrementalPixel(bufferIndex, blue, buf1)
				bufferIndex, green = utils.CalculateDecrementalPixel(bufferIndex, green, buf1)
				bufferIndex, red = utils.CalculateDecrementalPixel(bufferIndex, red, buf1)

				var closest = utils.GetBestFullColor(colors, red, green, blue)
				var deltaR = red - (closest >> 16 & 0xFF)
				var deltaG = green - (closest >> 8 & 0xFF)
				var deltaB = blue - (closest & 0xFF)

				if x > 0 {
					buf1[bufferIndex] = deltaB >> 1
					buf1[bufferIndex-1] = deltaG >> 1
					buf1[bufferIndex-2] = deltaR >> 1
				}

				if hasNextY {
					if x < widthMinus {
						buf2[bufferIndex+6] = deltaB >> 2
						buf2[bufferIndex+5] = deltaG >> 2
						buf2[bufferIndex+4] = deltaR >> 2
					}
					buf2[bufferIndex+3] = deltaB >> 2
					buf2[bufferIndex+2] = deltaG >> 2
					buf2[bufferIndex+1] = deltaR >> 2
				}
				data[index] = utils.GetBestColorRGB(fullColors, closest)
			}
		}
	}
	return C.CBytes(data)
}

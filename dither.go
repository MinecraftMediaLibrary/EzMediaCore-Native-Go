package main

//// formatter: off

import (
	"C"
	"math/rand"
	"unsafe"
)

//// formatter: on

func main() {}

//export filterLiteDither
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

				bufferIndex, red = calculateIncrementalPixel(bufferIndex, red, buf1)
				bufferIndex, green = calculateIncrementalPixel(bufferIndex, green, buf1)
				bufferIndex, blue = calculateIncrementalPixel(bufferIndex, blue, buf1)

				var closest = getBestFullColor(colors, red, green, blue)
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

				data[index] = getBestColorRGB(fullColors, closest)
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

				bufferIndex, blue = calculateDecrementalPixel(bufferIndex, blue, buf1)
				bufferIndex, green = calculateDecrementalPixel(bufferIndex, green, buf1)
				bufferIndex, red = calculateDecrementalPixel(bufferIndex, red, buf1)

				var closest = getBestFullColor(colors, red, green, blue)
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
				data[index] = getBestColorRGB(fullColors, closest)
			}
		}
	}
	return C.CBytes(data)
}

//export floydSteinbergDither
func floydSteinbergDither(
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
				var hasNextX = x < widthMinus
				var index = yIndex + x
				var rgb = buffer[index]
				var red = rgb >> 16 & 0xFF
				var green = rgb >> 8 & 0xFF
				var blue = rgb & 0xFF

				bufferIndex, red = calculateIncrementalPixel(bufferIndex, red, buf1)
				bufferIndex, green = calculateIncrementalPixel(bufferIndex, green, buf1)
				bufferIndex, blue = calculateIncrementalPixel(bufferIndex, blue, buf1)

				var closest = getBestFullColor(colors, red, green, blue)
				var deltaR = red - (closest >> 16 & 0xFF)
				var deltaG = green - (closest >> 8 & 0xFF)
				var deltaB = blue - (closest & 0xFF)

				if hasNextX {
					buf1[bufferIndex] = int(0.4375 * float32(deltaR))
					buf1[bufferIndex+1] = int(0.4375 * float32(deltaG))
					buf1[bufferIndex+2] = int(0.4375 * float32(deltaB))
				}
				if hasNextY {
					if x > 0 {
						buf2[bufferIndex-6] = int(0.1875 * float32(deltaR))
						buf2[bufferIndex-5] = int(0.1875 * float32(deltaG))
						buf2[bufferIndex-4] = int(0.1875 * float32(deltaB))
					}
					buf2[bufferIndex-3] = int(0.3125 * float32(deltaR))
					buf2[bufferIndex-2] = int(0.3125 * float32(deltaG))
					buf2[bufferIndex-1] = int(0.3125 * float32(deltaB))
					if hasNextX {
						buf2[bufferIndex] = int(0.0625 * float32(deltaR))
						buf2[bufferIndex+1] = int(0.0625 * float32(deltaG))
						buf2[bufferIndex+2] = int(0.0625 * float32(deltaB))
					}
				}

				data[index] = getBestColorRGB(fullColors, closest)
			}
		} else {
			var bufferIndex = width + (width << 1) - 1
			var buf1 = ditherBuffer[1]
			var buf2 = ditherBuffer[0]
			for x := width - 1; x >= 0; x-- {
				var hasNextX = x > 0
				var index = yIndex + x
				var rgb = buffer[index]
				var red = rgb >> 16 & 0xFF
				var green = rgb >> 8 & 0xFF
				var blue = rgb & 0xFF

				bufferIndex, blue = calculateDecrementalPixel(bufferIndex, blue, buf1)
				bufferIndex, green = calculateDecrementalPixel(bufferIndex, green, buf1)
				bufferIndex, red = calculateDecrementalPixel(bufferIndex, red, buf1)

				var closest = getBestFullColor(colors, red, green, blue)
				var deltaR = red - (closest >> 16 & 0xFF)
				var deltaG = green - (closest >> 8 & 0xFF)
				var deltaB = blue - (closest & 0xFF)

				if hasNextX {
					buf1[bufferIndex] = int(0.4375 * float32(deltaR))
					buf1[bufferIndex-1] = int(0.4375 * float32(deltaG))
					buf1[bufferIndex-2] = int(0.4375 * float32(deltaB))
				}
				if hasNextY {
					if x > 0 {
						buf2[bufferIndex+6] = int(0.1875 * float32(deltaR))
						buf2[bufferIndex+5] = int(0.1875 * float32(deltaG))
						buf2[bufferIndex+4] = int(0.1875 * float32(deltaB))
					}
					buf2[bufferIndex+3] = int(0.3125 * float32(deltaR))
					buf2[bufferIndex+2] = int(0.3125 * float32(deltaG))
					buf2[bufferIndex+1] = int(0.3125 * float32(deltaB))
					if hasNextX {
						buf2[bufferIndex] = int(0.0625 * float32(deltaR))
						buf2[bufferIndex-1] = int(0.0625 * float32(deltaG))
						buf2[bufferIndex-2] = int(0.0625 * float32(deltaB))
					}
				}

				data[index] = getBestColorRGB(fullColors, closest)
			}
		}
	}
	return C.CBytes(data)
}

//export randomDither
func randomDither(
	colors []int,
	fullColors []byte,
	buffer []int,
	width, weight int) unsafe.Pointer {
	var min = -weight
	var length = len(buffer)
	var height = length / width
	data := make([]byte, length)
	for y := 0; y < height; y++ {
		var yIndex = y * width
		for x := 0; x < width; x++ {
			var index = yIndex + x
			var color = buffer[index]
			var r = (color >> 16) & 0xFF
			var g = (color >> 8) & 0xFF
			var b = (color) & 0xFF
			r = correctPixel(r + generateRandom(min, weight))
			g = correctPixel(g + generateRandom(min, weight))
			b = correctPixel(b + generateRandom(min, weight))
			data[index] = getBestColorSeparateRGB(fullColors, r, g, b)
		}
	}
	return C.CBytes(data)
}

//export simpleDither
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
			data[yIndex+x] = getBestColorRGB(fullColors, buffer[yIndex+x])
		}
	}
	return C.CBytes(data)
}

func calculateIncrementalPixel(bufferIndex int, color int, buf1 []int) (int, int) {
	bufferIndex++
	return calculatePixel(bufferIndex, color, buf1)
}

func calculateDecrementalPixel(bufferIndex int, color int, buf1 []int) (int, int) {
	bufferIndex--
	return calculatePixel(bufferIndex, color, buf1)
}

func calculatePixel(bufferIndex int, color int, buf1 []int) (int, int) {
	color += buf1[bufferIndex]
	color = correctPixel(color)
	return bufferIndex, color
}

func correctPixel(color int) int {
	if color > 255 {
		color = 255
	} else if color < 0 {
		color = 0
	}
	return color
}

func generateRandom(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func getBestColorRGB(fullColors []byte, rgb int) byte {
	return fullColors[rgb>>1<<14|rgb>>1<<7|rgb>>1]
}

func getBestColorSeparateRGB(fullColors []byte, red int, green int, blue int) byte {
	return fullColors[red>>1<<14|green>>1<<7|blue>>1]
}

func getBestFullColor(colors []int, red int, green int, blue int) int {
	return colors[red>>1<<14|green>>1<<7|blue>>1]
}

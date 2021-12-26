package utils

func CalculateIncrementalPixel(bufferIndex int, color int, buf1 []int) (int, int) {
	bufferIndex++
	return calculatePixel(bufferIndex, color, buf1)
}

func CalculateDecrementalPixel(bufferIndex int, color int, buf1 []int) (int, int) {
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

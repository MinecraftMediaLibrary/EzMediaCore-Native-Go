package utils

import "math/rand"

func GenerateRandom(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func GetBestColorRGB(fullColors []byte, rgb int) byte {
	return fullColors[rgb>>1<<14|rgb>>1<<7|rgb>>1]
}

func GetBestColorSeparateRGB(fullColors []byte, red int, green int, blue int) byte {
	return fullColors[red>>1<<14|green>>1<<7|blue>>1]
}

func GetBestFullColor(colors []int, red int, green int, blue int) int {
	return colors[red>>1<<14|green>>1<<7|blue>>1]
}

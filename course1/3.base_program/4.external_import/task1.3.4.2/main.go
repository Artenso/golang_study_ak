package main

import (
	"fmt"

	"github.com/ksrof/gocolors"
)

func main() {
	str := "Rainbow"
	fmt.Println(ColorizeRed(str))
	fmt.Println(ColorizeCustom(str, 255, 165, 0))
	fmt.Println(ColorizeYellow(str))
	fmt.Println(ColorizeGreen(str))
	fmt.Println(ColorizeCyan(str))
	fmt.Println(ColorizeBlue(str))
	fmt.Println(ColorizeMagenta(str))
	fmt.Println(ColorizeWhite(str))
}

// Get your color with RGB values
func RGB(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func ColorizeRed(a string) string {
	return gocolors.Red(a, "")
}

func ColorizeGreen(a string) string {
	return gocolors.Green(a, "")
}

func ColorizeBlue(a string) string {
	return gocolors.Blue(a, "")
}

func ColorizeYellow(a string) string {
	return gocolors.Yellow(a, "")
}

func ColorizeMagenta(a string) string {
	return gocolors.Magenta(a, "")
}

func ColorizeCyan(a string) string {
	return gocolors.Cyan(a, "")
}

func ColorizeWhite(a string) string {
	return gocolors.White(a, "")
}

func ColorizeCustom(a string, r, g, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, a)
}

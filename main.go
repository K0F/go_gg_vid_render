package main

import "github.com/fogleman/gg"
import "fmt"
import "math"

var frameCount = 0
var noFrames = 25 * 100

func main() {
	dc := gg.NewContext(1920, 1080)

	for frameCount < noFrames {
		// background
		dc.DrawRectangle(0, 0, 1920, 1080)
		dc.SetRGB(0, 0, 0)
		dc.Fill()

		i := 1
		for i < 100 {
			val := math.Sin(float64(frameCount)/float64(noFrames)/float64(i)*math.Pi*2) / 2.0
			i++
			dc.DrawCircle(1920/2+val*400, 1080/2, 400-float64(i)*4)
			dc.SetRGB(1, 1, 1)
			dc.Stroke()
		}
		// dc.Fill()

		filename := fmt.Sprintf("%05d_out.png", frameCount)
		fmt.Println("Saving: " + filename)
		dc.SavePNG(filename)

		frameCount++
	}
}

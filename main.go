package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/remeh/sizedwaitgroup"
	"math"
)

var frameCount = 0
var noFrames = 25 * 100

func main() {
	swg := sizedwaitgroup.New(8)

	for frameCount < noFrames {

		dc := gg.NewContext(1920, 1080)
		swg.Add()

		go func(framecount int) {
			// background
			defer swg.Done()
			dc.DrawRectangle(0, 0, 1920, 1080)
			dc.SetRGB(0, 0, 0)
			dc.Fill()

			i := 1
			for i < 100 {
				val := math.Sin(float64(framecount)/float64(noFrames)/(float64(i)/10.0)*math.Pi*2) / 2.0
				i++
				dc.DrawCircle(1920/2+val*400, 1080/2, 400-float64(i)*4)
				dc.SetRGB(1, 1, 1)
				dc.Stroke()
			}
			// dc.Fill()

			filename := fmt.Sprintf("%05d_out.png", framecount)
			fmt.Println("Saving: " + filename)
			dc.SavePNG(filename)
		}(frameCount)

		frameCount++

	}

	swg.Wait()
	fmt.Printf("%d frames saved\n", frameCount)
	fmt.Println(`to create video run: ffmpeg -r 25 -i %05d_out.png -c:v h264 -pix_fmt yuv420p -y out.mp4 && mpv out.mp4`)
}

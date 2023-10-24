package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func getThumb(c echo.Context) error {
	err := ffmpeg.Input("./stream/hls/.m3u8", ffmpeg.KwArgs{"ss": "1"}).
		Output("./thumbs/bbb.gif", ffmpeg.KwArgs{"s": "320x240", "pix_fmt": "rgb24", "t": "3", "r": "3"}).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return err
	}
	return c.File("./thumbs/bbb.gif")

}

func main() {
	e := echo.New()
	// Routes
	e.GET("/thumb", getThumb)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

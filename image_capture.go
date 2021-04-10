package main

import (
	"fmt"
	"os"
	"strconv"

	"gocv.io/x/gocv"
	"github.com/otiai10/gosseract"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tsaveimage [file] [image dir]")
		return
	}
	file := os.Args[1]
	saveFile := os.Args[2]

	video, err := gocv.VideoCaptureFile(file)
	if err != nil {
		fmt.Printf("Error opening video capture file: %s\n", file)
		return
	}
	defer video.Close()

	client := gosseract.NewClient()
    defer client.Close()

	img := gocv.NewMat()
	defer img.Close()
	num := 0

    for {
        if ok := video.Read(&img); !ok {
            fmt.Printf("cannot read video\n")
            return
        }
        if img.Empty() {
            fmt.Printf("no image on device \n")
            return
        }

	    gocv.IMWrite(saveFile + strconv.Itoa(num) + ".jpg", img)
	    client.SetImage(saveFile + strconv.Itoa(num) + ".jpg")
	    num++

	    text, _ := client.Text()
	    fmt.Println(text)

	}


}
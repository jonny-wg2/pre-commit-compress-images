package main

import (
	"fmt"
	"os"

	"github.com/gographics/imagick/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	filename := os.Args[1]

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err := mw.ReadImage(filename)
	if err != nil {
		fmt.Println("Error reading image:", err)
		return
	}

	quality := 75

	err = mw.SetImageCompressionQuality(uint(quality))
	if err != nil {
		fmt.Println("Error setting compression quality:", err)
		return
	}

	err = mw.WriteImage(filename)
	if err != nil {
		fmt.Println("Error writing image:", err)
		return
	}

	fmt.Println("Image compressed successfully.")
}

package file

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/adzsx/stegopher/internal/utils"
)

var (
	img image.Image
)

func read(name string) {
	f, err := os.Open(name)
	utils.Err(err)
	defer f.Close()

	fType := utils.GetType(name)

	if fType == "png" {
		img, err = png.Decode(f)
		utils.Err(err)
		log.Println(img)
	} else if fType == "jpeg" {
		img, err = jpeg.Decode(f)
		utils.Err(err)
		log.Println(img)
	} else if fType == "gif" {
		img, err = gif.Decode(f)
		utils.Err(err)
	} else {
		log.Fatalln("Invalid file format")
	}
}

func Extract(input utils.Input) {
	read(input.Image)
}

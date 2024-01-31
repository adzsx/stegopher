package main

import (
	"log"
	"os"

	"github.com/adzsx/stegopher/internal/file"
	"github.com/adzsx/stegopher/internal/utils"
)

var (
	help string = `
stegopher usage:
	stegopher [flags]

Flags:
	-m MESSAE		Message to embed into image
	-i PATH			Input image
	-o PATH			Output image for embedding result
	`
)

func main() {
	log.SetFlags(0)

	args := os.Args

	input := utils.Format(args)

	if input.Help {
		log.Println(help)
		os.Exit(0)
	}

	file.Write(input)
}

package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Image   string
	Message string
	Output  string
	Help    bool
}

func Format(args []string) Input {

	input := Input{}

	if InSclice(args, "-h") || InSclice(args, "--help") {
		input.Help = true
		return input
	}

	if len(args) < 3 {
		log.Println("Not enough arguments\nEnter --help for help")
		os.Exit(0)
	}

	for index, element := range args[0:] {
		switch element[1:] {
		case "m":
			input.Message = args[index+1]
		case "i":
			input.Image = args[index+1]
		case "o":
			input.Output = args[index+1]
		}
	}

	if input.Image == "" {
		log.Fatalln("No input Image defined")
	}
	if input.Message == "" {
		log.Fatalln("No Message definied")
	}

	if input.Output == "" {
		f, err := os.Stat(input.Image)
		Err(err)
		name := f.Name()

		imgType := GetType(input.Image)
		name = strings.Replace(name, "."+imgType, "", 1)
		Err(err)
		output := name + "_out." + imgType

		counter := 0
		created := false

		for !created {
			counter += 1
			_, err = os.Open(output)
			if err != nil {
				created = true
			} else {
				if counter == 1 {
					output = strings.Replace(output, "."+imgType, "", 1) + "_" + strconv.Itoa(counter) + "." + imgType
				} else {
					output = strings.Replace(output, "_"+strconv.Itoa(counter-1)+"."+imgType, "", 1) + "_" + strconv.Itoa(counter) + "." + imgType
				}
			}
		}

		input.Output = output
	}

	return input
}

package utils

import (
	"os"
	"strings"
)

func Err(err error) {
	if err != nil {
		panic(err)
	}
}

func GetType(path string) string {
	stat, err := os.Stat(path)
	Err(err)

	img, err := os.Open(path)
	Err(err)
	defer img.Close()

	fType := stat.Name()[len(stat.Name())-3:]

	return strings.Replace(fType, "peg", "jpeg", 1)
}

func InSclice(s []string, element string) bool {
	for _, v := range s {
		if element == v {
			return true
		}
	}
	return false
}

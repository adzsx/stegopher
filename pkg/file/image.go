package file

import (
	"os"
	"strings"

	"github.com/adzsx/stegopher/internal/utils"
)

func GetType(path string) string {
	stat, err := os.Stat(path)
	utils.Err(err)

	img, err := os.Open(path)
	utils.Err(err)
	defer img.Close()

	fType := stat.Name()[len(stat.Name())-3:]

	return strings.Replace(fType, "peg", "jpeg", 1)
}

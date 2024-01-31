package file

import "github.com/adzsx/stegopher/internal/utils"

func Write(input utils.Input) {
	read(input.Image)
	print(img)
}

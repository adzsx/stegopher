package main

import "github.com/adzsx/stegopher/pkg/file"

func main() {
	f := "/home/adzsx/cs/security/tools/stegopher/image.png"
	file.Embed(f, "Test")
}

package main

import "fmt"
import "log"
import "os"
import "image"
import "image/png"

import "github.com/seletskiy/slorry/kpch/crack"

func OpenImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening %s: %s\n", path, err)
		return nil
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatalf("error decode image %s: %s\n", path, err)
		return nil
	}

	return img
}

func main() {
	path := os.Args[1]
	img := OpenImage(path)

	for _, s := range crack.Crack(img) {
		fmt.Print(s.Sym.Char)
	}

	fmt.Println()
}

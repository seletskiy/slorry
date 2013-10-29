package main

import "os"
import "log"
import "fmt"
import "path/filepath"
import "image"
//import "image/color"
import "image/png"

func ProcessFile(path string, info os.FileInfo, err error) error {
    if info.IsDir() {
        return nil
    }

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

    bounds := img.Bounds()

    dst := image.NewGray(image.Rect(0, 0, bounds.Dx() / 2, bounds.Dy() / 2))

    _ = dst

    for y := 0; y < bounds.Dy() - 1; y += 1 {
        for x := 0; x < bounds.Dx() - 1; x += 1 {
            r, _, _, _ := img.At(x, y).RGBA()
            if r == 0 {
                fmt.Print("X")
            } else {
                fmt.Print(" ")
            }
        }

        fmt.Println("")
    }

    log.Print(path)

    var value int

    fmt.Print(">>> ")
    fmt.Scanf("%d", &value)

    os.Rename(path, fmt.Sprintf("%d.png", value))

    return nil
}

func main() {
    filepath.Walk(os.Args[1], ProcessFile)
}

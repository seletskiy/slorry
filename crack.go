package main

import "log"
import "os"
import "image"
import "image/png"
import "fmt"

import . "symbols"
import . "symbols/maps"
import . "features"
import . "features/masks"

var _ = fmt.Print
var debug = false

func MatchFeatures(img image.Image, x int, y int) []*ConcreteFeature {
    features := []*Feature {
        &PipeLeft,
        &PipeRight,
        &CornerLeftBottom,
        &CornerRightBottom,
        &CornerRightTop,
        &CaveLeftSmall,
        &ArcTopLeft,
        &ArcTopRight,
        &LongHole,
        &DescentLeft,
        &CaveLeftBig,
        &AscentLeft,
        &BridgeLeft,
        &DescentLeftLong,
        &Snail,
        &MediumHole,
        &Pad,
    }

    result := make([]*ConcreteFeature, 0)

    for _, feature := range features {
        match := MatchFeature(feature, img, x, y)
        if match != nil {
            result = append(result, match)
        }
    }

    return result
}

func main() {
    path := os.Args[1]

    img := OpenImage(path)

    bounds := img.Bounds()

    matches := make([]*ConcreteFeature, 0)

    for y := 0; y < bounds.Dy(); y++ {
        for x := 0; x < bounds.Dx(); x++ {
            matches = append(matches, MatchFeatures(img, x, y)...)
        }
    }

    startX := -1
    endX := 0

    for y := 0; y < bounds.Dy(); y++ {
        for x := 0; x < bounds.Dx(); x++ {
            printed := false
            for _, m := range matches {
                if m.X == x && m.Y == y {
                    if debug {
                        fmt.Printf("%c", m.Symbol)
                    }
                    printed = true
                    break
                }
            }

            if printed {
                continue
            }

            r, g, b, _ := img.At(x, y).RGBA()
            if r + g + b == 0 {
                if startX == -1 {
                    startX = x
                } else if startX > x {
                    startX = x
                } else if x > endX {
                    endX = x
                }
                if debug {
                    fmt.Print(".")
                }
            } else {
                if debug {
                    fmt.Print(" ")
                }
            }
        }

        if debug {
            fmt.Println("")
        }
    }

    symbols := []Symbol {
        Zero,
        One,
        Two,
        Three,
        Four,
        Five,
        Six,
        Seven,
        Eight,
        Nine,
    }

    chunkSize := (endX - startX) / 5
    overlap := 4
    shift := 1

    var total FeatureSet

    for chunk := 0; chunk < 5; chunk++ {
        rangeStart := startX + chunk * chunkSize - overlap + chunk * shift
        rangeEnd := startX + (chunk + 1) * chunkSize + overlap + chunk * shift
        local := make([]*ConcreteFeature, 0)
        if debug {
            fmt.Printf("\n%d - %d\n", rangeStart, rangeEnd)
        }
        for _, m := range matches {
            if m.X >= rangeStart && m.X <= rangeEnd {
                local = append(local, m)
            }
        }

        max := 0.0
        var match Symbol
        var used FeatureSet

        for _, symbol := range symbols {
            weight, matched := symbol.Match(local, total)
            if debug {
                fmt.Printf("%s %3.2f\n", symbol.Char, weight)
            }
            if weight > max {
                max = weight
                match = symbol
                used = matched
            }
        }

        total = append(total, used...)

        if debug {
            fmt.Println("=======")
        }

        _ = match
        fmt.Printf("%s", match.Char)
    }

    fmt.Printf("\n")
}

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

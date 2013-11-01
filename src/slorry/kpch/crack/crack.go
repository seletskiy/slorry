package crack

import "fmt"
import "log"
import "image"
import "image/color"

import . "slorry/kpch/symbols"
import . "slorry/kpch/symbols/maps"
import . "slorry/kpch/features"
import . "slorry/kpch/features/patterns"

type SymbolMatch struct {
    Sym Symbol
    Weight float64
}


func Sharpify(img image.Image) *image.Gray {
    bounds := img.Bounds()
    dx := bounds.Dx()
    dy := bounds.Dy()

    gray := image.NewGray(bounds)

    for y := 0; y < dy; y++ {
        for x := 0; x < dx; x++ {
            if x == 0 || y == 0 || y == dy - 1 || x == dx - 1 {
                gray.Set(x, y, color.Gray{255})
                continue
            }

            weight := 0
            for i := -1; i <= 1; i += 1 {
                for j := -1; j <= 1; j += 1 {
                    col := img.At(x + j, y + i)
                    sat := int(gray.ColorModel().Convert(col).(color.Gray).Y)
                    sat = 255 - sat

                    if i == 0 && j == 0 {
                        weight += 9 * sat
                    } else {
                        weight -= sat
                    }
                }
            }

            pix := 0

            if weight < 0 {
                pix = 255
            } else if weight >= 150 {
                pix = 0
            } else {
                pix = 255
            }

            gray.Set(x, y, color.Gray{uint8(pix)})
        }
    }

    return gray
}

func FindFeatures(img image.Image, x int, y int) FeatureSet {
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
        &Cornice,
        &NotchTop,
        &NotchBottom,
    }

    result := make(FeatureSet, 0)

    for _, feature := range features {
        match := MatchFeature(feature, img, x, y)
        if match != nil {
            result = append(result, match)
        }
    }

    return result
}

func FindAllFeatures(img image.Image) FeatureSet {
    dx := img.Bounds().Dx()
    dy := img.Bounds().Dy()
    matches := make(FeatureSet, 0)
    for y := 0; y < dy; y++ {
        for x := 0; x < dx; x++ {
            matches = append(matches, FindFeatures(img, x, y)...)
        }
    }

    return matches
}

func PrintImage(img image.Image, features FeatureSet) {
    dx := img.Bounds().Dx()
    dy := img.Bounds().Dy()
    for y := 0; y < dy; y++ {
        row := ""
        for x := 0; x < dx; x++ {
            printed := false
            for _, m := range features {
                if m.X == x && m.Y == y {
                    row += fmt.Sprintf("%c", m.Symbol)
                    printed = true
                    break
                }
            }

            if printed {
                continue
            }

            sat := img.At(x, y).(color.Gray).Y
            if sat == 0 {
                row += "."
            } else {
                row += " "
            }
        }

        log.Println(row)
    }
}

func Chunkify(img image.Image, features FeatureSet,
        count int, overlap int, shift int) []FeatureSet {
    dx := img.Bounds().Dx()
    dy := img.Bounds().Dy()
    start := dx
    end := 0
    for x := 0; x < dx; x++ {
        for y := 0; y < dy; y++ {
            if img.At(x, y).(color.Gray).Y == 0 {
                if x < start {
                    start = x
                }
                
                if x > end {
                    end = x
                }
            }
        }
    }

    chunks := make([]FeatureSet, 0)
    size := (end - start) / count

    for chunk := 0; chunk < count; chunk++ {
        windowStart := start + chunk * size - overlap + chunk * shift
        windowEnd := start + (chunk + 1) * size + overlap + chunk * shift

        log.Printf(fmt.Sprintf(
            fmt.Sprintf("%%%ds%%%ds", windowStart, windowEnd - windowStart),
            "S", "E"))

        window := make(FeatureSet, 0)

        for _, f := range features {
            if f.X >= windowStart && f.X <= windowEnd {
                window = append(window, f)
            }
        }

        chunks = append(chunks, window)
    }

    return chunks
}

func FindSymbols(chunks []FeatureSet) []SymbolMatch {
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

    result := make([]SymbolMatch, 0)
    usedFeatures := make(FeatureSet, 0)

    for _, chunk := range chunks {
        var matchedSymbol Symbol
        var symbolFeatures FeatureSet

        maxWeight := 0.0

        for _, symbol := range symbols {
            weight, matchedFeatures := symbol.Match(chunk, usedFeatures)
            if weight > maxWeight {
                maxWeight = weight
                matchedSymbol = symbol
                symbolFeatures = matchedFeatures
            }

            log.Printf("::: %s: %3.2f", symbol.Char, weight)
        }

        log.Printf("=== %s: %3.2f", matchedSymbol.Char, maxWeight)
        result = append(result, SymbolMatch{matchedSymbol, maxWeight})
        usedFeatures = append(usedFeatures, symbolFeatures...)
    }

    return result
}

func Crack(source image.Image) []SymbolMatch {
    img := Sharpify(source)
    features := FindAllFeatures(img)
    chunks := Chunkify(img, features, 5, 4, 1)

    PrintImage(img, features)

    return FindSymbols(chunks)
}

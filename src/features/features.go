package features

import . "image"

type FeatureMask []string

type Feature struct {
    Id string
    Variants []FeatureMask
}

type ConcreteFeature struct {
    X int
    Y int
    Symbol uint8
    Origin *Feature
}

func MatchFeature(feature *Feature, img Image, x int, y int) *ConcreteFeature {
    for _, mask := range feature.Variants {
        concrete := MatchMask(mask, img, x, y)
        if concrete != nil {
            concrete.Origin = feature
            return concrete
        }
    }

    return nil
}

func MatchMask(mask FeatureMask, img Image, x int, y int) *ConcreteFeature {
    bx := x
    by := y
    bounds := img.Bounds()

    var symbol uint8 = 'X'

    for i := 0; i < len(mask); i++ {
        for j := 0; j < len(mask[i]); j++ {
            if mask[i][j] == ' ' {
                continue
            }

            dx := x + j
            dy := y + i

            if mask[i][j] >= 'A' && mask[i][j] <= 'Z' {
                bx = dx
                by = dy
                symbol = mask[i][j]
            }

            r, g, b, _ := img.At(dx, dy).RGBA()

            if dx < 0 || dy < 0 || dx > bounds.Dx() || dy > bounds.Dy() {
                return nil
            } else if r + g + b != 0 {
                return nil
            }
        }
    }

    return &ConcreteFeature{bx, by, symbol, nil}
}

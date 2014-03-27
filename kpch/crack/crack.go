package crack

import "fmt"
import "image"
import "image/color"

import "github.com/seletskiy/slorry/kpch/symbols"
import "github.com/seletskiy/slorry/kpch/symbols/maps"
import "github.com/seletskiy/slorry/kpch/features"
import "github.com/seletskiy/slorry/kpch/features/patterns"

type SymbolMatch struct {
	Sym    symbols.Symbol
	Weight float64
}

func Sharpify(img image.Image) *image.Gray {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	gray := image.NewGray(bounds)

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			if x == 0 || y == 0 || y == dy-1 || x == dx-1 {
				gray.Set(x, y, color.Gray{255})
				continue
			}

			weight := 0
			for i := -1; i <= 1; i += 1 {
				for j := -1; j <= 1; j += 1 {
					col := img.At(x+j, y+i)
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

func FindFeatures(img image.Image, x int, y int) features.FeatureSet {
	fs := []*features.Feature{
		&patterns.PipeLeft,
		&patterns.PipeRight,
		&patterns.CornerLeftBottom,
		&patterns.CornerRightBottom,
		&patterns.CornerRightTop,
		&patterns.CaveLeftSmall,
		&patterns.ArcTopLeft,
		&patterns.ArcTopRight,
		&patterns.LongHole,
		&patterns.DescentLeft,
		&patterns.CaveLeftBig,
		&patterns.AscentLeft,
		&patterns.BridgeLeft,
		&patterns.DescentLeftLong,
		&patterns.Snail,
		&patterns.MediumHole,
		&patterns.Pad,
		&patterns.Cornice,
		&patterns.NotchTop,
		&patterns.NotchBottom,
	}

	result := make(features.FeatureSet, 0)

	for _, feature := range fs {
		match := features.MatchFeature(feature, img, x, y)
		if match != nil {
			result = append(result, match)
		}
	}

	return result
}

func FindAllFeatures(img image.Image) features.FeatureSet {
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()
	matches := make(features.FeatureSet, 0)
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			matches = append(matches, FindFeatures(img, x, y)...)
		}
	}

	return matches
}

func PrintImage(img image.Image, features features.FeatureSet) {
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

	}
}

func Chunkify(img image.Image, fs features.FeatureSet,
	count int, overlap int, shift int) []features.FeatureSet {
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

	chunks := make([]features.FeatureSet, 0)
	size := (end - start) / count

	for chunk := 0; chunk < count; chunk++ {
		windowStart := start + chunk*size - overlap + chunk*shift
		windowEnd := start + (chunk+1)*size + overlap + chunk*shift

		window := make(features.FeatureSet, 0)

		for _, f := range fs {
			if f.X >= windowStart && f.X <= windowEnd {
				window = append(window, f)
			}
		}

		chunks = append(chunks, window)
	}

	return chunks
}

func FindSymbols(chunks []features.FeatureSet) []SymbolMatch {
	ss := []symbols.Symbol{
		maps.Zero,
		maps.One,
		maps.Two,
		maps.Three,
		maps.Four,
		maps.Five,
		maps.Six,
		maps.Seven,
		maps.Eight,
		maps.Nine,
	}

	result := make([]SymbolMatch, 0)
	usedFeatures := make(features.FeatureSet, 0)

	for _, chunk := range chunks {
		var matchedSymbol symbols.Symbol
		var symbolFeatures features.FeatureSet

		maxWeight := 0.0

		for _, symbol := range ss {
			weight, matchedFeatures := symbol.Match(chunk, usedFeatures)
			if weight > maxWeight {
				maxWeight = weight
				matchedSymbol = symbol
				symbolFeatures = matchedFeatures
			}

		}

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

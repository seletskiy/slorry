package symbols

import "log"

import "github.com/seletskiy/slorry/kpch/features"
import "github.com/seletskiy/slorry/kpch/features/align"

var _ = log.Print

type Symbol struct {
	Char     string
	AlignMap align.AlignMap
}

func (s Symbol) Match(fs features.FeatureSet, used features.FeatureSet) (float64, features.FeatureSet) {
	permutations := make(chan features.FeatureSet)
	go func() {
		Permutations(fs, s.AlignMap.Features, make(features.FeatureSet, 0),
			permutations)
		permutations <- nil
	}()

	var matched features.FeatureSet

	maxWeight := 0.0
	for x := range permutations {
		if len(x) == 0 {
			break
		}

		alignWeight := 0.0
		featuresWeight := float64(len(x)) / float64(len(s.AlignMap.Features))

		for _, link := range s.AlignMap.Aligns {
			if link.A >= len(x) {
				continue
			}

			if link.B >= len(x) {
				continue
			}

			if link.Align.Match(x[link.A], x[link.B]) {
				alignWeight += 1.0
			}

			for _, f := range used {
				if f == x[link.A] || f == x[link.B] {
					alignWeight -= 0.4
				}
			}
		}

		featuresWeight /= float64(len(s.AlignMap.Aligns))

		weight := alignWeight * featuresWeight

		if weight > maxWeight {
			maxWeight = weight
			matched = x
		}
	}

	return maxWeight, matched
}

func Permutations(fs features.FeatureSet, anchors []*features.Feature,
	result features.FeatureSet, out chan features.FeatureSet) {
	if len(anchors) == 0 || len(fs) == 0 {
		packet := make(features.FeatureSet, len(result))

		copy(packet, result)

		out <- packet

		return
	}

	for j, anchor := range anchors {
		//log.Println(j, anchor.Id)
		found := false
		for i, f := range fs {
			if f.Origin.Id == anchor.Id {
				rest := make(features.FeatureSet, i)
				copy(rest, fs[:i])
				Permutations(
					append(rest, fs[i+1:]...), anchors[j+1:],
					append(result, f),
					out)
				found = true
			}
		}

		if found {
			return
		}
	}

	Permutations(nil, nil, result, out)
}

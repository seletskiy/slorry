package symbols

import "log"

import . "slorry/kpch/features"
import . "slorry/kpch/features/align"

var _ = log.Print

type Symbol struct {
    Char string
    AlignMap AlignMap
}

func (s Symbol) Match(features FeatureSet, used FeatureSet) (float64, FeatureSet) {
    permutations := make(chan FeatureSet)
    go func() {
        Permutations(features, s.AlignMap.Features, make(FeatureSet, 0),
            permutations)
        permutations <- nil
    }()

    var matched FeatureSet

    maxWeight := 0.0
    for x := range permutations {
        if len(x) == 0 {
            break
        }

        alignWeight := 0.0
        featuresWeight := float64(len(x))/ float64(len(s.AlignMap.Features))

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

func Permutations(features FeatureSet, anchors []*Feature,
        result FeatureSet, out chan FeatureSet) {
    if len(anchors) == 0 || len(features) == 0 {
        packet := make(FeatureSet, len(result))

        copy(packet, result)

        out <- packet

        return
    }

    for j, anchor := range anchors {
        //log.Println(j, anchor.Id)
        found := false
        for i, f := range features {
            if f.Origin.Id == anchor.Id {
                rest := make(FeatureSet, i)
                copy(rest, features[:i])
                Permutations(
                    append(rest, features[i+1:]...), anchors[j+1:],
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

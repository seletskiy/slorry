package symbols

import "fmt"

import . "slorry/kpch/features"
import . "slorry/kpch/features/align"

var _ = fmt.Print

type Symbol struct {
    Char string
    AlignMap AlignMap
}

func (s Symbol) Match(features FeatureSet, used FeatureSet) (float64, FeatureSet) {
    permutations := make(chan FeatureSet)
    go func() {
        Perms(features, s.AlignMap.Features, make(FeatureSet, 0),
            permutations)
        permutations <- nil
    }()

    var matched FeatureSet

    max := 0.0
    //fmt.Printf("S %s\n", s.Char)
    for x := range permutations {
        if len(x) == 0 {
            break
        }

        weight := 0.0

        for _, link := range s.AlignMap.Aligns {
            ok := link.Align.Match(x[link.A], x[link.B])
            if ok {
                weight += 1.0
            }

            for _, f := range used {
                if f == x[link.A] || f == x[link.B] {
                    weight -= 0.4
                }
            }
        }

        if weight >= max {
            max = weight
            matched = x
        }
    }

    return max / float64(len(s.AlignMap.Aligns)), matched
}

func Perms(features FeatureSet, anchors []*Feature,
        result FeatureSet, out chan FeatureSet) bool {
    if len(anchors) == 0 {
        packet := make(FeatureSet, len(result))
        
        copy(packet, result)

        out <- packet

        return true
    }

    if len(features) == 0 {
        return false
    }

    anchor := anchors[0]
    for i, f := range features {
        if f.Origin.Id == anchor.Id {
            rest := features[i:]
            proceed := Perms(
                append(rest, features[:i]...), anchors[1:],
                append(result, features[i]),
                out)

            if !proceed {
                return false
            }
        }
    }

    return true
}

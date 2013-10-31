package align

import . "slorry/kpch/features"

type Align struct {
    Match func(*ConcreteFeature, *ConcreteFeature) bool
}

var LeftOf = Align {
    Match: func (i *ConcreteFeature, j *ConcreteFeature) bool {
        if i.X < j.X {
            return true
        } else {
            return false
        }
    },
}

var RightOf = Align {
    Match: func (i *ConcreteFeature, j *ConcreteFeature) bool {
        return LeftOf.Match(j, i)
    },
}

var Above = Align {
    Match: func (i *ConcreteFeature, j *ConcreteFeature) bool {
        if i.Y < j.Y {
            return true
        } else {
            return false
        }
    },
}

var Under = Align {
    Match: func (i *ConcreteFeature, j *ConcreteFeature) bool {
        return Above.Match(j, i)
    },
}

var SurelyAbove = Align {
    Match: func (i *ConcreteFeature, j *ConcreteFeature) bool {
        if j.Y - i.Y > 5 {
            return true
        } else {
            return false
        }
    },
}

var SurelyUnder = Align {
    Match: func (i *ConcreteFeature, j *ConcreteFeature) bool {
        return SurelyAbove.Match(j, i)
    },
}

var StraightLine = Align {
    Match: func (i *ConcreteFeature, j *ConcreteFeature) bool {
        if i.Y == j.Y || i.X == j.X {
            return true
        } else {
            return false
        }
    },
}

type AlignLink struct {
    A int
    B int
    Align *Align
}

type AlignFeatures []*Feature

type AlignMap struct {
    Features AlignFeatures
    Aligns []AlignLink
}

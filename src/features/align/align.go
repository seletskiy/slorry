package align

import . "features"

type Aligner interface {
    Match(*ConcreteFeature, *ConcreteFeature) bool
}

type Align struct {

}

type LeftOf Align

func (a LeftOf) Match(i *ConcreteFeature, j *ConcreteFeature) bool {
    if i.X < j.X {
        return true
    } else {
        return false
    }
}


type RightOf Align

func (a RightOf) Match(i *ConcreteFeature, j *ConcreteFeature) bool {
    if i.X > j.X {
        return true
    } else {
        return false
    }
}

type Above Align

func (a Above) Match(i *ConcreteFeature, j *ConcreteFeature) bool {
    if i.Y < j.Y {
        return true
    } else {
        return false
    }
}

type Under Align

func (a Under) Match(i *ConcreteFeature, j *ConcreteFeature) bool {
    if i.Y > j.Y {
        return true
    } else {
        return false
    }
}

type SurelyAbove Align

func (a SurelyAbove) Match(i *ConcreteFeature, j *ConcreteFeature) bool {
    if j.Y - i.Y > 5 {
        return true
    } else {
        return false
    }
}

type SurelyUnder Align

func (a SurelyUnder) Match(i *ConcreteFeature, j *ConcreteFeature) bool {
    if i.Y - j.Y > 5 {
        return true
    } else {
        return false
    }
}

type StraightLine Align

func (a StraightLine) Match(i *ConcreteFeature, j *ConcreteFeature) bool {
    if i.Y == j.Y || i.X == j.X {
        return true
    } else {
        return false
    }
}

type AlignLink struct {
    A int
    B int
    Align Aligner
}

type AlignFeatures []*Feature

type AlignMap struct {
    Features AlignFeatures
    Aligns []AlignLink
}

func (am AlignMap) Match(i *ConcreteFeature, j *ConcreteFeature) bool {
    return true
}


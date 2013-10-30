package maps

import . "features"
import . "features/align"
import . "features/masks"
import . "symbols"

var Zero = Symbol {
    Char: "0",
    AlignMap: AlignMap {
        []*Feature {
            &ArcTopLeft,
            &ArcTopRight,
            &LongHole,
        },
        []AlignLink {
            {0, 1, &LeftOf},
            {0, 1, &StraightLine},
            {2, 0, &RightOf},
            {2, 1, &LeftOf},
            {2, 0, &Under},
            {2, 1, &Under},
        },
    },
}

var One = Symbol {
    Char: "1",
    AlignMap: AlignMap {
        []*Feature {
            &PipeLeft,
            &PipeRight,
            &CornerRightTop,
            &Pad,
        },
        []AlignLink {
            {0, 1, &LeftOf},
            {0, 1, &StraightLine},
            {2, 0, &Above},
            {2, 1, &Above},
            {3, 0, &Under},
            {3, 1, &Under},
            {3, 2, &Under},
        },
    },
}

var Two = Symbol {
    Char: "2",
    AlignMap: AlignMap {
        []*Feature {
            &ArcTopLeft,
            &ArcTopRight,
            &Snail,
            &AscentLeft,
        },
        []AlignLink {
            {0, 1, &LeftOf},
            {0, 2, &LeftOf},
            {1, 2, &RightOf},
            {2, 0, &Under},
            {3, 2, &SurelyUnder},
            {3, 2, &LeftOf},
        },
    },
}


var Three = Symbol {
    Char: "3",
    AlignMap: AlignMap {
        []*Feature {
            &CornerRightTop,
            &ArcTopRight,
            &CaveLeftBig,
        },
        []AlignLink {
            {0, 1, &Above},
            {1, 2, &Above},
            {2, 1, &LeftOf},
            {2, 0, &LeftOf},
        },
    },
}

var Four = Symbol {
    Char: "4",
    AlignMap: AlignMap {
        []*Feature {
            &CornerRightTop,
            &DescentLeft,
            &BridgeLeft,
        },
        []AlignLink {
            {0, 1, &RightOf},
            {0, 2, &LeftOf},
            {0, 2, &SurelyAbove},
            {1, 2, &SurelyAbove},
        },
    },
}

var Five = Symbol {
    Char: "5",
    AlignMap: AlignMap {
        []*Feature {
            &CornerRightTop,
            &CornerRightBottom,
            &CaveLeftSmall,
        },
        []AlignLink {
            {0, 1, &Above},
            {0, 1, &StraightLine},
            {2, 0, &LeftOf},
            {2, 1, &LeftOf},
            {2, 0, &Under},
            {2, 1, &Under},
        },
    },
}

var Six = Symbol {
    Char: "6",
    AlignMap: AlignMap {
        []*Feature {
            &ArcTopLeft,
            &ArcTopRight,
            &CornerRightTop,
            &MediumHole,
        },
        []AlignLink {
            {0, 1, &LeftOf},
            {0, 1, &StraightLine},
            {2, 0, &Above},
            {2, 1, &Above},
            {2, 0, &RightOf},
            {2, 1, &LeftOf},
            {3, 0, &RightOf},
            {1, 3, &RightOf},
            {2, 3, &SurelyAbove},
        },
    },
}

var Seven = Symbol {
    Char: "7",
    AlignMap: AlignMap {
        []*Feature {
            &DescentLeftLong,
            &DescentLeftLong,
            &Pad,
        },
        []AlignLink {
            {0, 1, &LeftOf},
            {0, 1, &StraightLine},
            {2, 0, &Above},
            {2, 1, &Above},
        },
    },
}

var Eight = Symbol {
    Char: "8",
    AlignMap: AlignMap {
        []*Feature {
            &ArcTopLeft,
            &ArcTopRight,
            &ArcTopLeft,
            &ArcTopRight,
        },
        []AlignLink {
            {0, 1, &LeftOf},
            {0, 1, &StraightLine},
            {2, 3, &LeftOf},
            {2, 3, &StraightLine},
            {0, 2, &SurelyAbove},
            {1, 3, &SurelyAbove},
        },
    },
}

var Nine = Symbol {
    Char: "9",
    AlignMap: AlignMap {
        []*Feature {
            &ArcTopLeft,
            &ArcTopRight,
            &MediumHole,
            &AscentLeft,
        },
        []AlignLink {
            {0, 1, &LeftOf},
            {0, 1, &StraightLine},
            {0, 2, &LeftOf},
            {1, 2, &RightOf},
            {3, 2, &Under},
        },
    },
}

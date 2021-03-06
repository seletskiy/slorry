package patterns

import . "slorry/kpch/features"

var PipeLeft = Feature {
    Id: `PipeLeft`,
    Variants: []FeatureMask {
        {
            ` |`,
            ` |`,
            ` |`,
            ` |`,
            `+A`,
            `| `,
            `| `,
            `| `,
        }, {
            `  |`,
            `  |`,
            `  |`,
            `  |`,
            `+-A`,
            `|  `,
            `|  `,
            `|  `,
        },
    },
}

var PipeRight = Feature {
    Id: `PipeRight`,
    Variants: []FeatureMask {
        {
            `| `,
            `| `,
            `| `,
            `| `,
            `B+`,
            ` |`,
            ` |`,
            ` |`,
        }, {

            `|  `,
            `|  `,
            `|  `,
            `|  `,
            `B-+`,
            `  |`,
            `  |`,
            `  |`,
        },
    },
}

var CornerLeftBottom = Feature {
    Id: `CornerLeftBottom`,
    Variants: []FeatureMask {
        {
            `|    `,
            `|    `,
            `|    `,
            `D----`,
        },
    },
}

var CornerRightBottom = Feature {
    Id: `CornerRightBottom`,
    Variants: []FeatureMask {
        {
            `    |`,
            `    |`,
            `    |`,
            `----E`,
        },
    },
}

var CornerRightTop = Feature {
    Id: `CornerRightTop`,
    Variants: []FeatureMask {
        {
            `----F`,
            `    |`,
            `    |`,
        },
    },
}

var CaveLeftSmall = Feature {
    Id: `CaveLeftSmall`,
    Variants: []FeatureMask {
        {
            `---. `,
            `    G`,
            ` --. `,
        },
        {
            `----G`,
            ` .   `,
            `-----`,
        },
    },
}


var ArcTopRight = Feature {
    Id: `ArcTopRight`,
    Variants: []FeatureMask {
        {
            ` .   `,
            `  K  `,
            `   . `,
            `   | `,
            `   | `,
            `   | `,
        },
        {
            `.   `,
            ` .  `,
            ` K  `,
            `  . `,
            `  | `,
            `  | `,
        },
    },
}

var ArcTopLeft = Feature {
    Id: `ArcTopLeft`,
    Variants: []FeatureMask {
        {
            `   .`,
            `  L `,
            ` .  `,
            ` |  `,
            ` |  `,
            ` |  `,
        },
        {
            `   .`,
            `  . `,
            `  L `,
            ` .  `,
            ` |  `,
            ` |  `,
            ` |  `,
        },
        {
            `    ..`,
            `  L.  `,
            ` .    `,
            ` |    `,
            ` |    `,
            ` |    `,
        },
    },
}

var LongHole = Feature {
    Id: `LongHole`,
    Variants: []FeatureMask {
        {
            ` .M `,
            `/  \`,
            `|  |`,
            `|  |`,
            `|  |`,
            `|  |`,
            `|  |`,
        },
    },
}

var DescentLeft = Feature {
    Id: `DescentLeft`,
    Variants: []FeatureMask {
        {
            `   ----`,
            `  N    `,
            ` .     `,
            `.      `,
            `.      `,
        },
        {
            `   ----`,
            `  N    `,
            ` .     `,
            ` .     `,
            `.      `,
        },
    },
}

var CaveLeftBig = Feature {
    Id: `CaveLeftBig`,
    Variants: []FeatureMask {
        {
            `--. `,
            `   O`,
            `   |`,
            `   |`,
            `   |`,
            `--. `,
        },
        {
            `---. `,
            `   | `,
            `    O`,
            `   | `,
            `---. `,
        },
    },
}

var AscentLeft = Feature {
    Id: `AscentLeft`,
    Variants: []FeatureMask {
        {
            `  .  `,
            ` .   `,
            ` .   `,
            `.    `,
            `P----`,
        },
        {
            ` .   `,
            ` .   `,
            `.    `,
            `.    `,
            `P----`,
        },
        {
            `  .  `,
            `  .  `,
            ` .   `,
            `.    `,
            `P----`,
        },
    },
}

var BridgeLeft = Feature {
    Id: `BridgeLeft`,
    Variants: []FeatureMask {
        {
            `|  `,
            `|  `,
            `|  `,
            `++ `,
            ` | `,
            ` Q `,
            ` | `,
            `++ `,
            `|  `,
            `|  `,
            `|  `,
        },
        {
            `|  `,
            `|  `,
            `|  `,
            `++ `,
            ` Q `,
            ` | `,
            `++ `,
            `|  `,
            `|  `,
            `|  `,
        },
    },
}

var DescentLeftLong = Feature {
    Id: `DescentLeftLong`,
    Variants: []FeatureMask {
        {
            `   .`,
            `   .`,
            `  . `,
            `  R `,
            `  . `,
            ` .  `,
            ` .  `,
            ` .  `,
        },
        {
            `   .`,
            `    `,
            `  R `,
            `  . `,
            ` .  `,
            ` .  `,
            ` .  `,
            `    `,
            `.   `,
        },
    },
}

var Snail = Feature {
    Id: `Snail`,
    Variants: []FeatureMask {
        {
            `  ..`,
            `  ..`,
            `..S `,
            `  . `,
        },
        {
            ` ...`,
            ` . .`,
            `.S .`,
            `  . `,
            `  . `,
        },
    },
}

var MediumHole = Feature {
    Id: `MediumHole`,
    Variants: []FeatureMask {
        {
            ` .T.`,
            `.  .`,
            ` ...`,
        },
        {
            `  T `,
            ` . .`,
            `.  .`,
            ` ...`,
        },
        {
            `.T.`,
            `. .`,
            `.  `,
            `...`,
        },
        {
            `.T.`,
            `  .`,
            `  .`,
            `...`,
        },
        {
            `.T.`,
            `. .`,
            `...`,
        },
        {
            `.T.`,
            `.  `,
            `...`,
        },
        {
            `.T.`,
            `  .`,
            `...`,
        },
    },
}

var Pad = Feature {
    Id: `Pad`,
    Variants: []FeatureMask {
        {
            `----U----`,
        },
    },
}

var Cornice = Feature {
    Id: `Cornice`,
    Variants: []FeatureMask {
        {
            `.    `,
            `...V `,
            `  .  `,
            `  .  `,
            `  ...`,
        },
        {
            `.     `,
            ` ...V `,
            `   .  `,
            `   .  `,
            `  .   `,
            `  ....`,
        },
        {
            `.    `,
            `..V  `,
            `  .  `,
            ` .   `,
            ` .   `,
            ` ....`,
        },
        {
            `.     `,
            ` ..V  `,
            `   .  `,
            `  .   `,
            `  .   `,
            `  ....`,
        },
        {
            `.    `,
            `..V  `,
            `  .  `,
            `..   `,
            `.    `,
            `j....`,
        },
        {
            `.    `,
            `...V `,
            `     `,
            `  .  `,
            `  ...`,
        },
    },
}

var NotchTop = Feature {
    Id: `NotchTop`,
    Variants: []FeatureMask {
        {
            `  .  `,
            ` .   `,
            ` .   `,
            `W..  `,
            `   ..`,
            `    .`,
            `     `,
        },
        {
            `  .  `,
            `  .  `,
            ` .   `,
            ` W.  `,
            `   ..`,
            `    .`,
            `     `,
        },
        {
            `   . `,
            `  .  `,
            ` .   `,
            ` W.  `,
            `   ..`,
            `    .`,
            `     `,
        },
    },
}

var NotchBottom = Feature {
    Id: `NotchBottom`,
    Variants: []FeatureMask {
        {
            `..Y`,
            `  .`,
            `  .`,
            ` . `,
            `.  `,
        },
        {
            `.. `,
            `  Y`,
            ` . `,
            ` . `,
            `.  `,
        },
        {
            `.   `,
            ` ...`,
            `  Y `,
            ` .. `,
            ` .  `,
            ` .  `,
        },
        {
            `.   `,
            ` ...`,
            `  Y `,
            `  . `,
            ` .  `,
            ` .  `,
        },
    },
}

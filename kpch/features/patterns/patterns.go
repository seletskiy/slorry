package patterns

import "github.com/seletskiy/slorry/kpch/features"

var PipeLeft = features.Feature{
	Id: `PipeLeft`,
	Variants: []features.FeatureMask{
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

var PipeRight = features.Feature{
	Id: `PipeRight`,
	Variants: []features.FeatureMask{
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

var CornerLeftBottom = features.Feature{
	Id: `CornerLeftBottom`,
	Variants: []features.FeatureMask{
		{
			`|    `,
			`|    `,
			`|    `,
			`D----`,
		},
	},
}

var CornerRightBottom = features.Feature{
	Id: `CornerRightBottom`,
	Variants: []features.FeatureMask{
		{
			`    |`,
			`    |`,
			`    |`,
			`----E`,
		},
	},
}

var CornerRightTop = features.Feature{
	Id: `CornerRightTop`,
	Variants: []features.FeatureMask{
		{
			`----F`,
			`    |`,
			`    |`,
		},
	},
}

var CaveLeftSmall = features.Feature{
	Id: `CaveLeftSmall`,
	Variants: []features.FeatureMask{
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

var ArcTopRight = features.Feature{
	Id: `ArcTopRight`,
	Variants: []features.FeatureMask{
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

var ArcTopLeft = features.Feature{
	Id: `ArcTopLeft`,
	Variants: []features.FeatureMask{
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

var LongHole = features.Feature{
	Id: `LongHole`,
	Variants: []features.FeatureMask{
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

var DescentLeft = features.Feature{
	Id: `DescentLeft`,
	Variants: []features.FeatureMask{
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

var CaveLeftBig = features.Feature{
	Id: `CaveLeftBig`,
	Variants: []features.FeatureMask{
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

var AscentLeft = features.Feature{
	Id: `AscentLeft`,
	Variants: []features.FeatureMask{
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

var BridgeLeft = features.Feature{
	Id: `BridgeLeft`,
	Variants: []features.FeatureMask{
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

var DescentLeftLong = features.Feature{
	Id: `DescentLeftLong`,
	Variants: []features.FeatureMask{
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

var Snail = features.Feature{
	Id: `Snail`,
	Variants: []features.FeatureMask{
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

var MediumHole = features.Feature{
	Id: `MediumHole`,
	Variants: []features.FeatureMask{
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

var Pad = features.Feature{
	Id: `Pad`,
	Variants: []features.FeatureMask{
		{
			`----U----`,
		},
	},
}

var Cornice = features.Feature{
	Id: `Cornice`,
	Variants: []features.FeatureMask{
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

var NotchTop = features.Feature{
	Id: `NotchTop`,
	Variants: []features.FeatureMask{
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

var NotchBottom = features.Feature{
	Id: `NotchBottom`,
	Variants: []features.FeatureMask{
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

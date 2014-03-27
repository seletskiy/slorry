package maps

import "github.com/seletskiy/slorry/kpch/features"
import "github.com/seletskiy/slorry/kpch/features/align"
import "github.com/seletskiy/slorry/kpch/features/patterns"
import "github.com/seletskiy/slorry/kpch/symbols"

var Zero = symbols.Symbol{
	Char: "0",
	AlignMap: align.AlignMap{
		[]*features.Feature{
			&patterns.ArcTopLeft,
			&patterns.ArcTopRight,
			&patterns.LongHole,
		},
		[]align.AlignLink{
			{0, 1, &align.LeftOf},
			{0, 1, &align.StraightLine},
			{2, 0, &align.RightOf},
			{2, 1, &align.LeftOf},
			{2, 0, &align.Under},
			{2, 1, &align.Under},
		},
	},
}

var One = symbols.Symbol{
	Char: "1",
	AlignMap: align.AlignMap{
		[]*features.Feature{
			&patterns.PipeLeft,
			&patterns.PipeRight,
			&patterns.CornerRightTop,
			&patterns.Pad,
		},
		[]align.AlignLink{
			{0, 1, &align.LeftOf},
			{0, 1, &align.StraightLine},
			{2, 0, &align.Above},
			{2, 1, &align.Above},
			{3, 0, &align.Under},
			{3, 1, &align.Under},
			{3, 2, &align.Under},
		},
	},
}

var Two = symbols.Symbol{
	Char: "2",
	AlignMap: align.AlignMap{
		[]*features.Feature{
			&patterns.ArcTopLeft,
			&patterns.ArcTopRight,
			&patterns.Snail,
			&patterns.AscentLeft,
		},
		[]align.AlignLink{
			{0, 1, &align.LeftOf},
			{0, 2, &align.LeftOf},
			{1, 2, &align.RightOf},
			{2, 0, &align.Under},
			{3, 2, &align.SurelyUnder},
			{3, 2, &align.LeftOf},
		},
	},
}

var Three = symbols.Symbol{
	Char: "3",
	AlignMap: align.AlignMap{
		[]*features.Feature{
			&patterns.Pad,
			&patterns.ArcTopRight,
			&patterns.CaveLeftBig,
			&patterns.Cornice,
		},
		[]align.AlignLink{
			{0, 1, &align.Above},
			{1, 2, &align.Above},
			{2, 1, &align.LeftOf},
			{3, 0, &align.LeftOf},
			{3, 2, &align.Above},
		},
	},
}

var Four = symbols.Symbol{
	Char: "4",
	AlignMap: align.AlignMap{
		[]*features.Feature{
			&patterns.CornerRightTop,
			&patterns.DescentLeft,
			&patterns.BridgeLeft,
			&patterns.CornerRightTop,
		},
		[]align.AlignLink{
			{0, 1, &align.RightOf},
			{0, 2, &align.LeftOf},
			{0, 2, &align.SurelyAbove},
			{1, 2, &align.SurelyAbove},
			{3, 0, &align.SurelyUnder},
			{3, 1, &align.SurelyUnder},
			{3, 2, &align.Under},
			{3, 0, &align.LeftOf},
			{3, 2, &align.LeftOf},
		},
	},
}

var Five = symbols.Symbol{
	Char: "5",
	AlignMap: align.AlignMap{
		[]*features.Feature{
			&patterns.CornerRightTop,
			&patterns.CornerRightBottom,
			&patterns.CaveLeftSmall,
		},
		[]align.AlignLink{
			{0, 1, &align.Above},
			{0, 1, &align.StraightLine},
			{2, 0, &align.LeftOf},
			{2, 1, &align.LeftOf},
			{2, 0, &align.Under},
			{2, 1, &align.Under},
		},
	},
}

var Six = symbols.Symbol{
	Char: "6",
	AlignMap: align.AlignMap{
		[]*features.Feature{
			&patterns.ArcTopLeft,
			&patterns.ArcTopRight,
			&patterns.CornerRightTop,
			&patterns.MediumHole,
			&patterns.NotchTop,
		},
		[]align.AlignLink{
			{0, 1, &align.LeftOf},
			{0, 1, &align.StraightLine},
			{2, 0, &align.Above},
			{2, 1, &align.Above},
			{2, 0, &align.RightOf},
			{2, 1, &align.LeftOf},
			{3, 0, &align.RightOf},
			{1, 3, &align.RightOf},
			{2, 3, &align.SurelyAbove},
			{4, 0, &align.RightOf},
			{4, 1, &align.LeftOf},
			{4, 2, &align.Under},
		},
	},
}

var Seven = symbols.Symbol{
	Char: "7",
	AlignMap: align.AlignMap{
		[]*features.Feature{
			&patterns.DescentLeftLong,
			&patterns.DescentLeftLong,
			&patterns.Pad,
		},
		[]align.AlignLink{
			{0, 1, &align.SurelyLeftOf},
			{0, 1, &align.StraightLine},
			{2, 0, &align.Above},
			{2, 1, &align.Above},
		},
	},
}

var Eight = symbols.Symbol{
	Char: "8",
	AlignMap: align.AlignMap{
		[]*features.Feature{
			&patterns.ArcTopLeft,
			&patterns.ArcTopRight,
			&patterns.ArcTopLeft,
			&patterns.ArcTopRight,
		},
		[]align.AlignLink{
			{0, 1, &align.LeftOf},
			{0, 1, &align.StraightLine},
			{2, 3, &align.LeftOf},
			{2, 3, &align.StraightLine},
			{0, 2, &align.SurelyAbove},
			{1, 3, &align.SurelyAbove},
		},
	},
}

var Nine = symbols.Symbol{
	Char: "9",
	AlignMap: align.AlignMap{
		[]*features.Feature{
			&patterns.ArcTopLeft,
			&patterns.ArcTopRight,
			&patterns.MediumHole,
			&patterns.AscentLeft,
			&patterns.NotchBottom,
		},
		[]align.AlignLink{
			{0, 1, &align.LeftOf},
			{0, 2, &align.LeftOf},
			{1, 2, &align.RightOf},
			{3, 2, &align.LeftOf},
			{4, 0, &align.SurelyUnder},
			{4, 1, &align.SurelyUnder},
			{4, 1, &align.LeftOf},
			{4, 2, &align.SurelyUnder},
			{4, 3, &align.Above},
			{4, 3, &align.RightOf},
		},
	},
}

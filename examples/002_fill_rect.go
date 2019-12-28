package main

import (
	. "github.com/buchanae/ink/color"
	. "github.com/buchanae/ink/dd"
	. "github.com/buchanae/ink/gfx"
)

func Ink(doc *Doc) {

	Fill{
		Mesh: Rect{
			XY{0.2, 0.2},
			XY{0.8, 0.8},
		},
		Color: Blue,
	}.Draw(doc)
}

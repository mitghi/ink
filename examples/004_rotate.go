package main

import (
	. "github.com/buchanae/ink/color"
	. "github.com/buchanae/ink/dd"
	. "github.com/buchanae/ink/gfx"
)

func Ink(z *Layer) {

	r := Rect{
		XY{0.2, 0.2},
		XY{0.8, 0.8},
	}

	s := z.Shader(r)
	s.Set("a_pivot", r.Center())
	s.Set("a_rot", 0.4)
	s.Set("a_color", Red)
}
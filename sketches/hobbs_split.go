package main

import (
	"github.com/buchanae/ink/app"
	"github.com/buchanae/ink/color"
	. "github.com/buchanae/ink/dd"
	. "github.com/buchanae/ink/gfx"
	"github.com/buchanae/ink/rand"
)

func Ink(doc *app.Doc) {
	rand.SeedNow()
	Clear(doc, color.White)

	a := Triangle{
		B: XY{1, 0},
		C: XY{0, 1},
	}

	b := Triangle{
		A: XY{1, 1},
		B: XY{1, 0},
		C: XY{0, 1},
	}

	tris := recursive(8, a, b)
	p := rand.Palette()

	layer := doc.NewLayer()

	var strokes []Stroke
	for _, t := range tris {
		Fill{t, rand.Color(p)}.Draw(layer)
		strokes = append(strokes, t.Stroke())
	}

	for _, stk := range strokes {
		stk.Width = 0.0005
		Fill{stk, color.Black}.Draw(layer)
	}

	cir := Circle{
		XY:       XY{0.5, 0.5},
		Radius:   0.4,
		Segments: 100,
	}
	Cut{
		Shape:  cir,
		Source: layer,
	}.Draw(doc)

	cs := cir.Stroke()
	cs.Width = 0.005
	Fill{cs, color.Black}.Draw(doc)

}

func recursive(depth int, tris ...Triangle) []Triangle {
	if depth == 0 {
		return nil
	}
	var out []Triangle
	for _, t := range tris {
		a, b := split(t)
		out = append(out, a, b)
		out = append(out, recursive(depth-1, a, b)...)
	}
	return out
}

func split(t Triangle) (Triangle, Triangle) {

	edges := t.Edges()
	lens := [3]float32{
		edges[0].SquaredLength(),
		edges[1].SquaredLength(),
		edges[2].SquaredLength(),
	}

	do := func(long, a, b Line) (Triangle, Triangle) {
		mid := long.Interpolate(rand.Range(0.3, 0.7))
		return Triangle{
				mid, a.A, a.B,
			}, Triangle{
				mid, b.A, b.B,
			}
	}

	switch {
	case lens[0] >= lens[1] && lens[0] >= lens[2]:
		return do(edges[0], edges[1], edges[2])
	case lens[1] >= lens[0] && lens[1] >= lens[2]:
		return do(edges[1], edges[0], edges[2])
	default:
		return do(edges[2], edges[0], edges[1])
	}
}

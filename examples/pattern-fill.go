package main

import "github.com/vecno-io/go-magi"

func main() {
	im, err := magi.LoadPNG("examples/baboon.png")
	if err != nil {
		panic(err)
	}
	pattern := magi.NewSurfacePattern(im, magi.RepeatBoth)
	dc := magi.NewContext(600, 600)
	dc.MoveTo(20, 20)
	dc.LineTo(590, 20)
	dc.LineTo(590, 590)
	dc.LineTo(20, 590)
	dc.ClosePath()
	dc.SetFillStyle(pattern)
	dc.Fill()
	dc.SavePNG("out.png")
}

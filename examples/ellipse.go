package main

import "github.com/vecno-io/go-magi"

func main() {
	const S = 1024
	dc := magi.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(magi.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	if im, err := magi.LoadImage("examples/gopher.png"); err == nil {
		dc.DrawImageAnchored(im, S/2, S/2, 0.5, 0.5)
	}
	dc.SavePNG("out.png")
}

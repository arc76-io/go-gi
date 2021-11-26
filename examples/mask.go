package main

import (
	"log"

	"github.com/vecno-io/go-magi"
)

func main() {
	im, err := magi.LoadImage("examples/baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := magi.NewContext(512, 512)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	dc.Clip()
	dc.DrawImage(im, 0, 0)
	dc.SavePNG("out.png")
}

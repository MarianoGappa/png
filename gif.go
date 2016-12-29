package main

import (
	"bufio"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math/rand"
	"os"
	"time"
)

func drawGIF() {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	p := []color.Color{
		color.RGBA{244, 67, 54, 255},
		color.RGBA{233, 30, 99, 255},
		color.RGBA{156, 39, 176, 255},
	}

	img1 := image.NewPaletted(image.Rect(0, 0, 99, 99), p)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img1.SetColorIndex(x, y, uint8(rnd.Intn(3)))
		}
	}
	img2 := image.NewPaletted(image.Rect(0, 0, 99, 99), p)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img2.SetColorIndex(x, y, uint8(rnd.Intn(3)))
		}
	}
	img3 := image.NewPaletted(image.Rect(0, 0, 99, 99), p)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img3.SetColorIndex(x, y, uint8(rnd.Intn(3)))
		}
	}

	g := gif.GIF{
		Image: []*image.Paletted{img1, img2, img3},
		Delay: []int{10, 10, 10},
	}

	f, _ := os.Create("out.gif")
	defer f.Close()
	w := bufio.NewWriter(f)
	err := gif.EncodeAll(w, &g)
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()

}

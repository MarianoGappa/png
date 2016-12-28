package main

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func drawPNG(s string) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	img := image.NewRGBA(image.Rect(0, 0, 999, 999))

	var r, g, b, a uint32
	var rd, gd, bd, ad uint32
	var rl, gl, bl, al uint32
	var rt, gt, bt, at uint32
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if x > 0 && y > 0 && x < 1000 && y < 1000 {
				rd, gd, bd, ad = img.At(x-1, y-1).RGBA()
				rl, gl, bl, al = img.At(x-1, y).RGBA()
				rt, gt, bt, at = img.At(x, y-1).RGBA()
				if ad > 0 {
					rd = uint32(float64(rd) / float64(ad) * float64(256))
					gd = uint32(float64(gd) / float64(ad) * float64(256))
					bd = uint32(float64(bd) / float64(ad) * float64(256))
				}
				if al > 0 {
					rl = uint32(float64(rl) / float64(al) * float64(256))
					gl = uint32(float64(gl) / float64(al) * float64(256))
					bl = uint32(float64(bl) / float64(al) * float64(256))
				}
				if at > 0 {
					rt = uint32(float64(rt) / float64(at) * float64(256))
					gt = uint32(float64(gt) / float64(at) * float64(256))
					bt = uint32(float64(bt) / float64(at) * float64(256))
				}

				switch rnd.Intn(3) {
				case 0:
					r = rd
				case 1:
					r = rl
				case 2:
					r = rt
				}
				if r > 0 && r < 255 {
					r = r + uint32(rnd.Intn(3)-1)
				}

				switch rnd.Intn(3) {
				case 0:
					g = gd
				case 1:
					g = gl
				case 2:
					g = gt
				}
				if g > 0 && g < 255 {
					g = g + uint32(rnd.Intn(3)-1)
				}

				switch rnd.Intn(3) {
				case 0:
					b = bd
				case 1:
					b = bl
				case 2:
					b = bt
				}
				if b > 0 && b < 255 {
					b = b + uint32(rnd.Intn(3)-1)
				}

				switch rnd.Intn(3) {
				case 0:
					a = ad
				case 1:
					a = al
				case 2:
					a = at
				}
				if a > 0 && a < 255 {
					a = a + uint32(rnd.Intn(3)-1)
				}
			}

			if r == 0 {
				r = uint32(rnd.Intn(256))
			}
			if g == 0 {
				g = uint32(rnd.Intn(256))
			}
			if b == 0 {
				b = uint32(rnd.Intn(256))
			}
			if a == 0 {
				a = uint32(rnd.Intn(256))
			}
			img.SetRGBA(
				x,
				y,
				color.RGBA{
					uint8(r),
					uint8(g),
					uint8(b),
					uint8(255),
				})

		}
	}
	f, _ := os.Create("out.png")
	defer f.Close()
	w := bufio.NewWriter(f)
	png.Encode(w, img)
	w.Flush()
}

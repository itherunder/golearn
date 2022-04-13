package mandelbrot_

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"testing"
)

func TestMandelbrot(t *testing.T) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 4096, 4096
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	// png.Encode(os.Stdout, img) // NOTE: ignoring errors
	file, _ := os.Create("mandelbrot.png")
	defer file.Close()
	// png.Encode(file, img)
	png.Encode(file, superSampling(img))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{123 - contrast*n, 255 - contrast*n, 211 - contrast*n, 255}
		}
	}
	return color.RGBA{12, 235, 211, 255}
}

func superSampling(img image.Image) image.Image {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	newImage := image.NewRGBA(img.Bounds())
	for x := 0; x < width-1; x++ {
		for y := 0; y < height-1; y++ {
			var col color.RGBA
			r00, g00, b00, a00 := img.At(x, y).RGBA()
			r01, g01, b01, a01 := img.At(x, y+1).RGBA()
			r10, g10, b10, a10 := img.At(x+1, y).RGBA()
			r11, g11, b11, a11 := img.At(x+1, y+1).RGBA()
			col.R = uint8((uint(uint8(r00)) + uint(uint8(r01)) + uint(uint8(r10)) + uint(uint8(r11))) / 4)
			col.G = uint8((uint(uint8(g00)) + uint(uint8(g01)) + uint(uint8(g10)) + uint(uint8(g11))) / 4)
			col.B = uint8((uint(uint8(b00)) + uint(uint8(b01)) + uint(uint8(b10)) + uint(uint8(b11))) / 4)
			col.A = uint8((uint(uint8(a00)) + uint(uint8(a01)) + uint(uint8(a10)) + uint(uint8(a11))) / 4)
			newImage.Set(x, y, col)
		}
	}
	return newImage
}

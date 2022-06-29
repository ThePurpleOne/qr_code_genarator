package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// ! COLORS
type color_t int
const (
	BLACK color_t = iota
	WHITE
	GRAY
	RED
	GREEN
	BLUE
)



type pixels struct {
	w, h int
	data []color_t
}


func create_pixel_array(w, h int)  pixels{
	pixel_array := make([]color_t, w*h)
	for i := 0; i < w*h; i++ {
		pixel_array[i] = GRAY
	}
	return pixels{w, h, pixel_array}
}

func (p* pixels) set_pixel(x, y int, value color_t) {
	p.data[y * p.h + x] = value
}

func (p* pixels) get_pixel(x, y int) color_t {
	return p.data[y * p.h + x]
}

func (p pixels)to_img(scaler int) image.Image {
	surface := image.Rect(0, 0, p.w * scaler, p.h * scaler)
	img := image.NewRGBA(surface)

	for y := 0 ; y < p.h * scaler; y++ {
		for x := 0; x < p.w * scaler; x++ {
			
			switch p.data[p.w * (y/scaler) + x/scaler] {
				case BLACK:
					img.Set(x, y, color.Black)
				case WHITE:
					img.Set(x, y, color.White)
				case GRAY:
					img.Set(x, y, color.Gray{128})
				case RED:
					img.Set(x, y, color.RGBA{255, 0, 0, 255})
				case GREEN:
					img.Set(x, y, color.RGBA{0, 255, 0, 255})
				case BLUE:
					img.Set(x, y, color.RGBA{0, 0, 255, 255})
			}
		}
	}
	return img
}

func (p pixels)save_to_png(filename string, scaler int) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	err = png.Encode(f, p.to_img(scaler))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Image created")
}



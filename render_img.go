package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)


type pixels struct {
	w, h int
	data []bool
}

func create_pixel_array(w, h int)  pixels{
	pixel_array := make([]bool, w*h)
	for i := 0; i < w*h; i++ {
		pixel_array[i] = false
	}
	return pixels{w, h, pixel_array}
}

func (p* pixels) set_pixel(x, y int, value bool) {
	p.data[y * p.h + x] = value
}

func (p* pixels) get_pixel(x, y int) bool {
	return p.data[y * p.h + x]
}

func (p pixels)to_img(scaler int) image.Image {
	surface := image.Rect(0, 0, p.w * scaler, p.h * scaler)
	img := image.NewRGBA(surface)

	for y := 0 ; y < p.h * scaler; y++ {
		for x := 0; x < p.w * scaler; x++ {
			if p.data[p.w * (y/scaler) + x/scaler] {
				img.Set(x, y, color.White)
			}else{
				img.Set(x, y, color.Black)
			}
		}
	}
	return img
}

func (p pixels)save_to_png(scaler int, filename string) {
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



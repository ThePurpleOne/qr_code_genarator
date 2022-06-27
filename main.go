package main

func main() {

	pixel_scaler := 10
	w, h := 21, 21

	pix := create_pixel_array(w, h)

	for i := 0; i < w*h; i++ {
		if i%2 == 0 {
			pix.data[i] = true
		} else {
			pix.data[i] = false
		}
	}

	pix.save_to_png("test.png", pixel_scaler)

	a := create_poly([]int64{4, 56, 69, 2, 4}, 5)
	a.show()

}
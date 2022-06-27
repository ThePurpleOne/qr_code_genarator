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

	// ! TESTING POLYNOMIAL
	a := create_poly([]int64{3, 3, 1, 0, 2}, 5)
	b := create_poly([]int64{1, 1, 3, -4, 2}, 5)
	a.show()
	b.show()
	add_ab := a.add(b)
	add_ab.show()

	mul_ab := a.mul(b)
	mul_ab.show()
}
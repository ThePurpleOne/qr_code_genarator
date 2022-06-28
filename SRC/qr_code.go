package main

type qr_code struct {
	pix     pixels
	version int
	margin  int
}

func create_qr_code(version int, margin int) qr_code {

	nb_modules := (version-1)*4 + 21
	return qr_code{create_pixel_array(nb_modules+margin+1, nb_modules+margin+1), version, margin}
}

/**
*	Brief : Add Quiete Zone to the QR Code
**/
func (q *qr_code) add_margin() {

	for y := 0; y < q.pix.h; y++ {
		for x := 0; x < q.pix.w; x++ {
			if x < q.margin || x >= q.pix.w-q.margin || y < q.margin || y >= q.pix.h-q.margin {
				q.pix.set_pixel(x, y, true)
			}
		}
	}
}

func (q *qr_code) add_finders() {

	q.add_finder(q.margin+3, q.margin+3)         // TOP LEFT FINDER
	q.add_finder(q.pix.w-q.margin-4, q.margin+3) // TOP RIGHT FINDER
	q.add_finder(q.margin+3, q.pix.h-q.margin-4) // BOTTOM LEFT FINDER
}

func (q *qr_code) add_finder(center_x, center_y int) {

	y := center_y - 2
	for x := center_x - 2; x < center_x+3; x++ { // ! X AXIS
		q.pix.set_pixel(x, y, true)
		q.pix.set_pixel(x, y+4, true)
	}
	x := center_x - 2
	for y := center_y - 2; y < center_y+3; y++ { // ! Y AXIS
		q.pix.set_pixel(x, y, true)
		q.pix.set_pixel(x+4, y, true)
	}

	// ADDIN SEPARATOR AROUND THE FINDER,
	// GETS INSIDE THE MARGIN BUT ITS WHITE ANYWAY

	y = center_y - 4
	for x := center_x - 3; x < center_x+5; x++ { // ! X AXIS
		q.pix.set_pixel(x, y, true)
		q.pix.set_pixel(x, y+8, true)
	}
	x = center_x - 4
	for y := center_y - 3; y < center_y+5; y++ { // ! Y AXIS
		q.pix.set_pixel(x, y, true)
		q.pix.set_pixel(x+8, y, true)
	}
}

func (q *qr_code) add_timing_patterns() {

	y := q.margin + 6
	for x := q.margin + 8; x < q.pix.w-q.margin-8; x++ {
		if x%2 == 0 {
			q.pix.set_pixel(x, y, true)
		}
	}
	x := q.margin + 6
	for y := q.margin + 8; y < q.pix.w-q.margin-8; y++ {
		if y%2 == 0 {
			q.pix.set_pixel(x, y, true)
		}
	}
}

func (q *qr_code) add_alignement_patterns() {
	// 6, 28, 50
	// HARD CODED FOR NOW
	a, b, c := 6, 26, 46

	q.add_alignement_pattern(q.margin+a, q.margin+b)
	q.add_alignement_pattern(q.margin+b, q.margin+a)

	//q.add_alignement_pattern(q.margin+a, q.margin+c) // ONTO FINDER
	//q.add_alignement_pattern(q.margin+c, q.margin+a) // ONTO FINDER

	q.add_alignement_pattern(q.margin+b, q.margin+c)
	q.add_alignement_pattern(q.margin+c, q.margin+b)

	//q.add_alignement_pattern(q.margin+a, q.margin+a) // ONTO FINDER
	q.add_alignement_pattern(q.margin+b, q.margin+b)
	q.add_alignement_pattern(q.margin+c, q.margin+c)
}

func (q *qr_code) add_alignement_pattern(center_x, center_y int) {

	for x := center_x - 2; x < center_x+3; x++ {
		for y := center_y - 2; y < center_y+3; y++ {
			if x < center_x-1 || x > center_x+1 || y < center_y-1 || y > center_y+1 {
				q.pix.set_pixel(x, y, false)
			} else {
				q.pix.set_pixel(x, y, true)
			}
		}
	}

	q.pix.set_pixel(center_x, center_y, false)
}

func (q *qr_code) save_to_png(filename string, scale int) {
	q.pix.save_to_png(filename, scale)
}
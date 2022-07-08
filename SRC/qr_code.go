package main

type correction_level int

// https://www.thonky.com/qr-code-tutorial/format-version-information
const (
	MEDIUM   correction_level = iota // 15 %
	LOW                              // 7 %
	HIGH                             // 30 %
	QUARTILE                         // 25 %
)

type qr_code struct {
	pix            pixels
	version        int
	margin         int
	formating_mode int
	correct        correction_level
}

const MODE_INDICATOR = 0b0010 // MODE BYTE
const ERROR_CORRECTION = LOW

func create_qr_code(version int, margin int) qr_code {
	nb_modules := (version-1)*4 + 21
	return qr_code{create_pixel_array(nb_modules+margin+1, nb_modules+margin+1), version, margin, 0, MEDIUM}
}

func (q *qr_code) modify_correction_level(correct correction_level) {
	q.correct = correct
}

func (q *qr_code) modify_format(new_format int) {
	q.formating_mode = new_format
}

func (q *qr_code) add_margin() {

	for y := 0; y < q.pix.h; y++ {
		for x := 0; x < q.pix.w; x++ {
			if x < q.margin || x >= q.pix.w-q.margin || y < q.margin || y >= q.pix.h-q.margin {
				q.pix.set_pixel(x, y, WHITE)
			}
		}
	}
}

func (q *qr_code) add_timing_patterns() {

	y := q.margin + 6
	for x := q.margin + 8; x < q.pix.w-q.margin-8; x++ {
		if x%2 == 0 {
			q.pix.set_pixel(x, y, WHITE)
			q.pix.set_pixel(y, x, WHITE)
		} else {
			q.pix.set_pixel(x, y, BLACK)
			q.pix.set_pixel(y, x, BLACK)
		}
	}
}

// ! FINDERS
func (q *qr_code) add_finders() {

	q.add_finder(q.margin+3, q.margin+3)         // TOP LEFT FINDER
	q.add_finder(q.pix.w-q.margin-4, q.margin+3) // TOP RIGHT FINDER
	q.add_finder(q.margin+3, q.pix.h-q.margin-4) // BOTTOM LEFT FINDER
}

// TODO : GO WITH LESS LOOPS
func (q *qr_code) add_finder(center_x, center_y int) {

	// ! WHITE CONTOUR
	y := center_y - 4
	for x := center_x - 4; x < center_x+5; x++ {
		q.pix.set_pixel(x, y, WHITE)
		q.pix.set_pixel(x, y+8, WHITE)
		q.pix.set_pixel(y, x, WHITE)
		q.pix.set_pixel(y+8, x, WHITE)
	}

	// ! INSIDE BLACK CONTOUR
	y = center_y - 3
	for x := center_x - 3; x < center_x+4; x++ {
		q.pix.set_pixel(x, y, BLACK)
		q.pix.set_pixel(x, y+6, BLACK)
		q.pix.set_pixel(y, x, BLACK)
		q.pix.set_pixel(y+6, x, BLACK)
	}

	// ! INSIDE WHITE CONTOUR AROUND SQUARE
	y = center_y - 2
	for x := center_x - 2; x < center_x+3; x++ {
		q.pix.set_pixel(x, y, WHITE)
		q.pix.set_pixel(x, y+4, WHITE)
		q.pix.set_pixel(y, x, WHITE)
		q.pix.set_pixel(y+4, x, WHITE)
	}

	// ! INSIDE BLACK SQUARE
	y = center_y - 1
	for x := center_x - 1; x < center_x+2; x++ {
		q.pix.set_pixel(x, y, BLACK)
		q.pix.set_pixel(x, y+1, BLACK)
		q.pix.set_pixel(x, y+2, BLACK)
	}
}

// ! ALIGNEMENTS
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
				q.pix.set_pixel(x, y, BLACK)
			} else {
				q.pix.set_pixel(x, y, WHITE)
			}
		}
	}

	q.pix.set_pixel(center_x, center_y, BLACK)
}

func (q *qr_code) add_dark_module() {
	q.pix.set_pixel(q.margin+8, (4*q.version)+9, BLACK)
}

func (q *qr_code) save_to_png(filename string, scale int) {
	q.pix.save_to_png(filename, scale)
}

// ! ----------------------------------------------
// ! -------------------- DATA --------------------
// ! ----------------------------------------------

const MAX_DATA = 102

type data_block struct {
	data             []byte
	error_correction []byte
}

type data_group struct {
	block1 data_block
	block2 data_block
}

// TAKES DATA STRING AND PUTS IT INTO STRING OF '0's and '1's WITH ERROR CORRECTION
// IT HAS TO FILL BOTH THE BLOCKS, IF THE DATA ISNT LONG ENOUGH, IT WILL BE PADDED
func create_data_block(data string) (data_block, data_block) {
	// SIZES AND ALL HARD CODED FOR NOW
	// MAX DATA :
	// 50 bytes / block pour for group 1
	// 51 bytes / block pour for group 2

	// PAD WITH HARD CODED 11101100 00010001 if needed
	for len(data) < MAX_DATA {
		if len(data)%2 == 0 {
			data += string(0b11101100)
		} else {
			data += string(0b00010001)
		}
	}

	// SPLIT INTO TWO BLOCKS
	block1 := data[:50]
	block2 := data[50:]

	// ENCODE DATA
	block1_encoded := []byte(block1)
	block2_encoded := []byte(block2)

	

}

func create_data_group(block1, block2 data_block) data_group {
	return data_group{block1, block2}
}

// ! -------------------- UTILS --------------------

func encode_text(text string) []byte {
	return []byte(text)
}

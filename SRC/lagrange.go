package main

func get_lagrange_poly(points []float64) polynomial {

	// CREATE POLYNOMIALS
	//poly_l :=
	poly := create_poly(make([]float64, len(points)))

	for x_i := 0; x_i < len(points); x_i++ {
		y_i := points[x_i]
		poly_n := create_poly([]float64{1})

		for x_k := 0; x_k < len(points); x_k++ {

			if x_i != x_k {
				d := (float64(x_i) - float64(x_k))
				poly_n = poly_n.mul(create_poly([]float64{-float64(x_k) * 1 / d, 1 / d}))
			}
		}
		poly_n = poly_n.mul(create_poly([]float64{float64(y_i)}))
		poly = poly.add(poly_n)
	}

	return poly
}

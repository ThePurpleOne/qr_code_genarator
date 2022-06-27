package main

import "math"

func get_lagrange_poly(points []float64, prime float64) polynomial {

	// CREATE POLYNOMIALS
	poly_l := make([]float64, len(points))
	poly := create_poly(poly_l, prime)

	for x_i := 0; x_i < len(points); x_i++ {
		y_i := points[x_i]
		poly_n := create_poly([]float64{1}, prime)

		for x_k := 0; x_k < len(points); x_k++ {
			y_k := points[x_k]
			if x_i != x_k {
				d := inverse_mul(math.Mod((y_i-y_k), prime), prime)
				poly_n = poly_n.mul(create_poly([]float64{math.Mod(-y_k*d, prime), math.Mod(d, prime)}, prime))
			}
		}
		poly = poly.add(poly_n)
	}
	return poly
}

//poly_l = [0] * len(l)
//poly = polynome(poly_l)
//for x_i, y_i in enumerate(l):
//	poly_n = polynome([1])
//	for x_k, y_k in enumerate(l):
//		if y_k[0] != y_i[0]:
//			d = inverse_mult((y_i[0]-y_k[0]) % poly_n.prime_mod, poly_n.prime_mod)
//			poly_n = poly_n.mul(polynome([-y_k[0] * d % poly_n.prime_mod, d % poly_n.prime_mod]))
//	poly_n = poly_n.mul(polynome([y_i[1] % poly_n.prime_mod]))
//	poly = poly.add(poly_n)
//return poly

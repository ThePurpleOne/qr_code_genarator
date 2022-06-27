package main

import (
	"fmt"
)

func get_lagrange_poly(points []float64) polynomial {

	// CREATE POLYNOMIALS
	//poly_l := 
	poly := create_poly(make([]float64, len(points)))

	for x_i := 0; x_i < len(points); x_i++ {
		y_i := points[x_i]
		poly_n := create_poly([]float64{1})

		for x_k := 0; x_k < len(points); x_k++ {
			fmt.Printf("x_i : %d\n", x_i)
			fmt.Printf("x_k : %d\n", x_k)

			if x_i != x_k {
				test := (float64(x_i)-float64(x_k))
				fmt.Printf("x_i - x_k : %f\n", test)
				d := (float64(x_i) - float64(x_k))
				//poly_n = poly_n.mul(create_poly([]float64{math.Mod(-float64(x_k)*d, prime), math.Mod(d, prime)}, prime))
				poly_n = poly_n.mul(create_poly([]float64{-float64(x_k) * 1 / d, 1 / d }))
			}
		}
		poly_n = poly_n.mul(create_poly([]float64{float64(y_i)}))
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

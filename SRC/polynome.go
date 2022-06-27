package main

import (
	"fmt"
	"math"
)

type polynomial struct {
	coefs []float64
	mod float64
}

func create_poly(coefs []float64, mod float64) polynomial {
	return polynomial{coefs, mod}
}

func (p* polynomial)show() {
	for i := len(p.coefs) - 1; i >= 0; i-- {
		power := ""
		coef := ""
		
		if p.coefs[i] != 0 { 	// AVOID SHOWING 0
			coef = fmt.Sprintf("%.02f", p.coefs[i])			
			if i != 0 { 		// NOT 0's degree
				power = fmt.Sprintf("x^%d + ", i)
			}
		}else{
			coef = ""
		}
		fmt.Printf("%s%s", coef, power)
	}
	fmt.Println()
}

func (p polynomial)add(p2 polynomial) polynomial{
	if p.mod != p2.mod {
		panic("Cannot add polynomials with different mod")
	}

	// PADDING WITH 0 IF LENGTHS ARE DIFFERENT
	if len(p.coefs) > len(p2.coefs) {
		for len(p.coefs) > len(p2.coefs) {
			p2.coefs = append(p2.coefs, 0)
		}
	}else{
		for len(p2.coefs) > len(p.coefs) {
			p.coefs = append(p.coefs, 0)
		}
	}

	out_coefs := make([]float64, len(p.coefs))
	for i := 0; i < len(p.coefs); i++ {
		out_coefs[i] = p.coefs[i] + p2.coefs[i]
	}
	return create_poly(out_coefs, p.mod)
}

func (p polynomial)mul(p2 polynomial) polynomial{
	out_coefs := make([]float64, len(p.coefs) + len(p2.coefs) - 1)
	for i := 0; i < len(p.coefs); i++ {
		for j := 0; j < len(p2.coefs); j++ {
			out_coefs[i+j] += math.Mod(p.coefs[i] * p2.coefs[j], p.mod)
		}
	}
	return create_poly(out_coefs, p.mod)
}

//def mul(self, poly_2):
//	coeff_poly_res = [0] * (len(self.coefs) + len(poly_2.coefs) - 1)
//	for index_1, value_1 in enumerate(self.coefs):
//		for index_2, value_2 in enumerate(poly_2.coefs):
//			coeff_poly_res[index_1 + index_2] += (value_1 * value_2) % self.prime_mod
//	return polynome(coeff_poly_res)




func (p polynomial)eval(x float64) float64{
	out := float64(0)
	// USING HORNER's METHOD
	for i := len(p.coefs) - 1; i >= 0; i-- {
		out = (out * x + p.coefs[i])
	}
	return math.Mod(out, p.mod)
}
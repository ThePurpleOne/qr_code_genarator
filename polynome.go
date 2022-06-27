package main

import (
	"fmt"
)

type polynomial struct {
	coefs []int64
	mod int64
}

func create_poly(coefs []int64, mod int64) polynomial {
	return polynomial{coefs, mod}
}

func (p* polynomial)show() {
	for i := len(p.coefs) - 1; i >= 0; i-- {
		power := ""
		coef := ""
		
		if p.coefs[i] != 0 { 	// AVOID SHOWING 0
			coef = fmt.Sprintf("%d", p.coefs[i])			
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

	out_coefs := make([]int64, len(p.coefs))
	for i := 0; i < len(p.coefs); i++ {
		out_coefs[i] = p.coefs[i] + p2.coefs[i]
	}
	return create_poly(out_coefs, p.mod)
}
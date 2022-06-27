package main

import (
	"fmt"
)

type polynomial struct {
	coefs []int64
	mod uint64
}

func create_poly(coefs []int64, mod uint64) polynomial {
	return polynomial{coefs, mod}
}

func (p* polynomial)show() {
	for i := len(p.coefs) - 1; i >= 0; i-- {

		//SHOW i
		//fmt.Printf("%d ", p.coefs[i])
		power := ""
		coef := p.coefs[i]
		if p.coefs[i] != 0 { 	// AVOID SHOWING 0
			if i != 0 { 		// NOT 0's degree
				power = fmt.Sprintf("x^%d + ", i)
			}
		}else{
			coef = ' '
		}
		fmt.Printf("%d%s", coef, power)
	}
	fmt.Println()
}
package main

import (
	"fmt"
	"math"
)

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func extended_euclide(a int64, b int64) (int64, int64, int64) {
	a, b = abs(a), abs(b)

	// INIT
	old_r, r := a, b
	old_s, s := int64(1), int64(0)

	for r != 0 {
		q := old_r / r
		old_r, r = r, (old_r - q*r)
		old_s, s = s, (old_s - q*s)
	}

	bez_t := int64(0)
	if b != 0 {
		bez_t = (old_r - old_s * a) / b
	}

	// GCD, x, y
	return old_r, old_s, bez_t
}

func euclide_check(a, b, x, y, pgcd int64) bool {
	if pgcd == (a*x + b*y) {
		return true
	}
	return false
}

func inverse_mul(a, p float64) float64 {
	
	if a < 0{
		a += p
	}

	gcd, x, y := extended_euclide(int64(a), int64(p))
	fmt.Printf("A : %f\n", a)
	fmt.Printf("P : %f\n", p)
	fmt.Printf("GCD : %d\n", gcd)

	if gcd != 1 {
		panic("gcd != 1")
	}

	fmt.Printf("a : %f\n", a)
	fmt.Printf("x : %d\n", x)

	xf := float64(x)
	
	if euclide_check(int64(a), int64(p), x, y, gcd){
		println("TRUE")
	}else{
		println("FALSE")
	}


	c := math.Mod(a * xf, p)

	fmt.Printf("c = %f\n", c)

	return float64(x)
}
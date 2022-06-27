package main

import (
	"fmt"
	"math"
)

func main() {

	// ! TESTING PIXEL RENDERING
	fmt.Println("\n\n\n ----- TESTING PIXELS RENDERING -----")

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
	fmt.Println("\n\n\n ----- TESTING POLINOMIALS -----")

	//a := create_poly([]int64{3, 3, 1, 0, 2}, 5)
	a := create_poly([]float64{1, 3, 2})
	b := create_poly([]float64{3, -4, 2})
	
	fmt.Printf("a : ")
	a.show()
	fmt.Printf("b : ")
	b.show()
	
	add_ab := a.add(b)
	fmt.Printf("a + b : ")
	add_ab.show()
	
	mul_ab := a.mul(b)
	fmt.Printf("a * b : ")
	mul_ab.show()
	
	x_ := float64(7)
	ax := a.eval(x_)
	fmt.Printf("a(%f) = %f\n", x_, ax)

	// ! TESTING EXTENDED EUCLIDE 
	fmt.Println("\n\n\n ----- TESTING EUCLIDE -----")
	nb_a := float64(499115)
	nb_b := float64(1197)
	gcd, x, y := extended_euclide(int64(nb_a), int64(nb_b))

	fmt.Printf("gcd(%f, %f) = %d\n", nb_a, nb_b, gcd)

	// SHOW MULTIPLILCATIVE INVERSES
	fmt.Printf("nb_a = %f\n", nb_a)
	fmt.Printf("nb_b = %f\n", nb_b)
	fmt.Printf("%f^-1 mod %f = %d\n", nb_a, nb_b, x)
	fmt.Printf("%f^-1 mod %f = %d\n", nb_b, nb_a, y)
	
	fmt.Printf("Mult Inverse of %f mod %f : %f\n", nb_a, nb_b, inverse_mul(nb_a, nb_b))
	fmt.Printf("Mult Inverse of %f mod %f : %f\n", nb_b, nb_a, inverse_mul(nb_b, nb_a))

	// ! TESTING LAGRANGE  
	fmt.Println("\n\n\n ----- TESTING LAGRANGE -----")

	test1 := 200.1
	test2 := 100.05

	fmt.Printf("FLOAT MOD : %f %% %f = %f\n", test1, test2, math.Mod(test1, test2) )

	list_of_points := []float64{4, 3, 6}
	lgr := get_lagrange_poly(list_of_points)
	lgr.show()
}
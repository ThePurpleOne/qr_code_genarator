package main

import "fmt"

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
	a := create_poly([]int64{88, 99897, 787, 58, 4, 5, 7, 9, 4, 1, 3, 2}, 5)
	b := create_poly([]int64{1, 1, 3, -4, 2}, 5)
	
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
	
	x := int64(7)
	ax := a.eval(x)
	//bx := b.eval(x)
	fmt.Printf("a(%d) = %d\n", x, ax)

	// ! TESTING EXTENDED EUCLIDE 
	fmt.Println("\n\n\n ----- TESTING EUCLIDE -----")
	nb_a := int64(4991)
	nb_b := int64(1197)
	gcd, x, y := extended_euclide(nb_a, nb_b)
	fmt.Printf("gcd(%d, %d) = %d\n", nb_a, nb_b, gcd)
	
	// SHOW MULTIPLILCATIVE INVERSES
	fmt.Printf("nb_a = %d\n", nb_a)
	fmt.Printf("nb_b = %d\n", nb_b)
	fmt.Printf("%d^-1 mod %d = %d\n", nb_a, nb_b, x)
	fmt.Printf("%d^-1 mod %d = %d\n", nb_b, nb_a, y)

}
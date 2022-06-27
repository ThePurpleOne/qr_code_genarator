package main

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func extended_euclide(a int64, b int64) (int64, int64, int64) {
	a, b = abs(a), abs(b)

	if b > a {
		a, b = b, a
	}

	if b == 0 {
		panic("b == 0")
	}

	// INIT
	old_r, r := a, b
	old_s, s := int64(1), int64(0)
	old_t, t := int64(0), int64(1)

	for r != 0 {
		q := old_r / r
		old_r, r = r, (old_r - q*r)
		old_s, s = s, (old_s - q*s)
		old_t, t = t, (old_t - q*t)
	}

	// GCD, x, y
	return old_r, old_s, old_t
}

//func extended_euclide(a int64, b int64) (int64, int64, int64) {
//	a, b = abs(a), abs(b)

//	if b > a {
//		a, b = b, a
//	}

//	if b == 0 {
//		panic("b == 0")
//	}

//	// INITALIZING ARRAYS FOR r, q, x, y
//	r := make([]int64, 100)
//	q := make([]int64, 100)
//	x := make([]int64, 100)
//	y := make([]int64, 100)
//	step := 2

//	// INIT
//	r[0] = a
//	r[1] = b

//	x[0] = 1
//	x[1] = 0

//	y[0] = 0
//	y[1] = 1

//	for r[1] != 0 {
//		r[step] = a % b
//		q[step] = a / b
//		x[step] = x[step-2] - q[step]*x[step-1]
//		y[step] = y[step-2] - q[step]*y[step-1]

//		a, b = b, r[step]
//		step++
//	}

//	// PGCD, x, y
//	return r[step-2], x[step-2], y[step-2]
//}

func euclide_check(a, b, x, y, pgcd int64) bool {
	if pgcd == (a*x + b*y) {
		return true
	}
	return false
}
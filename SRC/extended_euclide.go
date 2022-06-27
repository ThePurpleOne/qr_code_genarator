package main

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
		bez_t = (old_r - old_s*a) / b
	}

	// GCD, x, y
	return old_r, old_s, bez_t
}

func euclide_check(a, b, x, y, pgcd int64) bool {
	return pgcd == (a*x + b*y)
}

func inverse_mul(a, p float64) float64 {

	if a < 0 {
		a += p
	}

	gcd, x, _ := extended_euclide(int64(a), int64(p))

	if gcd != 1 {
		panic("gcd != 1")
	}
	return float64(x)
}
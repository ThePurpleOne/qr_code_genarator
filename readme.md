# QR CODE GENERATOR

## Polynomial manpulation

### Create Polynomial
> Second parameter is a modulo
```go
b := create_poly([]int64{1, 1, 3, -4, 2}, 5)
```

### Show Polynomial
```go
b.show()
```
OUTPUTS 
```go
2x^4 + -4x^3 + 3x^2 + 1x^1 + 1
```

### Add Polynomial
```go
add_ab := a.add(b)
```

### Multiply Polynomial
```go
mul_ab := a.mul(b)
```

### Evaluate Polynomial
```go
x := int64(7)
ax := a.eval(x)
```

--- 

## PNG Rendering API
### Create Pixels object
```go
pix := create_pixel_array(w, h)
```

### Set a pixel
```go
pix.set_pixel(x, y, true)
```

### Get a pixel
```go
pix.get_pixel(x, y)
```

### Encode pixels into PNG
```go
img := pix.to_img()
```

### SAVE Pixels object to PNG
```go
pix.save_to_png(10, "image.png")
```

### Output
![](ASSETS/test.png)

---

## Multiplicative Inverse
### Extended Euclidean Algorithm
> x, y are Bézout coefficients
```go
gcd, x, y := extended_euclidean(a, b)
```

### Find Multiplicative Inverse
> Find a^-1 mod b
```go
x := mult_inverse(a, b)
```
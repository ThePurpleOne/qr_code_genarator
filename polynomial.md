# Polynomial manpulation

## Create Polynomial
> Second parameter is a modulo
```go
b := create_poly([]int64{1, 1, 3, -4, 2}, 5)
```

## Show Polynomial
```go
b.show()
```
OUTPUTS 
```go
2x^4 + -4x^3 + 3x^2 + 1x^1 + 1
```

## Add Polynomial
```go
add_ab := a.add(b)
```

## Multiply Polynomial
```go
mul_ab := a.mul(b)
```

## Evaluate Polynomial
```go
x := int64(7)
ax := a.eval(x)
```
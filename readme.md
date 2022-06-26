## PNG RENDERING API
### CREATE PIXELS OBJECT
```go
pix := create_pixel_array(w, h)
```

### SET A PIXEL
```go
pix.set_pixel(x, y, true)
```

### GET A PIXEL
```go
pix.get_pixel(x, y)
```

### ENCODE PIXELS INTO PNG
```go
img := pix.to_img()
```

### SAVE PIXELS OBJECT INTO A PNG FILE
```go
pix.save_to_png("image.png")
```
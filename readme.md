## PNG RENDERING API
### CREATE PIXELS OBJECT
To create a pixels object, containing the size of the array and the pixel data (BOOL -> true is White, false is Black, 'cause only writing black and white for now).

```go
pix := create_pixel_array(w, h)
```

### SET A PIXEL
To set a pixel, you need to specify the x and y coordinates, and the value of the pixel.

```go
set_pixel(pix, x, y, true)
```

### GET A PIXEL
To get a pixel, you need to specify the x and y coordinates.

```go
get_pixel(pix, x, y)
```

### ENCODE PIXELS INTO PNG
Encode the pixels object into a PNG.

```go
img := pix.to_img()
```

### SAVE PIXELS OBJECT INTO A PNG FILE
Save the pixels object into a PNG file.

```go
pix.save_to_png("image.png")
```
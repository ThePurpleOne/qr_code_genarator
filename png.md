# PNG Rendering API
## Create Pixels object
```go
pix := create_pixel_array(w, h)
```

## Set a pixel
```go
pix.set_pixel(x, y, true)
```

## Get a pixel
```go
pix.get_pixel(x, y)
```

## Encode pixels into PNG
```go
img := pix.to_img()
```

## SAVE Pixels object to PNG
```go
pix.save_to_png(10, "image.png")
```

## Output
![](ASSETS/test.png)
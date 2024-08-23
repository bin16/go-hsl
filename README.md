``` go
// from color.Color
var clr = hsl.FromColor(color.RGBA{255, 0, 0, 255})

// New()
var clr = hsl.New(0, 100, 50)

// struct
var clr = hsl.HSL{0, 1.0, .5}

clr.RGBA() 
// 65535, 0, 0, 65535

clr.ToRGBA() // -> color.RGBA
// color.RGBA{255, 0, 0, 255}
```

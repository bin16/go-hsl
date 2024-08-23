``` go
// from color.Color
var clr = hsl.FromColor(color.RGBA{255, 0, 0, 255})

// New()
var clr = hsl.New(0, 100, 50)

// struct
var clr = hsl.HSL{0, 1.0, .5}

// to color.RGBA{}
clr.ToRGBA()
// color.RGBA{0xff, 0, 0, 0xff}

// for color.Color
clr.RGBA() 
// 0xffff, 0, 0, 0xffff
```

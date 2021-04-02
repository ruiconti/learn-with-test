package structs

import (
    "math"
)

type Rectangle struct {
    Width  float64
    Height float64
}

// function 
func Perimeter(rec Rectangle) float64 {
    return 2 * (rec.Width + rec.Height) 
}

// methods
func (r Rectangle) Area() float64 {
    return r.Width * r.Height 
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
    Height float64
    Base   float64
}

func (t Triangle) Area() float64 {
    return (t.Height * t.Base) / 2
}

// interface
type Shape interface {
    Area() float64
}

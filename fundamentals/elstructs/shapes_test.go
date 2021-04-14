package structs

import (
    "testing"
)

func TestPerimeter(t *testing.T) {
    got := Perimeter(Rectangle{10.0, 10.0})
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }

    got = Perimeter(Rectangle{20.0, 25.0})
    want = 90.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {
    areaTests := []struct {
        shape Shape
        want  float64
    }{
        {shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
        {shape: Circle{Radius: 10}, want: 314.1592653589793},
        {shape: Triangle{Height: 12, Base: 6}, want: 36.0},
    }

    for _, st := range areaTests {
        got := st.shape.Area()
        if got != st.want {
            t.Errorf("%#v got %.2f want %.2f", got, st.want)
        }
    }
}


package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 40.0},
		{Circle{10.0}, 62.83185307179586},
		{Triangle{10.0, 10.0}, -0.1},
	}

	//t.Run("Testing rectangle perimeter", func(t *testing.T) {
	//	rectangle := Rectangle{10.0, 10.0}
	//	got := rectangle.Perimeter()
	//	want := 40.0
	//
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//})
	//
	//t.Run("Testing circle perimeter", func(t *testing.T) {
	//	circle := Circle{10}
	//	got := circle.Perimeter()
	//	want := 62.83185307179586
	//
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//})
	//
	//t.Run("Testing triangle perimeter", func(t *testing.T) {
	//	triangle := Triangle{10, 5}
	//	got := triangle.Perimeter()
	//	want := -0.1
	//
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//})

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

	//checkArea := func(t testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//}
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{12, 6}
	//	checkArea(t, rectangle, 72.0)
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10.0}
	//	checkArea(t, circle, 314.1592653589793)
	//})
}

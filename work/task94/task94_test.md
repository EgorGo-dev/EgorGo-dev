package main

import "testing"

func TestRectangleArea(t *testing.T) {
    tests := []struct {
        name string
        w, h float64
        want float64
    }{
        {"square", 4, 4, 16},
        {"rectangle", 3, 5, 15},
        {"zero width", 0, 5, 0},
        {"fraction", 1.5, 2, 3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := Rectangle{Width: tt.w, Height: tt.h}
            if got := r.Area(); got != tt.want {
                t.Errorf("Area() = %v; want %v", got, tt.want)
            }
        })
    }
}
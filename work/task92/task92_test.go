package main

import "testing"

func TestMultiply(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 6},
        {"zero", 0, 5, 0},
        {"negative", -2, 3, -6},
        {"both negative", -2, -3, 6},
        {"one", 1, 7, 7},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Multiply(tt.a, tt.b); got != tt.want {
                t.Errorf("Multiply(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
package main

import "testing"

func TestGenerate(t *testing.T) {
    ch := Generate(5)
    want := []int{1, 2, 3, 4, 5}
    i := 0
    for v := range ch {
        if i >= len(want) || v != want[i] {
            t.Fatalf("got %d at position %d; want %d", v, i, want[i])
        }
        i++
    }
    if i != len(want) {
        t.Errorf("got %d values; want %d", i, len(want))
    }
}
package main

import "testing"

func TestDivide(t *testing.T) {
    got, err := Divide(10, 2)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if got != 5 {
        t.Errorf("Divide(10,2) = %d; want 5", got)
    }

    _, err = Divide(10, 0)
    if err == nil {
        t.Error("expected error for division by zero, got nil")
    }
}
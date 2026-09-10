package main

import "testing"

func TestSum(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"empty", []int{}, 0},
        {"one", []int{5}, 5},
        {"many", []int{1, 2, 3, 4, 5}, 15},
        {"with negatives", []int{-1, 5, -2}, 2},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Sum(tt.nums); got != tt.want {
                t.Errorf("Sum(%v) = %d; want %d", tt.nums, got, tt.want)
            }
        })
    }
}
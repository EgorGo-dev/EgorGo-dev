package main

import "testing"

func TestSpeaker(t *testing.T) {
    tests := []struct {
        name string
        s    Speaker
        want string
    }{
        {"dog", Dog{}, "Гав"},
        {"cat", Cat{}, "Мяу"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.s.Speak(); got != tt.want {
                t.Errorf("Speak() = %q; want %q", got, tt.want)
            }
        })
    }
}
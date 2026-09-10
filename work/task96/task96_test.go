package main

import (
    "encoding/json"
    "testing"
)

func TestUserJSON(t *testing.T) {
    u := User{Name: "Егор", Age: 14}
    data, err := json.Marshal(u)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    want := `{"name":"Егор","age":14}`
    if string(data) != want {
        t.Errorf("got %s; want %s", string(data), want)
    }
}
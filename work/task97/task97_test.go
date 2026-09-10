package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()

    HelloHandler(rec, req)

    if rec.Code != http.StatusOK {
        t.Errorf("status = %d; want %d", rec.Code, http.StatusOK)
    }
    want := "Hello, Go Web!\n"
    if rec.Body.String() != want {
        t.Errorf("body = %q; want %q", rec.Body.String(), want)
    }
}
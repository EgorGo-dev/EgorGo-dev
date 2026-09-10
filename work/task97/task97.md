Дан HTTP-обработчик:

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Go Web!")
}

Напишите тест TestHelloHandler с httptest.NewRecorder.
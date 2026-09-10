Дана функция:

func Generate(n int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := 1; i <= n; i++ {
            ch <- i
        }
    }()
    return ch
}

Напишите тест TestGenerate, проверяющий числа от 1 до 5.
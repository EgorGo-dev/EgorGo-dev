Дана функция:

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

Напишите тест TestDivide, проверяющий успех и ошибку.
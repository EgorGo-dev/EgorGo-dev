Дана структура:

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

Напишите тест TestUserJSON, проверяющий сериализацию в JSON.
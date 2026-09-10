Дан интерфейс и две реализации:

type Speaker interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Гав" }

type Cat struct{}
func (c Cat) Speak() string { return "Мяу" }

Напишите табличный тест TestSpeaker.
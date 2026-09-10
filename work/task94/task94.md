Дана структура и метод:

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

Напишите табличный тест TestRectangleArea.
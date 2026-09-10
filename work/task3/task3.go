package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

func currentDayOfTheWeek() string {
	now := TimeNow()
	switch now.Weekday() {
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	case time.Sunday:
		return "Воскресенье"
	}
	return ""
}

func dayOrNight() string {
	now := TimeNow()
	hour := now.Hour()
	if hour >= 10 && hour <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	now := TimeNow()
	today := int(now.Weekday()) // число от 0 до 6
	if today == 5 {
		return 0
	}
	days := (5 - today + 7) % 7
	return days
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	currentDay := currentDayOfTheWeek()
	return strings.EqualFold(answer, currentDay)
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}
	boolean := strings.EqualFold(answer, dayOrNight())
	return boolean, nil
}

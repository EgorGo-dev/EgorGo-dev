# 🚀 EgorGo-dev — Мои решения задач по Go

[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://go.dev/)
[![SQLite](https://img.shields.io/badge/SQLite-3-blue?logo=sqlite)](https://sqlite.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-blue?logo=postgresql)](https://postgresql.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![GitHub](https://img.shields.io/badge/GitHub-181717?logo=github)](https://github.com/EgorGo-dev)
[![GitHub stars](https://img.shields.io/github/stars/EgorGo-dev/go-projects?style=social)](https://github.com/EgorGo-dev/go-projects)
[![GitHub last commit](https://img.shields.io/github/last-commit/EgorGo-dev/go-projects)](https://github.com/EgorGo-dev/go-projects)

<div align="center">
  <img src="https://img.shields.io/badge/Go_задач-200%2B-blue?style=for-the-badge&logo=go" />
  <img src="https://img.shields.io/badge/SQL_задач-100%2B-green?style=for-the-badge&logo=sqlite" />
  <img src="https://img.shields.io/badge/Сертификаты-6-orange?style=for-the-badge&logo=academia" />
</div>

<br>

Здесь я храню решения задач по Go из **Яндекс Лицея** и других источников. Репозиторий показывает мой прогресс в изучении языка, владение Git и GitHub.

---

### 👔 Для рекрутера / работодателя

<details>
<summary><b>Хотите быстро оценить мой уровень? Начните с этих 4 задач уровня Boss</b> (нажмите, чтобы развернуть)</summary>

1. **[REST API на Go](./work/HARDtask1/hardTask.go)** — HTTP-сервер с middleware, JSON и синхронизацией.
2. **[SQL подзапросы](./work/HARDtask2/hardTask2.sql)** — коррелированные подзапросы, `NOT EXISTS`, обработка дубликатов.
3. **[Оконные функции в PostgreSQL](./work/HARDtask3/hardTask3.sql)** — `DISTINCT ON`, `RANK`, сложные CTE.
4. **[Параллелизм в Go](./work/HARDtask4/hardTask4.go)** — горутины, каналы, `sync.WaitGroup`, worker pool.

**Ключевые навыки:** Go · горутины · каналы · синхронизация · HTTP · JSON · SQL · PostgreSQL · оконные функции · CTE · middleware.

</details>

---

## 📖 Содержание

| 👤 **Обо мне** | 🛠 **Стек** | 📊 **Прогресс** | 📝 **Задачи** |
|:---:|:---:|:---:|:---:|
| 🎯 **Цели** | 🏆 **Сертификаты** | 🙏 **Благодарности** | 📬 **Контакты** |

---

## 👨‍💻 Обо мне

- Мне **14 лет**, начал программировать в 13 с C++, через год перешёл на Go.
- 🎓 Учусь в [Яндекс Лицее](https://lyceum.yandex.ru/) (направление: Веб-разработка на Go).
- 🥇 **1 место** на курсах: Go-1, Go-2, Go-3, SQL.
- Решил **200+ задач** на Go, активно изучаю SQL (SQLite/PostgreSQL).
- 🚀 Интересуюсь бэкендом, высоконагруженными системами и облаками.
- Всегда открыт к обсуждению кода и новым идеям.

---

## 🛠 Стек

**Языки и технологии:**  
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![SQLite](https://img.shields.io/badge/SQLite-003B57?style=for-the-badge&logo=sqlite&logoColor=white) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)

**Инструменты:** Git · GitHub · VS Code  
**Пакеты Go (Стандартная библиотека):** `fmt` · `time` · `strings` · `strconv` · `log` · `errors` · `slices` · `io` · `os` · `sync` · `regexp` · `context` · `net/http` · `testing` · `math`  
**Пакеты Go (Сторонние):** `go.uber.org/zap`

---

## 📊 Прогресс

| Направление | Количество | Статус |
|-------------|------------|--------|
| **Задачи на Go** | 200+ | ✅ Активно решаю |
| **Задачи на SQL** | 100+ | ✅ В процессе изучения |
| **Сертификаты** | 6 из 7 | ✅ Получены |
| **Текущий курс** | Веб-разработка на Go | 🔄 В процессе |
| **Свои проекты** | 1 | ✅ Первый готов! |

---

## 🚀 Мои проекты

### 🧮 calc-api

REST API калькулятор на Go. Принимает JSON, считает в горутине, логирует запросы через middleware.

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![HTTP](https://img.shields.io/badge/HTTP-005571?style=flat)
![JSON](https://img.shields.io/badge/JSON-000000?style=flat&logo=json&logoColor=white)

[**→ Посмотреть проект**](https://github.com/EgorGo-dev/calc-api)

---

## 📝 Решённые задачи

<details>
<summary><b>📋 Контрольные работы и сложные задачи</b> (нажмите, чтобы развернуть)</summary>

## 📑 Оглавление

- 📘 [Задачи 1–30 (Go)](#tasks-1-30) — базовые и продвинутые концепции Go
- 🧩 [Задачи 31–60 (Go + SQLite)](#tasks-31-60) — переход к работе с базами данных
- 🗄️ [Задачи 61–90 (SQLite)](#tasks-61-90) — чистый SQL, JOIN, оконные функции и CTE
- 🧪 [Задачи 91–100 (Тестирование)](#tasks-91-100) — unit-тесты, table-driven, httptest, бенчмарки

---

<a id="tasks-1-30"></a>
## Задачи 1–30 (Go)

| № | Тема | Условие | Решение |
|---|------|---------|---------|
| &nbsp;1 | 🔖 Интерфейсы | [📄 условие](./work/task1/task1.md) | [💻 Go-решение](./work/task1/task1.go) |
| &nbsp;2 | 🔖 Строки | [📄 условие](./work/task2/task2.md) | [💻 Go-решение](./work/task2/task2.go) |
| &nbsp;3 | 🔖 Время | [📄 условие](./work/task3/task3.md) | [💻 Go-решение](./work/task3/task3.go) |
| &nbsp;4 | 🔖 Массивы | [📄 условие](./work/task4/task4.md) | [💻 Go-решение](./work/task4/task4.go) |
| &nbsp;5 | 🔖 Форматирование | [📄 условие](./work/task5/task5.md) | [💻 Go-решение](./work/task5/task5.go) |
| &nbsp;6 | 🔖 Сортировка | [📄 условие](./work/task6/task6.md) | [💻 Go-решение](./work/task6/task6.go) |
| &nbsp;7 | 🔖 Потоки | [📄 условие](./work/task7/task7.md) | [💻 Go-решение](./work/task7/task7.go) |
| &nbsp;8 | 🔖 Контекст | [📄 условие](./work/task8/task8.md) | [💻 Go-решение](./work/task8/task8.go) |
| &nbsp;9 | 🔖 HTTP | [📄 условие](./work/task9/task9.md) | [💻 Go-решение](./work/task9/task9.go) |
| 10 | 🔖 HTTP | [📄 условие](./work/task10/task10.md) | [💻 Go-решение](./work/task10/task10.go) |
| 11 | 🔖 HTTP | [📄 условие](./work/task11/task11.md) | [💻 Go-решение](./work/task11/task11.go) |
| 12 | 🔖 HTTP | [📄 условие](./work/task12/task12.md) | [💻 Go-решение](./work/task12/task12.go) |
| 13 | 🔖 Таймауты | [📄 условие](./work/task13/task13.md) | [💻 Go-решение](./work/task13/task13.go) |
| 14 | 🔖 Каналы | [📄 условие](./work/task14/task14.md) | [💻 Go-решение](./work/task14/task14.go) |
| 15 | 🔖 Каналы | [📄 условие](./work/task15/task15.md) | [💻 Go-решение](./work/task15/task15.go) |
| 16 | 🔖 Контекст | [📄 условие](./work/task16/task16.md) | [💻 Go-решение](./work/task16/task16.go) |
| 17 | 🔖 Тестирование | [📄 условие](./work/task17/task17.md) | [💻 Go-решение](./work/task17/task17.go) |
| 18 | 🔖 Тестирование | [📄 условие](./work/task18/task18.md) | [💻 Go-решение](./work/task18/task18.go) |
| 19 | 🔖 Дженерики | [📄 условие](./work/task19/task19.md) | [💻 Go-решение](./work/task19/task19.go) |
| 20 | 🔖 Контекст | [📄 условие](./work/task20/task20.md) | [💻 Go-решение](./work/task20/task20.go) |
| 21 | 🔖 Параллелизм | [📄 условие](./work/task21/task21.md) | [💻 Go-решение](./work/task21/task21.go) |
| 22 | 🔖 Логирование | [📄 условие](./work/task22/task22.md) | [💻 Go-решение](./work/task22/task22.go) |
| 23 | 🔖 JSON | [📄 условие](./work/task23/task23.md) | [💻 Go-решение](./work/task23/task23.go) |
| 24 | 🔖 JSON | [📄 условие](./work/task24/task24.md) | [💻 Go-решение](./work/task24/task24.go) |
| 25 | 🔖 HTTP | [📄 условие](./work/task25/task25.md) | [💻 Go-решение](./work/task25/task25.go) |
| 26 | 🔖 HTTP | [📄 условие](./work/task26/task26.md) | [💻 Go-решение](./work/task26/task26.go) |
| 27 | 🔖 HTTP | [📄 условие](./work/task27/task27.md) | [💻 Go-решение](./work/task27/task27.go) |
| 28 | 🔖 Curl | [📄 условие](./work/task28/task28.md) | [💻 Go-решение](./work/task28/task28.go) |
| 29 | 🔖 Сервер | [📄 условие](./work/task29/task29.md) | [💻 Go-решение](./work/task29/task29.go) |
| 30 | 🔖 Сервер | [📄 условие](./work/task30/task30.md) | [💻 Go-решение](./work/task30/task30.go) |

---

<a id="tasks-31-60"></a>
## Задачи 31–60 (Go + SQLite)

| № | Тема | Условие | Решение |
|---|------|---------|---------|
| 31 | 🔖 Метрики | [📄 условие](./work/task31/task31.md) | [💻 Go-решение](./work/task31/task31.go) |
| 32 | 🔖 Сервер | [📄 условие](./work/task32/task32.md) | [💻 Go-решение](./work/task32/task32.go) |
| 33 | 🔖 Middleware | [📄 условие](./work/task33/task33.md) | [💻 Go-решение](./work/task33/task33.go) |
| 34 | 🔖 Middleware | [📄 условие](./work/task34/task34.md) | [💻 Go-решение](./work/task34/task34.go) |
| 35 | 🔖 Метрики | [📄 условие](./work/task35/task35.md) | [💻 Go-решение](./work/task35/task35.go) |
| 36 | 🔖 Авторизация | [📄 условие](./work/task36/task36.md) | [💻 Go-решение](./work/task36/task36.go) |
| 37 | 🔖 Арифметика | [📄 условие](./work/task37/task37.md) | [🗄️ SQLite-решение](./work/task37/task37.sql) |
| 38 | 🔖 Конкатенация | [📄 условие](./work/task38/task38.md) | [🗄️ SQLite-решение](./work/task38/task38.sql) |
| 39 | 🔖 Арифметика | [📄 условие](./work/task39/task39.md) | [🗄️ SQLite-решение](./work/task39/task39.sql) |
| 40 | 🔖 DISTINCT | [📄 условие](./work/task40/task40.md) | [🗄️ SQLite-решение](./work/task40/task40.sql) |
| 41 | 🔖 REPLACE | [📄 условие](./work/task41/task41.md) | [🗄️ SQLite-решение](./work/task41/task41.sql) |
| 42 | 🔖 Строки | [📄 условие](./work/task42/task42.md) | [🗄️ SQLite-решение](./work/task42/task42.sql) |
| 43 | 🔖 Строки | [📄 условие](./work/task43/task43.md) | [🗄️ SQLite-решение](./work/task43/task43.sql) |
| 44 | 🔖 Строки | [📄 условие](./work/task44/task44.md) | [🗄️ SQLite-решение](./work/task44/task44.sql) |
| 45 | 🔖 Строки | [📄 условие](./work/task45/task45.md) | [🗄️ SQLite-решение](./work/task45/task45.sql) |
| 46 | 🔖 Агрегаты | [📄 условие](./work/task46/task46.md) | [🗄️ SQLite-решение](./work/task46/task46.sql) |
| 47 | 🔖 Агрегаты | [📄 условие](./work/task47/task47.md) | [🗄️ SQLite-решение](./work/task47/task47.sql) |
| 48 | 🔖 Агрегаты | [📄 условие](./work/task48/task48.md) | [🗄️ SQLite-решение](./work/task48/task48.sql) |
| 49 | 🔖 Агрегаты | [📄 условие](./work/task49/task49.md) | [🗄️ SQLite-решение](./work/task49/task49.sql) |
| 50 | 🔖 Агрегаты | [📄 условие](./work/task50/task50.md) | [🗄️ SQLite-решение](./work/task50/task50.sql) |
| 51 | 🔖 Агрегаты | [📄 условие](./work/task51/task51.md) | [🗄️ SQLite-решение](./work/task51/task51.sql) |
| 52 | 🔖 Подзапросы | [📄 условие](./work/task52/task52.md) | [🗄️ SQLite-решение](./work/task52/task52.sql) |
| 53 | 🔖 Условия | [📄 условие](./work/task53/task53.md) | [🗄️ SQLite-решение](./work/task53/task53.sql) |
| 54 | 🔖 Условия | [📄 условие](./work/task54/task54.md) | [🗄️ SQLite-решение](./work/task54/task54.sql) |
| 55 | 🔖 LIKE | [📄 условие](./work/task55/task55.md) | [🗄️ SQLite-решение](./work/task55/task55.sql) |
| 56 | 🔖 LIKE | [📄 условие](./work/task56/task56.md) | [🗄️ SQLite-решение](./work/task56/task56.sql) |
| 57 | 🔖 Группировка | [📄 условие](./work/task57/task57.md) | [🗄️ SQLite-решение](./work/task57/task57.sql) |
| 58 | 🔖 CASE | [📄 условие](./work/task58/task58.md) | [🗄️ SQLite-решение](./work/task58/task58.sql) |
| 59 | 🔖 Корреляция | [📄 условие](./work/task59/task59.md) | [🗄️ SQLite-решение](./work/task59/task59.sql) |
| 60 | 🔖 CASE | [📄 условие](./work/task60/task60.md) | [🗄️ SQLite-решение](./work/task60/task60.sql) |

---

<a id="tasks-61-90"></a>
## Задачи 61–90 (SQLite)

| № | Тема | Условие | Решение |
|---|------|---------|---------|
| 61 | 🔖 JOIN | [📄 условие](./work/task61/task61.md) | [🗄️ SQLite-решение](./work/task61/task61.sql) |
| 62 | 🔖 LEFT JOIN | [📄 условие](./work/task62/task62.md) | [🗄️ SQLite-решение](./work/task62/task62.sql) |
| 63 | 🔖 JOIN | [📄 условие](./work/task63/task63.md) | [🗄️ SQLite-решение](./work/task63/task63.sql) |
| 64 | 🔖 JOIN | [📄 условие](./work/task64/task64.md) | [🗄️ SQLite-решение](./work/task64/task64.sql) |
| 65 | 🔖 JOIN | [📄 условие](./work/task65/task65.md) | [🗄️ SQLite-решение](./work/task65/task65.sql) |
| 66 | 🔖 JOIN | [📄 условие](./work/task66/task66.md) | [🗄️ SQLite-решение](./work/task66/task66.sql) |
| 67 | 🔖 Подзапрос | [📄 условие](./work/task67/task67.md) | [🗄️ SQLite-решение](./work/task67/task67.sql) |
| 68 | 🔖 CTE | [📄 условие](./work/task68/task68.md) | [🗄️ SQLite-решение](./work/task68/task68.sql) |
| 69 | 🔖 CTE | [📄 условие](./work/task69/task69.md) | [🗄️ SQLite-решение](./work/task69/task69.sql) |
| 70 | 🔖 CTE | [📄 условие](./work/task70/task70.md) | [🗄️ SQLite-решение](./work/task70/task70.sql) |
| 71 | 🔖 NOT EXISTS | [📄 условие](./work/task71/task71.md) | [🗄️ SQLite-решение](./work/task71/task71.sql) |
| 72 | 🔖 Оконные | [📄 условие](./work/task72/task72.md) | [🗄️ SQLite-решение](./work/task72/task72.sql) |
| 73 | 🔖 ROW_NUMBER | [📄 условие](./work/task73/task73.md) | [🗄️ SQLite-решение](./work/task73/task73.sql) |
| 74 | 🔖 Оконные | [📄 условие](./work/task74/task74.md) | [🗄️ SQLite-решение](./work/task74/task74.sql) |
| 75 | 🔖 RANK | [📄 условие](./work/task75/task75.md) | [🗄️ SQLite-решение](./work/task75/task75.sql) |
| 76 | 🔖 RANK | [📄 условие](./work/task76/task76.md) | [🗄️ SQLite-решение](./work/task76/task76.sql) |
| 77 | 🔖 ROW_NUMBER | [📄 условие](./work/task77/task77.md) | [🗄️ SQLite-решение](./work/task77/task77.sql) |
| 78 | 🔖 DENSE_RANK | [📄 условие](./work/task78/task78.md) | [🗄️ SQLite-решение](./work/task78/task78.sql) |
| 79 | 🔖 LAG | [📄 условие](./work/task79/task79.md) | [🗄️ SQLite-решение](./work/task79/task79.sql) |
| 80 | 🔖 LEAD | [📄 условие](./work/task80/task80.md) | [🗄️ SQLite-решение](./work/task80/task80.sql) |
| 81 | 🔖 DENSE_RANK | [📄 условие](./work/task81/task81.md) | [🗄️ SQLite-решение](./work/task81/task81.sql) |
| 82 | 🔖 DENSE_RANK | [📄 условие](./work/task82/task82.md) | [🗄️ SQLite-решение](./work/task82/task82.sql) |
| 83 | 🔖 ROW_NUMBER | [📄 условие](./work/task83/task83.md) | [🗄️ SQLite-решение](./work/task83/task83.sql) |
| 84 | 🔖 Группировка | [📄 условие](./work/task84/task84.md) | [🗄️ SQLite-решение](./work/task84/task84.sql) |
| 85 | 🔖 Группировка | [📄 условие](./work/task85/task85.md) | [🗄️ SQLite-решение](./work/task85/task85.sql) |
| 86 | 🔖 Подзапрос | [📄 условие](./work/task86/task86.md) | [🗄️ SQLite-решение](./work/task86/task86.sql) |
| 87 | 🔖 Группировка | [📄 условие](./work/task87/task87.md) | [🗄️ SQLite-решение](./work/task87/task87.sql) |
| 88 | 🔖 Оконные | [📄 условие](./work/task88/task88.md) | [🗄️ SQLite-решение](./work/task88/task88.sql) |
| 89 | 🔖 ROW_NUMBER | [📄 условие](./work/task89/task89.md) | [🗄️ SQLite-решение](./work/task89/task89.sql) |
| 90 | 🔖 LAG | [📄 условие](./work/task90/task90.md) | [🗄️ SQLite-решение](./work/task90/task90.sql) |

---

<a id="tasks-91-100"></a>
## Задачи 91–100 (Тестирование)

| № | Тема | Условие | Решение |
|---|------|---------|---------|
| &nbsp;91 | 🔖 Unit-тесты | [📄 условие](./work/task91/task91.md) | [🧪 тест](./work/task91/task91_test.go) |
| &nbsp;92 | 🔖 Table-driven | [📄 условие](./work/task92/task92.md) | [🧪 тест](./work/task92/task92_test.go) |
| &nbsp;93 | 🔖 Ошибки | [📄 условие](./work/task93/task93.md) | [🧪 тест](./work/task93/task93_test.go) |
| &nbsp;94 | 🔖 Структуры | [📄 условие](./work/task94/task94.md) | [🧪 тест](./work/task94/task94_test.go) |
| &nbsp;95 | 🔖 Интерфейсы | [📄 условие](./work/task95/task95.md) | [🧪 тест](./work/task95/task95_test.go) |
| &nbsp;96 | 🔖 JSON | [📄 условие](./work/task96/task96.md) | [🧪 тест](./work/task96/task96_test.go) |
| &nbsp;97 | 🔖 HTTP | [📄 условие](./work/task97/task97.md) | [🧪 тест](./work/task97/task97_test.go) |
| &nbsp;98 | 🔖 Каналы | [📄 условие](./work/task98/task98.md) | [🧪 тест](./work/task98/task98_test.go) |
| &nbsp;99 | 🔖 WaitGroup | [📄 условие](./work/task99/task99.md) | [🧪 тест](./work/task99/task99_test.go) |
| 100 | 🔖 Бенчмарки | [📄 условие](./work/task100/task100.md) | [🧪 тест](./work/task100/task100_test.go) |

---

</details>

<br>

<details>
<summary><b>💀🔥 Сверхсложные задачи — БОСС-УРОВЕНЬ</b> (нажмите, чтобы развернуть)</summary>

> ⚠️ **Внимание!** Эти задачи требуют нестандартного мышления, глубокого понимания **Go** 💻 или **SQL (SQLite/PostgreSQL)** 🗄️ и железных нервов.  
> Подходите с осторожностью – здесь не место для слабаков! ⚔️

| ⚡ Задача | 🔖 Тема | 📄 Условие | 💻/🗄️ Решение | 📌 Язык | 🧠 Сложность |
|-----------|---------|------------|----------------|---------|--------------|
| **1.** | 🔖 REST API | 📄 [условие](./work/HARDtask1/hardTask.md) | 💻 [решение](./work/HARDtask1/hardTask.go) | **Go** | 🧠🧠🧠 |
| **2.** | 🔖 Подзапросы | 📄 [условие](./work/HARDtask2/hardTask2.md) | 🗄️ [решение](./work/HARDtask2/hardTask2.sql) | **SQLite** | 🧠 |
| **3.** | 🔖 Оконные | 📄 [условие](./work/HARDtask3/hardTask3.md) | 🗄️ [решение](./work/HARDtask3/hardTask3.sql) | **PostgreSQL** | 🧠🧠🧠🧠 |
| **4.** | 🔖 Параллелизм | 📄 [условие](./work/HARDtask4/hardTask4.md) | 💻 [решение](./work/HARDtask4/hardTask4.go) | **Go** | 🧠🧠🧠🧠 |
| **5.** | — | — | — | — | — |

</details>

<br>

<details>
<summary><b>📌 Мелкие задачи</b> (нажмите, чтобы развернуть)</summary>

Решено **более 150** небольших задач на **Go** и **SQL (SQLite + PostgreSQL)**:

**Go:**  
`циклы`, `условия`, `функции`, `строки`, `массивы`, `слайсы`, `мапы`, `структуры`, `методы`, `конструкторы`, `интерфейсы`, `ошибки`, `тестирование`, **горутины**, **каналы**, **синхронизация** (`sync.Mutex`, `sync.WaitGroup`), **контекст** (`context`), **тестирование конкурентного кода** (`go test -race`).

**SQL (общие темы для SQLite и PostgreSQL):**  
- **Основы:** `SELECT`, `FROM`, `AS`, `WHERE`, `AND`, `OR`, `NOT`, `IN`, `NOT IN`, `BETWEEN`, `LIKE`, `IS NULL`, `IS NOT NULL`, `ORDER BY`, `LIMIT`, `OFFSET`, `DISTINCT`.  
- **Агрегация:** `COUNT`, `SUM`, `AVG`, `MIN`, `MAX`, `GROUP BY`, `HAVING`.  
- **Строки:** `||`, `LENGTH`, `LOWER`, `UPPER`, `SUBSTR` (в SQLite) / `SUBSTRING` (в PostgreSQL), `POSITION`.  
- **Округление и преобразование:** `ROUND`, `CAST`.  
- **Условная логика:** `CASE WHEN`.  
- **Подзапросы:** скалярные, `IN`, `EXISTS`, `NOT EXISTS`, подзапросы в `FROM`.  
- **Соединения:** `INNER JOIN`, `LEFT JOIN`, `SELF JOIN`.  
- **CTE:** `WITH ... AS`, несколько CTE.  
- **Оконные функции:** `ROW_NUMBER`, `RANK`, `DENSE_RANK`, `SUM() OVER`, `AVG() OVER`, `LAG`, `LEAD`, `PARTITION BY`, `ORDER BY` в окне.  
- **Временные таблицы.**

**Специфическое для SQLite:**  
- Даты: `strftime`, `JULIANDAY`.  
- Ограничения: `LIMIT`/`OFFSET` (работает чуть иначе, чем в PostgreSQL).

**Специфическое для PostgreSQL:**  
- Даты: `EXTRACT`, `to_char`, `NOW()`.  
- Генерация ID: `SERIAL`, `BIGSERIAL`.  
- Возврат данных: `RETURNING`.  
- Тип данных: `JSONB`.  
- Работа с дубликатами, выбор максимума через подзапросы, ранжирование внутри групп.

*(Я не выложил их отдельно, потому что их очень много.)*

</details>

---

## 🎯 Мои цели

- [x] Пройти Go-1, Go-2, Go-3, SQL, Мат. мышление, PostgreSQL
- [x] Написать свой первый проект — [calc-api](https://github.com/EgorGo-dev/calc-api) 🎉
- [ ] Освоить веб-разработку на Go
- [ ] Изучить Docker
- [ ] Познакомиться с Kubernetes
- [ ] Создать свой первый веб-сервис на Go

---

## 🏆 Сертификаты

**Успешно завершенные курсы:**
| 🏆 Go-1 | 🏆 Go-2 | 🏆 Go-3 |
|---------|---------|---------|
| <img src="./certificate/certificate.jpg" width="300" alt="Сертификат Go-1"> <br> **Егор Геннадьевич Гордеев** <br> self-paced курс | <img src="./certificate/certificate 2.jpg" width="300" alt="Сертификат Go-2"> <br> **Егор Геннадьевич Гордеев** <br> self-paced курс | <img src="./certificate/certificate 3.jpg" width="300" alt="Сертификат Go-3"> <br> **Егор Геннадьевич Гордеев** <br> self-paced курс |
| <div align="center">🏆 **SQL**</div> | <div align="center">🏆 **Мат. мышление**</div> | <div align="center">🏆 **PostgreSQL**</div> |
| <img src="./certificate/certificate 4.jpg" width="300" alt="Сертификат SQL"> <br> **Егор Геннадьевич Гордеев** <br> self-paced курс | <img src="./certificate/certificate 5.jpg" width="300" alt="Сертификат Мат. мышление"> <br> **Егор Геннадьевич Гордеев** <br> self-paced курс | <img src="./certificate/certificate 6.jpg" width="300" alt="Сертификат PostgreSQL"> <br> **Егор Геннадьевич Гордеев** <br> self-paced курс |

**Планируемые курсы:**  
🚀 `Веб-разработка на Go` · `Docker` · `Kubernetes`

*Сертификаты об успешном завершении курсов:*
- **Go-1:** Разработка на Go: основы  
- **Go-2:** Разработка на Go: продвинутый уровень  
- **Go-3:** Разработка на Go: продолжение  
- **SQL:** SQL: шаг за шагом  
- **Математическое мышление для программистов**  
- **PostgreSQL**  

*Планируемые курсы: Веб-разработка на Go, Docker, Kubernetes*

---

## 🙏 Благодарности

- **[Яндекс Лицей](https://lyceum.yandex.ru/)** за отличные курсы.
- Всем, кто поддерживает меня звёздочками и добрыми словами.

---

## 📬 Контакты

[![GitHub](https://img.shields.io/badge/GitHub-EgorGo--dev-181717?style=for-the-badge&logo=github)](https://github.com/EgorGo-dev)  
[![Email](https://img.shields.io/badge/Email-gpp220512@gmail.com-D14836?style=for-the-badge&logo=gmail)](mailto:gpp220512@gmail.com)

---

⭐ Если этот репозиторий показался вам полезным — поставьте звёздочку!  
Спасибо, что заглянули! 😊
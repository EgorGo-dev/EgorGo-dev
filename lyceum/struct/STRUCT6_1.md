# 🗂 Структура проекта 6.1: PostgreSQL и Docker

*   📁 **[6_1/](../6_1/)** — *Корневая папка задачи*
    *   📁 **[config/](../6_1/config/)** — *Модуль конфигурации*
        *   📄 [config.go](../6_1/config/config.go) — *Чтение параметров БД и портов*
    *   📁 **[google/](../6_1/google/)** — *Официальные прото-файлы Google*
        *   📁 [api/](../6_1/google/api/)
    *   📁 **[pb/](../6_1/pb/)** — *Сгенерированный код (gRPC + Gateway)*
        *   📄 [order.pb.go](../6_1/pb/order.pb.go)
        *   📄 [order_grpc.pb.go](../6_1/pb/order_grpc.pb.go)
        *   📄 [order.pb.gw.go](../6_1/pb/order.pb.gw.go)
    *   📁 **[server/](../6_1/server/)** — *Ядро приложения*
        *   📁 **[migrations/](../6_1/server/migrations/)** — *SQL-миграции*
            *   📄 [001_init.sql](../6_1/server/migrations/001_init.sql) — *Создание таблицы `orders`*
        *   📄 [handlers.go](../6_1/server/handlers.go) — *Связка gRPC-методов и слоя `store`*
        *   📄 [main.go](../6_1/server/main.go) — *Запуск сервера, подключение к БД, миграции, shutdown*
    *   📁 **[store/](../6_1/store/)** — *Слой работы с базой данных*
        *   📄 [store.go](../6_1/store/store.go) — *Пул соединений `pgx` и SQL-запросы (CRUD)*
    *   📄 [.env.example](../6_1/.env.example) — *Пример переменных*
    *   📄 [docker-compose.yml](../6_1/docker-compose.yml) — *Оркестрация (App + PostgreSQL)*
    *   📄 [Dockerfile](../6_1/Dockerfile) — *Сборка контейнера*
    *   📄 [order.proto](../6_1/order.proto)
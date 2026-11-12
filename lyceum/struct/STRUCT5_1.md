# 🗂 Структура проекта 5.1: grpc-gateway и Graceful Shutdown

*   📁 **[5_1/](../5_1/)** — *Корневая папка задачи*
    *   📁 **[config/](../5_1/config/)** — *Модуль конфигурации*
        *   📄 [config.go](../5_1/config/config.go) — *Чтение портов gRPC и HTTP из `.env`*
    *   📁 **[google/](../5_1/google/)** — *Официальные прото-файлы Google*
        *   📁 [api/](../5_1/google/api/) — *Аннотации для HTTP-маршрутов*
    *   📁 **[pb/](../5_1/pb/)** — *Сгенерированный код*
        *   📄 [order.pb.go](../5_1/pb/order.pb.go)
        *   📄 [order_grpc.pb.go](../5_1/pb/order_grpc.pb.go)
        *   📄 [order.pb.gw.go](../5_1/pb/order.pb.gw.go) — *Сгенерированный HTTP-шлюз (Gateway)*
    *   📁 **[server/](../5_1/server/)** — *Серверная логика*
        *   📄 [main.go](../5_1/server/main.go) — *Запуск двух серверов (gRPC + HTTP) и Graceful Shutdown*
    *   📄 [.env.example](../5_1/.env.example)
    *   📄 [order.proto](../5_1/order.proto)
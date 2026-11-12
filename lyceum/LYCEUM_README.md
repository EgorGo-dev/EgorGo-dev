# Go course project

Репозиторий для заданий курса Go. Каждое задание — в своей папке `<лекция>_<задание>` (например, `3_1`).

## Подготовка

```bash
# protoc: https://protobuf.dev/installation/
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Задание 3.1 — gRPC OrderService

Proto-файл: `3_1/order.proto`. **Не редактируйте и не перемещайте его** — изменённый proto обнуляет проверку.

Генерация кода (из папки `3_1`, результат попадёт в `3_1/pb`):

```bash
cd 3_1
protoc --go_out=. --go-grpc_out=. order.proto
```

Требования, которые проверяет система:

- `go.mod` — в корне репозитория или в папке `3_1`;
- сгенерированные `pb/order.pb.go` и `pb/order_grpc.pb.go` закоммичены;
- в `3_1` есть `package main` с `func main()` (рекомендуется `3_1/server/main.go`), сервер слушает порт **50051**;
- хранилище — `map`, защищённая `sync.Mutex` / `sync.RWMutex` (или `sync.Map`), без гонок данных (`go build -race`);
- для несуществующего id методы `GetOrder`, `UpdateOrder`, `DeleteOrder` возвращают ошибку `codes.NotFound` с понятным сообщением;
- `UpdateOrder` не создаёт заказ, если его нет;
- код отформатирован `gofmt`, `go vet` без замечаний;
- в репозитории нет бинарников и папок IDE.

Запуск:

```bash
go run ./3_1/server
```

## Задание 4.1 — Конфигурация, Интерсепторы и Makefile

### Переменные окружения
Сервис использует следующие переменные:
* `GRPC_PORT` — порт, на котором запускается gRPC-сервер (по умолчанию `50051`).

### Настройка и запуск
1. Создайте файл `.env` на основе примера:
   ```bash
   cp 4_1/.env.example 4_1/.env
   ```

2. Используйте `Makefile` для управления проектом:
   * Сгенерировать gRPC код: `make generate`
   * Собрать бинарник: `make build`
   * Запустить сервер: `make run`

## Задание 5.1 — grpc-gateway, graceful shutdown

Папка `5_1`, proto `5_1/order.proto` с HTTP-аннотациями (**не редактировать**), `google/api/*.proto` уже лежат рядом.

```bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
cd 5_1
protoc -I . --go_out=. --go-grpc_out=. --grpc-gateway_out=. order.proto
```

| Метод | HTTP |
|---|---|
| CreateOrder | `POST /v1/orders` `{"item": "...", "quantity": 1}` |
| GetOrder | `GET /v1/orders/{id}` |
| UpdateOrder | `PUT /v1/orders/{id}` `{"item": "...", "quantity": 1}` |
| DeleteOrder | `DELETE /v1/orders/{id}` |
| ListOrders | `GET /v1/orders` |

- порты: gRPC — `GRPC_PORT`, HTTP — `HTTP_PORT`;
- по `SIGTERM` и `SIGINT` сервер корректно останавливается (`GracefulStop`, `http.Server.Shutdown`) и **печатает `Server Stopped`**;
- `Makefile` генерирует gRPC **и** grpc-gateway код, собирает и запускает сервер.

## Задание 6.1 — Docker, docker-compose, PostgreSQL

Папка `6_1`, proto `6_1/order.proto` (как в 5.1, **не редактировать**).

- `6_1/Dockerfile` — multi-stage, финальный образ alpine/distroless/scratch, `go.mod`/`go.sum` + `go mod download` отдельным слоем до `COPY . .`;
- `6_1/docker-compose.yml` — сервис (`build`) + `postgres`, общая сеть, `healthcheck` с `pg_isready`, `depends_on: condition: service_healthy`, volume `postgres_data` → `/var/lib/postgresql/data`, переменные из `.env`;
- `.env.example` с `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, `POSTGRES_HOST`, `POSTGRES_PORT`, `GRPC_PORT`, `HTTP_PORT`;
- заказы хранятся в PostgreSQL и не теряются после перезапуска сервиса;
- схему БД создавайте SQL-файлами в `6_1/migrations/*.sql` (или из кода при старте);
- graceful shutdown с выводом `Server Stopped` — как в 5.1.
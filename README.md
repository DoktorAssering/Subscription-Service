# Subscription Service

REST-сервис на Go для управления и агрегации данных об онлайн-подписках пользователей.

---

## 🚀 Быстрый старт

### 1. Клонирование репозитория

```bash
git clone https://github.com/your-org/subscription-service.git
cd subscription-service
```

### 2. Создание файла окружения

Скопируйте файл примера и заполните переменные:

```bash
cp .env.example .env
# или создайте вручную .env с содержимым:
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=postgres
# DB_NAME=subscriptions
```

### 3. Запуск базы данных через Docker

```bash
docker compose up -d
```

Проверьте, что контейнер PostgreSQL запущен и слушает порт 5432 (Либо любой другой указанный в .env + docker-compose.yml).

### 4. Установка зависимостей и генерация Swagger

```bash
# Скачиваем модули Go
go mod tidy

# Генерируем документацию Swagger (точка входа в cmd/main.go)
swag init -g cmd/main.go
```

### 5. Сборка и запуск приложения

```bash
# Запуск прямо через Go
go run ./cmd/main.go
```

Сервис будет доступен на `http://localhost:8080`.

---

## 📂 Структура проекта

```
subscription-service/
├── cmd/
│   └── main.go            # Точка входа
├── config/
│   └── config.go          # Настройка БД и окружения
├── handler/
│   └── subscription.go    # HTTP-ручки
├── model/
│   └── subscription.go    # Модель данных
├── repository/
│   └── subscription.go    # Работа с БД
├── service/
│   └── subscription.go    # Бизнес-логика
├── docs/                  # Сгенерированные Swagger-файлы
├── Dockerfile             # Сборка Go-приложения
├── docker-compose.yml     # Подъём БД
├── .env.example           # Пример файла окружения
├── go.mod
└── README.md
```

---

## 📋 API Endpoints

|  Метод | Путь                   | Описание                        |
| -----: | :--------------------- | :------------------------------ |
|   POST | `/subscriptions`       | Создать подписку                |
|    GET | `/subscriptions`       | Получить список подписок        |
|    GET | `/subscriptions/{id}`  | Получить подписку по ID         |
|    PUT | `/subscriptions/{id}`  | Обновить подписку               |
| DELETE | `/subscriptions/{id}`  | Удалить подписку                |
|    GET | `/subscriptions/total` | Суммарная стоимость по фильтрам |

### Пример тела запроса (POST /subscriptions)

```json
{
  "service_name": "Netflix",
  "price": 12.99,
  "user_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "start_date": "2025-07-22"
}
```

### Пример агрегирующего запроса

```
GET /subscriptions/total?service=Netflix&userId=3fa85f64-5717-4562-b3fc-2c963f66afa6
```

Ответ:

```json
{ "total_price": 12.99 }
```

---

## 🔧 Тестирование

* **Swagger UI**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

* **curl**:

  ```bash
  curl -X POST http://localhost:8080/subscriptions \
    -H "Content-Type: application/json" \
    -d '{ "service_name":"Netflix","price":12.99,"user_id":"3fa85f64-5717-4562-b3fc-2c963f66afa6","start_date":"2025-07-22" }'
  ```

* **Postman**: импортируйте коллекцию из папки `postman/Subscription API.json` (если есть).

---

## 🛠️ Отладка и советы

* После изменения модели перезапускайте приложение для авто-миграции (`db.AutoMigrate`).
* Проверяйте переменные окружения и совпадение портов `DB_HOST/DB_PORT`.
* В случае ошибок модулей Go — запускайте `go clean -modcache && go mod tidy`.

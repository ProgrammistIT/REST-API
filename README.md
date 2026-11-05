# Threat API

Небольшой REST API для хранения информации об угрозах по MD5-хешу файла.

Написан на Go с использованием PostgreSQL. API реализован на основе OpenAPI 3.0.

## Возможности

- `POST /info` — добавить угрозу
- `GET /info/{md5}` — получить угрозу по MD5
- `PUT /info/{md5}` — обновить информацию об угрозе

Для каждой угрозы хранятся MD5, название(name), описание(description), уровень опасности(severity) и время обнаружения(detected_time).

Уровень опасности может быть `low`, `medium`, `high` или `critical`.

## Запуск

Для запуска нужны Docker и Docker Compose.

```bash
docker-compose up --build -d
```

Сервер будет доступен по адресу `http://localhost:8080`.

Остановить контейнеры:

```bash
docker-compose down
```

Также можно запустить проект локально, предварительно настроив PostgreSQL и выполнив `init.sql`:

```bash
go run .
```

Строку подключения можно указать через переменную окружения `DATABASE_URL`.

## Примеры запросов

**Создание угрозы:**

```bash
curl -X POST http://localhost:8080/info \
  -H "Content-Type: application/json" \
  -d '{
    "md5": "d41d8cd98f00b204e9800998ecf8427e",
    "name": "Test Trojan",
    "description": "Test threat",
    "severity": "high"
  }'
```

**Получение угрозы:**

```bash
curl http://localhost:8080/info/d41d8cd98f00b204e9800998ecf8427e
```

**Обновление угрозы:**

```bash
curl -X PUT http://localhost:8080/info/d41d8cd98f00b204e9800998ecf8427e \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Trojan",
    "severity": "critical"
  }'
```

## Тесты

Для запуска тестов:

```bash
go test ./...
```

Добавлены тесты проверки MD5 и JSON-ответов с ошибками.

## Структура проекта

- `internal/model` — модель угрозы
- `internal/validator` — проверка данных
- `internal/database` — подключение к PostgreSQL
- `internal/handler` — обработчики HTTP-запросов
- `internal/error` — обработчики ошибок
- `init.sql` — создание таблицы в базе данных
- `Dockerfile`, `compose.yaml` — запуск через Docker

# soho
Game Service API

Описание

Game Service - микросервис, отвечающий за логику игры "Сапер с элементами матч-3". Он обрабатывает игровые действия, генерирует игровое поле и взаимодействует с другими сервисами (User Service, Task Service, Leaderboard Service).

Функциональность

Создание игровой сессии

Получение информации об игровой сессии

Обработка действий пользователя (открытие клеток, сбор спрайтов)

Начисление монет и бонусов

Используемые технологии

Язык: Go 1.21+

Веб-фреймворк: Gorilla Mux

База данных: PostgreSQL

Кэш: Redis

Контейнеризация: Docker

Установка и запуск

1. Клонирование репозитория

git clone https://github.com/yourrepo/game-service.git
cd game-service

2. Установка зависимостей

go mod tidy

3. Настройка переменных окружения

Создайте .env файл и добавьте:

DATABASE_URL=postgres://user:password@localhost:5432/gamedb?sslmode=disable
REDIS_URL=redis://localhost:6379
PORT=8080

4. Запуск сервера

go run main.go

API эндпоинты

1. Создание игровой сессии

POST /game?user_id={user_id}

Ответ:

{
    "id": "session_uuid",
    "user_id": "123",
    "status": "active"
}

2. Получение информации об игровой сессии

GET /game/{id}

Ответ:

{
    "id": "session_uuid",
    "user_id": "123",
    "status": "active",
    "field": [[0,1,0],[1,0,1]]
}

Запуск в Docker

Создайте Dockerfile:

FROM golang:1.21 AS build
WORKDIR /app
COPY . .
RUN go build -o game-service

FROM alpine
WORKDIR /app
COPY --from=build /app/game-service .
CMD ["./game-service"]

Соберите и запустите контейнер:

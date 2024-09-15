## Домашнее задание №15 «Go-клиент для БД онлайн магазина»

- Создайте пакет для работы в БД
- Реализуйте объект подключния 
- Реализуйте функции для выполнения запросов из предыдущего ДЗ
- Обеспечьте атомарность выполнения запросов с помозью транзакций
- В качестве транспортного слоя используйте серверную часть из ДЗ №13

### Критерии оценки
- Понятность и чистота кода - до 2 баллов
- Реализован слой взаимодействия с БД - 2 балла
- Используются транзакции - 2 балла
- Сервис отдает данные по HTTP - 2 балла

#### Зачёт от 6 баллов


# HW15 Go SQL

Это проект для выполнения домашнего задания №15 по курсу Golang от Otus. 
В этом проекте реализована работа с базой данных PostgreSQL с использованием Go. 
Проект включает в себя функции для вставки, обновления, удаления и выборки данных, а также создание индексов для ускорения выборки. 
Также добавлено CRUD-приложение, которое позволяет обращаться к методам из `db.go` по HTTP-запросам.

## Структура проекта

```
myapp/
├── main.go
├── config/
│   └── config.go
├── db/
│   └── db.go
├── handlers/
│   └── handlers.go
└── config.yml
```

## Установка и запуск

### Шаг 1: Клонирование репозитория

Клонируйте репозиторий на ваш локальный компьютер:

```sh
git clone https://github.com/ddpakhomov/home-work-golang-otus/hw15_go_sql.git
cd hw15_go_sql
```

### Шаг 2: Установка зависимостей

Установите необходимые зависимости:

```sh
go mod tidy
```

### Шаг 3: Настройка базы данных

Создайте файл `config.yml` в корневой директории проекта и добавьте следующие настройки:

```yaml
database:
  user: "postgres"
  password: ""
  dbname: "postgres-otus"
  host: "localhost"
  port: 5432
  sslmode: "disable"
```

### Шаг 4: Запуск контейнера PostgreSQL

Создайте файл `docker-compose.yml` в корневой директории проекта и добавьте следующие настройки:

```yaml
version: '3.9'
services:
  postgres-otus:
    image: postgres:16.2-alpine
    container_name: postgres-otus
    environment:
      POSTGRES_DB: "postgres-otus"
      POSTGRES_USER: "postgres"
      POSTGRES_PASSWORD: ""
      PGDATA: "/var/lib/postgresql/data/pgdata"
    volumes:
      - .data:/var/lib/postgresql/data
    ports:
      - "5432:5432"
```

Запустите контейнер с PostgreSQL:

```sh
docker-compose up -d
```

### Шаг 5: Запуск приложения

Запустите приложение:

```sh
go run main.go
```

## HTTP API

### Users

#### Получение списка пользователей

```sh
curl -X GET http://localhost:8080/users
```

#### Создание пользователя

```sh
curl -X POST http://localhost:8080/users -d '{
    "name": "John",
    "email": "john@example.com",
    "password": "password123"
}' -H "Content-Type: application/json"
```

#### Обновление пользователя

```sh
curl -X PUT http://localhost:8080/users -d '{
    "id": 1,
    "name": "John Smith",
    "email": "john.smith@example.com"
}' -H "Content-Type: application/json"
```

#### Удаление пользователя

```sh
curl -X DELETE http://localhost:8080/users?id=1
```

### Products

#### Получение списка продуктов

```sh
curl -X GET http://localhost:8080/products
```

#### Создание продукта

```sh
curl -X POST http://localhost:8080/products -d '{
    "name": "Product A",
    "price": 29.99
}' -H "Content-Type: application/json"
```

#### Обновление продукта

```sh
curl -X PUT http://localhost:8080/products -d '{
    "id": 1,
    "name": "Product B",
    "price": 39.99
}' -H "Content-Type: application/json"
```

#### Удаление продукта

```sh
curl -X DELETE http://localhost:8080/products?id=1
```

### Orders

#### Получение списка заказов по пользователю

```sh
curl -X GET http://localhost:8080/orders?user_id=1
```

#### Создание заказа

```sh
curl -X POST http://localhost:8080/orders -d '{
    "user_id": 1,
    "order_date": "2023-10-01 10:00:00",
    "total_amount": 59.98,
    "order_products": [
        {
            "product_id": 1,
            "quantity": 2
        }
    ]
}' -H "Content-Type: application/json
```

#### Удаление заказа
```sh
curl -X DELETE http://localhost:8080/orders?id=1
```

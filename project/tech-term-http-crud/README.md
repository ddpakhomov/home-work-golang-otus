# CRUD Приложение на Go с использованием MongoDB

Это приложение представляет собой веб-сервер, написанный на языке Go, который реализует CRUD (Create, Read, Update, Delete) операции для управления командами, хранящимися в базе данных MongoDB.

## Требования

- Go 1.16 или новее
- MongoDB

## Установка

1. Склонируйте репозиторий:

    ```sh
    git clone https://github.com/yourusername/yourrepository.git
    cd yourrepository
    ```

2. Установите зависимости:

    ```sh
    go mod tidy
    ```

3. Создайте файл конфигурации `config.yml` в корне проекта:

    ```yaml
    mongodb:
      uri: "mongodb://localhost:27017"
      database: "mydatabase"
      collection: "commands"
    ```

## Запуск

1. Запустите MongoDB сервер, если он еще не запущен.
2. Запустите приложение:

    ```sh
    go run main.go
    ```

Сервер будет слушать запросы на порту `8080`.

## API Эндпоинты
## API Маршруты

### Создание термина

- **URL:** `/terms`
- **Метод:** `POST`
- **Тело запроса:**
    ```json
    {
        "term": "example-term",
        "name": "Example Term",
        "description": "This is an example term",
        "theme": "example-theme"
    }
    ```
- **Ответ:**
    ```json
    {
        "InsertedID": "60d5f9b9f1d3f8b5a7c7b5c9"
    }
    ```

### Создание термина с темой

- **URL:** `/terms/{theme}`
- **Метод:** `POST`
- **Тело запроса:**
    ```json
    {
        "term": "example-term",
        "name": "Example Term",
        "description": "This is an example term",
        "theme": "example-theme"
    }
    ```
- **Ответ:**
    ```json
    {
        "InsertedID": "60d5f9b9f1d3f8b5a7c7b5c9"
    }
    ```

### Получение всех терминов

- **URL:** `/terms`
- **Метод:** `GET`
- **Ответ:**
    ```json
    [
        {
            "id": "60d5f9b9f1d3f8b5a7c7b5c9",
            "term": "example-term",
            "name": "Example Term",
            "description": "This is an example term",
            "theme": "example-theme"
        }
    ]
    ```

### Получение терминов по теме

- **URL:** `/terms/theme/{theme}`
- **Метод:** `GET`
- **Ответ:**
    ```json
    [
        {
            "id": "60d5f9b9f1d3f8b5a7c7b5c9",
            "term": "example-term",
            "name": "Example Term",
            "description": "This is an example term",
            "theme": "example-theme"
        }
    ]
    ```

### Получение термина по имени

- **URL:** `/terms/{term}`
- **Метод:** `GET`
- **Ответ:**
    ```json
    {
        "id": "60d5f9b9f1d3f8b5a7c7b5c9",
        "term": "example-term",
        "name": "Example Term",
        "description": "This is an example term",
        "theme": "example-theme"
    }
    ```

### Получение термина по теме и имени

- **URL:** `/terms/{theme}/{term}`
- **Метод:** `GET`
- **Ответ:**
    ```json
    {
        "id": "60d5f9b9f1d3f8b5a7c7b5c9",
        "term": "example-term",
        "name": "Example Term",
        "description": "This is an example term",
        "theme": "example-theme"
    }
    ```

### Обновление термина

- **URL:** `/terms/{theme}/{term}`
- **Метод:** `PUT`
- **Тело запроса:**
    ```json
    {
        "term": "updated-term",
        "name": "Updated Term",
        "description": "This is an updated term",
        "theme": "updated-theme"
    }
    ```
- **Ответ:**
    ```json
    {
        "id": "60d5f9b9f1d3f8b5a7c7b5c9",
        "term": "updated-term",
        "name": "Updated Term",
        "description": "This is an updated term",
        "theme": "updated-theme"
    }
    ```

### Удаление термина

- **URL:** `/terms/{term}`
- **Метод:** `DELETE`
- **Ответ:**
    ```json
    {
        "message": "Term 'example-term' in all themes successfully deleted"
    }
    ```

### Удаление термина по теме

- **URL:** `/terms/{theme}/{term}`
- **Метод:** `DELETE`
- **Ответ:**
    ```json
    {
        "message": "Term 'example-term' in theme 'example-theme' successfully deleted"
    }
    ```

## Основные компоненты

### Конфигурация

Конфигурация загружается из файла `config.yml` с помощью функции `LoadConfig`, расположенной в файле `config/config.go`. В конфигурации содержатся параметры подключения к MongoDB.

### Подключение к MongoDB

В функции `main` создается контекст с тайм-аутом в 10 секунд для подключения к MongoDB. Параметры подключения берутся из загруженной конфигурации. Используется клиент MongoDB для подключения к базе данных.

### Создание маршрутизатора

Используется библиотека `mux` для создания маршрутизатора, который обрабатывает HTTP-запросы. Определяются маршруты для различных CRUD операций: создание, получение всех команд, получение конкретной команды, обновление и удаление.

### Запуск сервера

Сервер запускается и начинает слушать запросы на порту `8080`.

## Зависимости

- `github.com/gorilla/mux`: для маршрутизации HTTP-запросов.
- `go.mongodb.org/mongo-driver/mongo`: для работы с MongoDB.
- `gopkg.in/yaml.v2`: для работы с YAML файлами.

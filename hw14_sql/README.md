## Домашнее задание №14 «БД онлайн магазина»

- Установите PostgreSQL
- Создайте БД и требуемых пользователей
- Создайте схему БД, содержащую следующие таблицы
1) Таблица ""Пользователи"" (Users) содержит информацию о пользователях, включая их уникальные идентификаторы (id), имена (name), электронные адреса (email) и пароли (password).
2) Таблица ""Заказы"" (Orders) отображает информацию о заказах, включая идентификаторы заказов (id), идентификаторы пользователей (user_id), даты заказов (order_date), общую стоимость заказов (total_amount). Связь между таблицами ""Пользователи"" и ""Заказы"" реализована через внешний ключ (FOREIGN KEY).
3) Таблица ""Товары"" (Products) содержит информацию о товарах, включая их идентификаторы (id), названия (name) и цены (price).
4) Таблица ""Заказы-Товары"" (OrderProducts) содержит информацию о отношении заказов к товарам (многие ко многим).

- Напишите запросы на вставку, редактирование и удаление пользователей и продуктов. 
- Напишите запрос на сохранение и удаление заказов
- Напишите запрос на выборку пользователей и выборку товаров
- Напишите запрос на выборку заказов по пользователю
- Напишите запрос на выборку статистики по пользователю (общая сумма заказов/средняя цена товара)
- Создайте требуемые индексы для ускорения выборки

### Критерии оценки
- Реализована схема БД - 2 балла;
- БД заполнена тестовыми данными - 2 балла;
- Реализованы DML и DQL запросы - 2 балла;

#### Зачёт от 6 баллов


### Создание схемы базы данных

```sql
CREATE TABLE Users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) NOT NULL,
                       password VARCHAR(100) NOT NULL
);

CREATE TABLE Orders (
                        id SERIAL PRIMARY KEY,
                        user_id INTEGER NOT NULL,
                        order_date TIMESTAMP NOT NULL,
                        total_amount DECIMAL(10, 2) NOT NULL,
                        FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE Products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE OrderProducts (
                               order_id INTEGER NOT NULL,
                               product_id INTEGER NOT NULL,
                               quantity INTEGER NOT NULL,
                               PRIMARY KEY (order_id, product_id),
                               FOREIGN KEY (order_id) REFERENCES Orders(id),
                               FOREIGN KEY (product_id) REFERENCES Products(id)
);
```

### Запросы на вставку, редактирование и удаление пользователей и продуктов

#### Вставка пользователя

```sql
INSERT INTO Users (name, email, password)
VALUES ('John', 'john@example.com', 'password123');
```

#### Редактирование пользователя

```sql
UPDATE Users
SET name = 'John Smith', email = 'john.smith@example.com'
WHERE id = 1;
```

#### Удаление пользователя

```sql
DELETE FROM Users
WHERE id = 1;
```

#### Вставка продукта

```sql
INSERT INTO Products (name, price)
VALUES ('Product A', 29.99);
```

#### Редактирование продукта

```sql
UPDATE Products
SET name = 'Product B', price = 39.99
WHERE id = 1;
```

#### Удаление продукта

```sql
DELETE FROM Products
WHERE id = 1;
```

### Запрос на сохранение и удаление заказов

#### Сохранение заказа

```sql
INSERT INTO Orders (user_id, order_date, total_amount)
VALUES (2, '2023-10-01 10:00:00', 59.98);

INSERT INTO OrderProducts (order_id, product_id, quantity)
VALUES (2, 2, 2);
```

#### Удаление заказа

```sql
DELETE FROM OrderProducts
WHERE order_id = 1;

DELETE FROM Orders
WHERE id = 1;
```

### Запрос на выборку пользователей и выборку товаров

#### Выборка всех пользователей

```sql
SELECT * FROM Users;
```

#### Выборка всех товаров

```sql
SELECT * FROM Products;
```

### Запрос на выборку заказов по пользователю

```sql
SELECT Orders.id, Orders.order_date, Orders.total_amount
FROM Orders
WHERE Orders.user_id = 2;
```

### Запрос на выборку статистики по пользователю

```sql
SELECT
    Users.id,
    Users.name,
    SUM(Orders.total_amount) AS total_spent,
    AVG(Products.price) AS average_product_price
FROM
    Users
        JOIN
    Orders ON Users.id = Orders.user_id
        JOIN
    OrderProducts ON Orders.id = OrderProducts.order_id
        JOIN
    Products ON OrderProducts.product_id = Products.id
GROUP BY
    Users.id, Users.name;
```

### Создание индексов для ускорения выборки

```sql
CREATE INDEX idx_users_email ON Users(email);
CREATE INDEX idx_orders_user_id ON Orders(user_id);
CREATE INDEX idx_orderproducts_order_id ON OrderProducts(order_id);
CREATE INDEX idx_orderproducts_product_id ON OrderProducts(product_id);
```


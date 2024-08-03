package db

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type DB struct {
	conn *sql.DB
}

func NewDB(connStr string) (*DB, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{conn: conn}, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) InsertUser(name, email, password string) error {
	_, err := db.conn.Exec("INSERT INTO Users (name, email, password) VALUES ($1, $2, $3)", name, email, password)
	return err
}

func (db *DB) UpdateUser(id int, name, email string) error {
	_, err := db.conn.Exec("UPDATE Users SET name = $1, email = $2 WHERE id = $3", name, email, id)
	return err
}

func (db *DB) DeleteUser(id int) error {
	_, err := db.conn.Exec("DELETE FROM Users WHERE id = $1", id)
	return err
}

func (db *DB) GetUsers() ([]User, error) {
	rows, err := db.conn.Query("SELECT id, name, email, password FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if scanErr := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); scanErr != nil {
			return nil, scanErr
		}
		users = append(users, user)
	}

	return users, nil
}

func (db *DB) InsertProduct(name string, price float64) error {
	_, err := db.conn.Exec("INSERT INTO Products (name, price) VALUES ($1, $2)", name, price)
	return err
}

func (db *DB) UpdateProduct(id int, name string, price float64) error {
	_, err := db.conn.Exec("UPDATE Products SET name = $1, price = $2 WHERE id = $3", name, price, id)
	return err
}

func (db *DB) DeleteProduct(id int) error {
	_, err := db.conn.Exec("DELETE FROM Products WHERE id = $1", id)
	return err
}

func (db *DB) GetProducts() ([]Product, error) {
	rows, err := db.conn.Query("SELECT id, name, price FROM Products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if scanErr := rows.Scan(&product.ID, &product.Name, &product.Price); scanErr != nil {
			return nil, scanErr
		}
		products = append(products, product)
	}

	return products, nil
}

func (db *DB) InsertOrder(userID int, orderDate string, totalAmount float64, orderProducts []OrderProduct) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	var orderID int
	err = tx.QueryRow("INSERT INTO Orders (user_id, order_date, total_amount) VALUES ($1, $2, $3) RETURNING id",
		userID, orderDate, totalAmount).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, op := range orderProducts {
		_, execErr := tx.Exec("INSERT INTO OrderProducts (order_id, product_id, quantity) VALUES ($1, $2, $3)",
			orderID, op.ProductID, op.Quantity)
		if execErr != nil {
			tx.Rollback()
			return execErr
		}
	}

	return tx.Commit()
}

func (db *DB) DeleteOrder(orderID int) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	_, execErr := tx.Exec("DELETE FROM OrderProducts WHERE order_id = $1", orderID)
	if execErr != nil {
		tx.Rollback()
		return execErr
	}

	_, execErr = tx.Exec("DELETE FROM Orders WHERE id = $1", orderID)
	if execErr != nil {
		tx.Rollback()
		return execErr
	}
	return tx.Commit()
}

func (db *DB) GetOrdersByUser(userID int) ([]Order, error) {
	rows, err := db.conn.Query("SELECT id, order_date, total_amount FROM Orders WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		if scanErr := rows.Scan(&order.ID, &order.OrderDate, &order.TotalAmount); scanErr != nil {
			return nil, scanErr
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (db *DB) GetUserStatistics(userID int) (UserStatistics, error) {
	var stats UserStatistics
	err := db.conn.QueryRow(`
                SELECT
                        SUM(o.total_amount) AS total_spent,
                        AVG(p.price) AS average_product_price
                FROM
                        Users u
                        JOIN Orders o ON u.id = o.user_id
                        JOIN OrderProducts op ON o.id = op.order_id
                        JOIN Products p ON op.product_id = p.id
                WHERE
                        u.id = $1
                GROUP BY
                        u.id
        `, userID).Scan(&stats.TotalSpent, &stats.AverageProductPrice)
	if err != nil {
		return stats, err
	}

	return stats, nil
}

func (db *DB) CreateIndexes() error {
	_, err := db.conn.Exec(`
                CREATE INDEX IF NOT EXISTS idx_users_email ON Users(email);
                CREATE INDEX IF NOT EXISTS idx_orders_user_id ON Orders(user_id);
                CREATE INDEX IF NOT EXISTS idx_orderproducts_order_id ON OrderProducts(order_id);
                CREATE INDEX IF NOT EXISTS idx_orderproducts_product_id ON OrderProducts(product_id);
        `)
	return err
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Order struct {
	ID          int
	OrderDate   string
	TotalAmount float64
}

type OrderProduct struct {
	OrderID   int
	ProductID int
	Quantity  int
}

type UserStatistics struct {
	TotalSpent          float64
	AverageProductPrice float64
}

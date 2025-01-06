package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect initializes the database connection and creates necessary tables.
func Connect() {
	dsn := "root:root@tcp(localhost:3307)/pizza_shop"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}

	log.Println("Connected to MySQL successfully!")

	// Create tables
	createTables()
}

// Close gracefully closes the database connection.
func Close() {
	if err := DB.Close(); err != nil {
		log.Printf("Failed to close the database connection: %v", err)
	} else {
		log.Println("Database connection closed.")
	}
}

func createTables() {
	// Create items table
	createItemTableQuery := `
    CREATE TABLE IF NOT EXISTS items (
        ItemID INT AUTO_INCREMENT PRIMARY KEY,
        Name VARCHAR(100) NOT NULL,
        Price DECIMAL(10, 2) NOT NULL,
        Category VARCHAR(100) NOT NULL
    );`
	if _, err := DB.Exec(createItemTableQuery); err != nil {
		log.Fatalf("Failed to create 'items' table: %v", err)
	}

	// Create bills table
	createBillTableQuery := `
	CREATE TABLE IF NOT EXISTS bills (
        BillID INT AUTO_INCREMENT PRIMARY KEY,
        TotalAmount DECIMAL(10, 2) NOT NULL,
        Date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	if _, err := DB.Exec(createBillTableQuery); err != nil {
		log.Fatalf("Failed to create 'bills' table: %v", err)
	}

	// Create bill_items table
	createBillItemTableQuery := `
	CREATE TABLE IF NOT EXISTS bill_items (
        BillItemID INT AUTO_INCREMENT PRIMARY KEY,
        BillID INT NOT NULL,
        ItemID INT NOT NULL,
        Quantity INT NOT NULL,
        Price DECIMAL(10, 2) NOT NULL,
        FOREIGN KEY (BillID) REFERENCES bills(BillID) ON DELETE CASCADE,
        FOREIGN KEY (ItemID) REFERENCES items(ItemID) ON DELETE CASCADE
    );`
	if _, err := DB.Exec(createBillItemTableQuery); err != nil {
		log.Fatalf("Failed to create 'bill_items' table: %v", err)
	}

	log.Println("Tables created successfully!")
}

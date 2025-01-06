package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

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

	createItemTableQuery := `
    CREATE TABLE IF NOT EXISTS items (
        ItemID INT AUTO_INCREMENT PRIMARY KEY,
        Name VARCHAR(100) NOT NULL,
        Price VARCHAR(100) NOT NULL,
		Category VARCHAR(100) NOT NULL 
    );
	`
	if _, err = DB.Exec(createItemTableQuery); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	createCartTableQuery := `
	CREATE TABLE IF NOT EXISTS carts (
    CartID INT AUTO_INCREMENT PRIMARY KEY,
    ItemID INT,	
    Quantity INT,
    FOREIGN KEY (ItemID) REFERENCES items(ItemID)
	);`

	if _, err = DB.Exec(createCartTableQuery); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	createInvoiceTableQuery := `
	CREATE TABLE IF NOT EXISTS bills (
    BillID INT AUTO_INCREMENT PRIMARY KEY,
    TotalAmount DECIMAL(10, 2),
	Date DATE
	);
	`

	if _, err = DB.Exec(createInvoiceTableQuery); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

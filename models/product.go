package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productCode"`
	Name        string `json:"name"`
	Inventory   int    `json:"inventory"`
	Price       int    `json:"price"`
	Status      string `json:"status"`
}

func GetProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var p Product

		if err := rows.Scan(&p.ID, &p.ProductCode, &p.Name, &p.Inventory, &p.Price, &p.Status); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, err
}

func (p *Product) GetProduct(db *sql.DB) error {
	return db.QueryRow("SELECT productCode, name, inventory, price, status from products WHERE id = ?", p.ID).Scan(&p.ProductCode, &p.Name, &p.Inventory, &p.Price, &p.Status)
}

func (p *Product) Save(db *sql.DB) error {
	statement, err := db.Prepare("INSERT INTO products (productCode, name, inventory, price, status) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer statement.Close()

	newProduct, err := statement.Exec(p.ProductCode, p.Name, p.Inventory, p.Price, p.Status)

	if err != nil {
		return err
	}

	insertedId, err := newProduct.LastInsertId()

	if err != nil {
		return err
	}

	p.ID = int(insertedId)

	return nil
}

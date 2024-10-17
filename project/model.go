package main

import (
	"database/sql"
	"log"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Get all products
func getProducts(db *sql.DB) ([]product, error) {
	query := "SELECT id, name, quantity, price FROM products"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var products []product
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error in row iteration: %v", err)
		return nil, err
	}

	return products, nil
}

// Get a single product by ID
func (p *product) getProduct(db *sql.DB) error {
	query := "SELECT name, quantity, price FROM products WHERE id = ?"
	row := db.QueryRow(query, p.ID)

	err := row.Scan(&p.Name, &p.Quantity, &p.Price)
	if err != nil {
		log.Printf("Error fetching product with ID %d: %v", p.ID, err)
		return err
	}

	return nil
}

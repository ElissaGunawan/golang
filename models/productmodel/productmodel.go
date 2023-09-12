package productmodel

import (
	"go-web/config"
	"go-web/entities"
)

func Getall() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT 
			products.id, 
			products.name,
			categories.id as category_id,
			categories.name as category_name,
			products.brand, 
			products.price,
			products.status FROM products
		JOIN categories ON products.category_id = categories.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Id,
			&product.Category.Name,
			&product.Brand,
			&product.Price,
			&product.Status,
		); err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`
		INSERT INTO products(
			name, category_id, brand, price, status
		) VALUES (?, ?, ?, ?, ?)`,
		product.Name,
		product.Category.Id,
		product.Brand,
		product.Price,
		product.Status,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
		SELECT 
			products.id, 
			products.name, 
			categories.name as category_name,
			products.brand, 
			products.price,
			products.status FROM products
		JOIN categories ON products.category_id = categories.id
		WHERE products.id = ?
	`, id)

	var product entities.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
		&product.Brand,
		&product.Price,
		&product.Status,
	)

	if err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(`
		UPDATE products SET 
			name = ?, 
			category_id = ?,
			brand = ?,
			price = ?,
			status = ?
		WHERE id = ?`,
		product.Name,
		product.Category.Id,
		product.Brand,
		product.Price,
		product.Status,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

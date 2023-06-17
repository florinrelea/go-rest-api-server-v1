# Create test table and insert examples
```
_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		productCode TEXT,
		name TEXT,
		inventory INTEGER,
		price INTEGER,
		status TEXT
		);`)

if err != nil {
  log.Fatal(err.Error())
}

_, err = DB.Exec(`INSERT INTO products (productCode, name, inventory, price, status) VALUES
  (1, "Product 1", 10, 100, "active"),
  (2, "Product 2", 20, 200, "active"),
  (3, "Product 3", 30, 300, "active"),
  (4, "Product 4", 40, 400, "active"),
  (5, "Product 5", 50, 500, "active"),
  (6, "Product 6", 60, 600, "active"),
  (7, "Product 7", 70, 700, "active"),
  (8, "Product 8", 80, 800, "active"),
  (9, "Product 9", 90, 900, "active"),
  (10, "Product 10", 100, 1000, "active")
  ;`)
```
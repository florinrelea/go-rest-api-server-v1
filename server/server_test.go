package server_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"example.com/server"
)

const tableProductCreationQuery = `CREATE TABLE IF NOT EXISTS products (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	productCode VARCHAR(25) NOT NULL,
	name VARCHAR(256) NOT NULL,
	price INT NOT NULL,
	inventory INT NOT NULL,
	status VARCHAR(64) NOT NULL
)`

var serverApp server.ServerApp

func TestMain(m *testing.M) {
	serverApp = server.ServerApp{}

	serverApp.Init(":9003")

	ensureTableExists()

	code := m.Run()

	clearProductTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := serverApp.DB.Exec(tableProductCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearProductTable() {
	serverApp.DB.Exec("DELETE FROM products")
	serverApp.DB.Exec("DELETE FROM sqlite_sequence WHERE name='products'")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	serverApp.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistingProduct(t *testing.T) {
	clearProductTable()

	req, _ := http.NewRequest("GET", "/products/11", nil)

	response := executeRequest(req)

	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	var m map[string]string

	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "sql: no rows in result set" {
		t.Errorf("Expected the 'error' key of the response to be set to 'sql: no rows in result set'. Got '%s'", m["error"])
	}
}

func TestCreateProduct(t *testing.T) {
	clearProductTable()

	payload := []byte(`{"productCode":"12345","name":"ProductTest","price":1,"inventory":1,"status":"testing"}`)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(payload))

	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["productCode"] != "12345" {
		t.Errorf("Expected productCode to be '12345'. Got '%v'", m["productCode"])
	}

	if m["name"] != "ProductTest" {
		t.Errorf("Expected product name to be 'ProductTest'. Got '%v'", m["name"])
	}

	if m["price"] != 1.0 {
		t.Errorf("Expected product price to be '1'. Got '%v'", m["price"])
	}

	if m["inventory"] != 1.0 {
		t.Errorf("Expected product inventory to be '1'. Got '%v'", m["inventory"])
	}

	if m["status"] != "testing" {
		t.Errorf("Expected product status to be 'testing'. Got '%v'", m["status"])
	}
}

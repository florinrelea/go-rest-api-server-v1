package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example.com/models"
	"github.com/gorilla/mux"
)

type ServerApp struct {
	Port   string
	DB     *sql.DB
	Router *mux.Router
}

// Helper functions
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Helper functions >> finish

func (a *ServerApp) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product id.")
		return
	}

	var p models.Product

	p.ID = id

	err = p.GetProduct(a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
	return
}

func (a *ServerApp) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts(a.DB)

	if err != nil {
		fmt.Printf("getProducts error %s\n", err.Error())

		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *ServerApp) createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Todo: data validation
	err = newProduct.Save(a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, newProduct)
}

func (a *ServerApp) Init(port string) {
	rtr := mux.NewRouter()
	DB, err := sql.Open("sqlite3", "./practice.db")

	if err != nil {
		log.Fatal(err.Error())
	}

	a.Router = rtr
	a.DB = DB
	a.Port = port

	a.Router.HandleFunc("/products", a.getAllProducts).Methods("GET")
	a.Router.HandleFunc("/products/{id}", a.getProduct).Methods("GET")
	a.Router.HandleFunc("/products", a.createProduct).Methods("POST")
}

func (a *ServerApp) Run() {
	fmt.Println("Server started on port", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}

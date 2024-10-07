package main

import (
	"database/sql"
	"github.com/gorilla/mux"

	- "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialise() error {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v",DbUser, DBPassword, DBName)
	var err error 
	app.DB, err = sql.Open("mysql", connectionString)
	if err!=nil{
		return err
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	return nil
}

func (app *App) Run (address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func sendResponse(w http.ResponseWriter,statusCode int, payload interface{}){
	response , _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.write(response)	
}

func sendError(w http.ResponseWriter, statusCode int , err string){
	error_message := map[string]string{"error":err}
	sendResponse(w, statusCode , error_message)
}

func (app * App) handleRoutes(){
	app.Router.HandleFunc("/api/products", app.getProducts).Methods("GET")
    app.Router.HandleFunc("/api/products/{id}", app.getProduct).Methods("GET")
    app.Router.HandleFunc("/api/products", app.createProduct).Methods("POST")
    app.Router.HandleFunc("/api/products/{id}", app.updateProduct).Methods("PUT")
    app.Router.HandleFunc("/api/products/{id}", app.deleteProduct).Methods("DELETE")
}
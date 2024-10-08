package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestMain(m *testing.M) {

	err := a.Initialise(DbUser , DBPassword, "test")
	if err != nil {
		log.Fatal("Erro occured while initialising the database")
	}
	createTable()
	m.Run()
}

func createTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS products (
        id int NOT NULL AUTO_INCREMENT ,
        name VARCHAR(255) NOT NULL,
        quantity int,
        price float(10,7),
		PRIMARY KEY (id)
    );`
    _, err := a.DB.Exec(createTableQuery)
    if err != nil {
        log.Fatal(err)
    }
}

func clearTable(){
	a.DB.Exec("DELETE from products")
	a.DB.Exec("ALTER table products AUTO_INCREMENT=1")
	log.println("clearTable")
}

func addProduct (name string, quantity int, price float64){
	query := fmt.Sprintf("INSERT into products(name, quantity, price) VALUES('%v',%v,%v)", name, quantity, price)
	_, err := a.DB.Exec(query)
	if err != nil {
		log.println(err)
	}
}

func TestGetProduct(t *testing.T) {	
	clearTable()
	addProduct("Keyboard",100,5000)
	request , _ := http.NewRequest("GET", "/products/1",nil)
	sendRequest()
}


func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errof("Expected status: %v, Received: %v", expectedStatusCode, actualStatusCode)
	}
}

func sendRequest(request *http.Request) *http.ResponseRecorder{
	recorder := httptest.NewRecorder()
	a.Router.ServerHTTP(recorder,Request)
	return recorder
}


func TestCrzateProduct(t *testing.T){
	clearTable()
	var product = []byte(`{"name":"chair", "quantity":1,"price":100}`)
	req, _ := http.NewRequest("POST","/product", bytes.NewBuffer(product))
	req.Header.Set("Content-Type","application/json")

	response := sendRequest(req)
	checkStatusCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes() &m)

	if m["name"] != "chair" {
		t.Errorf("Expected name: %v, Got: %v", "chair", m["name"])
	}
	log.Printf("%T", m["quantity"])
	if m["quantity"] != 1.0 {
		t.Errorf("Expected quantity: %v, Got: %v", 1.0, m["quantity"])
	}
}
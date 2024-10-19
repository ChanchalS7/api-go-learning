package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App 

func TestMain(m *testing.M){
	err := a.Initialise(DBUser,DBPassword,"test_go")
	

	if err !=nil{
		log.Fatal("Error occurred while initialising the database")
	}
	createTable()
	m.Run()//run all other test withing package

}

func createTable(){
	createTableQuery:=`CREATE TABLE IF NOT EXISTS products(
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	quantity INT,
	price FLOAT(10,7),
	PRIMARY KEY(id)
	);`
	_,err:=a.DB.Exec(createTableQuery)
	if err!=nil{
		log.Fatal(err)
	}
}

func clearTable(){
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER TABLE products AUTO_INCREMENT=1")
	log.Println("Clear table")
}
func addProduct(name string, quantity int, price float64){
	// Prepare the insert query using placeholders
	query := "INSERT INTO products (name, quantity, price) VALUES (?, ?, ?)"

	_,err:= a.DB.Exec(query,name,quantity,price)
	if err!=nil{
		log.Println(err)
	}
}
func TestGetProduct(t *testing.T){
clearTable()
addProduct("keyboard",100,500)
request,_:=http.NewRequest("GET","/product/1",nil)
response :=sendRequest(request)
checkStatusCode(t,http.StatusOK,response.Code)



}
func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int){
 if expectedStatusCode != actualStatusCode{
	t.Errorf("Expected status : %v, Received:%v", expectedStatusCode,actualStatusCode)
 }
}
func sendRequest(request *http.Request) *httptest.ResponseRecorder{
	recorder:=httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder
}


func TestCreateProduct(t *testing.T) {
    clearTable()
    var product = []byte(`{"name":"chair","quantity":1,"price":100}`)
    req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))
    req.Header.Set("Content-Type", "application/json")

    response := sendRequest(req)
    checkStatusCode(t, http.StatusCreated, response.Code)

    // if response.Code != http.StatusCreated {
    //     t.Errorf("Expected status code 201 but got %v. Response body: %s", response.Code, response.Body.String())
    // }
}

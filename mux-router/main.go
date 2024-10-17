package main

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)
//1.Create web server
//2. HTTP Verbs
/*
-GET
-POST
-PUT
-PATCH
-DELETE
*/
//3.REST ARCHITECTURE
/*
-Client-server architecture
- stateless communication
- uniform interface
- 
*/
//4. Project explanation
/*
- product api
-[id,name,quantity,price]
-database
- GET ->/products
- GET ->/products/id
- POST ->/product
- PUT ->/product/id
- DELETE ->/product/id

*/

/*
***MUX Router***
-create slice to store data in memory storage
- Create Product struct
- Create product slice
- Create some product
- Create handleRequests() method which have all the routes
- Write returnAllProducts() method which return all products
- Write getProduct() method which return single product
- Convert the id type from int to string
- test get-all and get single product
- Now use gorilla/mux router library


*/
type Product struct{
	Id 			string
	Name 		string 
	Quantity 	int 
	Price 		float64
}
//declare as a global variable
var Products []Product

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to homepage")
	fmt.Println("Endpoint hit:homepage")
}
//returnAllProducts() method) 
func returnAllProducts( w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint hit: returnAllProducts")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(Products)
}
func getProduct(w http.ResponseWriter, r *http.Request){
// fmt.Println(r.URL.Path)
// key:=r.URL.Path[len("/product/"):]
// for _,product:=range Products{
// 	if (product.Id)==key{
// 		json.NewEncoder(w).Encode(product)
// 	}
// }
vars:=mux.Vars(r)
key:=vars["id"]
for _,product:=range Products{
	if (product.Id)==key{
		json.NewEncoder(w).Encode(product)
	}

}

}
func handleRequests(){
	myRouter:=mux.NewRouter().StrictSlash(true)
	// Log a message that the server is starting
    log.Println("Server is starting on http://localhost:10000")
	myRouter.HandleFunc("/products",returnAllProducts)
	myRouter.HandleFunc("/product/{id}",getProduct)
	myRouter.HandleFunc("/",homepage)
	log.Fatal(http.ListenAndServe("localhost:10000",myRouter))

}
func main(){
	//product slice
	Products = []Product{
		Product{Id:"1", Name:"Chair", Quantity:100, Price:100.00},
		Product{Id:"2", Name:"Desk", Quantity:200,Price:200.00},
	}
	//call handleRequest method
	handleRequests()
	// http.HandleFunc("/",homepage)
	// http.ListenAndServe("localhost:10000",nil)
}
package main

import (
	"net/http"
	"fmt"
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

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to homepage")
	fmt.Println("Endpoint hit:homepage")
}
func main(){
	http.HandleFunc("/",homepage)
	http.ListenAndServe("localhost:10000",nil)
}
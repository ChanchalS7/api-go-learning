package main

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)
type App struct{
	Router *mux.Router
	DB *sql.DB
}

func (app *App) Initialise()error{
// Create the connection string
connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPassword, DBNAME)
var err error
//open connection 
app.DB,err=sql.Open("mysql",connectionString)

if err!=nil{
	return err
}
app.Router=mux.NewRouter().StrictSlash(true)
return nil
}

func (app *App) Run(address string){
	log.Fatal(http.ListenAndServe(address,app.Router))
}
func (app *App) handleRoutes(){
	app.Router.HandleFunc("/products",getProducts).Methods("GET")
}
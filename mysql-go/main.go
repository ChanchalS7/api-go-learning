package main
import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)

func checkError(e error){
	if e!=nil{
		log.Fatal(e)
	}
}
type Data struct{
	id int 
	name string
}
func main(){
	connectionString:=fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v",DBUser,DBPassword,DBNAME)
	db,err:=sql.Open("mysql",connectionString)
	checkError(err)
	defer db.Close()

	rows,err:=db.Query("SELECT * FROM data")
	checkError(err)
	for rows.Next(){
		var data Data
		err:=rows.Scan(&data.id,&data.name)
		checkError(err)
		fmt.Println(data)
	}

}
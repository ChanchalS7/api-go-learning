package main
import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)


// Data struct to hold the retrieved data
type Data struct {
	ID   int    // Exported field for ID
	Name string // Exported field for Name
}

// checkError checks for errors and logs them
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Create the connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", DBUser, DBPassword, DBNAME)

	// Open a connection to the database
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close() // Ensure the database connection is closed

	// Insert a new record
	result, err := db.Exec("INSERT INTO data (id, name) VALUES (?, ?)", 4, "xyz") // Use placeholders
	checkError(err)

	// Get the last inserted ID
	lastInsertedId, err := result.LastInsertId()
	checkError(err)
	fmt.Println("Last Inserted ID:", lastInsertedId)

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	checkError(err)
	fmt.Println("Rows Affected:", rowsAffected)

	// Query the database for all records
	rows, err := db.Query("SELECT id, name FROM data") // Ensure the query matches your table's structure
	checkError(err)
	defer rows.Close() // Ensure rows are closed after usage

	// Iterate through the rows and print the data
	for rows.Next() {
		var data Data
		err := rows.Scan(&data.ID, &data.Name) // Use the exported fields
		checkError(err)
		fmt.Println(data)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
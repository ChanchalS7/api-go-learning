package main

import "fmt"

func main() {
	app := App{}
	app.Initialise()
	app.Run("localhost:10000")
	fmt.Println("Server is running")
}
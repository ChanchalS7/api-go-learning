# Inventory Management System - REST API

This project is an example of a simple inventory management system built with Go, MySQL, and Gorilla Mux for routing. It includes endpoints for CRUD operations on product data.

## Prerequisites

- Go (1.19+)
- MySQL Server
- Postman or curl for API testing

## Project Structure

```bash
.
├── app.go            # Main application logic
├── app_test.go       # Unit tests
├── constants.go      # Database credentials
├── main.go           # Entry point for the server
├── model.go          # Database models for Product
└── README.md         # Project documentation


```

# API Go Learning

This project is a simple API built with Go, utilizing the Gorilla Mux router and MySQL for inventory management.

## Setup

### Clone the repository:
```bash
git clone https://github.com/ChanchalS7/api-go-learning.git
```

### Install dependencies:
1. Install the Gorilla Mux router package:
   ```bash
   go get -u github.com/gorilla/mux
   ```
   
2. Install the MySQL driver package:
   ```bash
   go get -u github.com/go-sql-driver/mysql
   ```

### Configure database:
Update the `constants.go` file with your MySQL database credentials:
```go
const DBUser = "root"
const DBPassword = "your_password"
const DBNAME = "inventory"
```

### Create the database:
Log into your MySQL instance and create the database named `inventory`.

## Run the application:
Start the server by running:
```bash
go run main.go
```
The server will run at [http://localhost:10000](http://localhost:10000).

## API Endpoints

### 1. Get All Products
- **Endpoint:** `/products`
- **Method:** GET
- **Description:** Fetch all products from the inventory.

**Example Request (using curl):**
```bash
curl -X GET http://localhost:10000/products
```

**Example Response:**
```json
[
  {
    "id": 1,
    "name": "keyboard",
    "quantity": 100,
    "price": 500.0
  }
]
```

### 2. Get a Product by ID
- **Endpoint:** `/product/{id}`
- **Method:** GET
- **Description:** Fetch a product by its ID.

**Example Request (using Postman or curl):**
```bash
curl -X GET http://localhost:10000/product/1
```

**Example Response:**
```json
{
  "id": 1,
  "name": "keyboard",
  "quantity": 100,
  "price": 500.0
}
```

### 3. Create a Product
- **Endpoint:** `/product`
- **Method:** POST
- **Description:** Add a new product to the inventory.

**Example Request:**
```bash
curl -X POST http://localhost:10000/product \
-H "Content-Type: application/json" \
-d '{"name":"chair","quantity":10,"price":100.0}'
```

**Example Response:**
```json
{
  "id": 2,
  "name": "chair",
  "quantity": 10,
  "price": 100.0
}
```

### 4. Update a Product
- **Endpoint:** `/product/{id}`
- **Method:** PUT
- **Description:** Update an existing product in the inventory.

**Example Request:**
```bash
curl -X PUT http://localhost:10000/product/1 \
-H "Content-Type: application/json" \
-d '{"name":"updated keyboard","quantity":50,"price":400.0}'
```

**Example Response:**
```json
{
  "id": 1,
  "name": "updated keyboard",
  "quantity": 50,
  "price": 400.0
}
```

### 5. Delete a Product
- **Endpoint:** `/product/{id}`
- **Method:** DELETE
- **Description:** Remove a product from the inventory by its ID.

**Example Request:**
```bash
curl -X DELETE http://localhost:10000/product/1
```

**Example Response:**
```json
{
  "result": "Successful deletion"
}
```


## Running Test Cases

Testing is crucial for ensuring that your API behaves as expected. This section outlines how to run existing test cases and provides guidance on writing your own tests.

### Prerequisites
Before running tests, ensure that:
- You have Go installed on your machine. You can verify this by running `go version` in your terminal.
- Your project dependencies are up to date. Run `go mod tidy` to ensure everything is in order.

### Running Tests
To execute the test cases included in the project, follow these steps:

1. **Navigate to the Project Directory:**
   Open your terminal and change to the directory of your project:
   ```bash
   cd api-go-learning
   ```

2. **Run the Tests:**
   Use the following command to run all tests in the project:
   ```bash
   go test ./...
   ```
   This command tells Go to look for tests in the current directory and all subdirectories.

3. **Example Output:**
   After running the tests, you should see output similar to this:
   ```
   ?   	your_module_name	[no test files]
   === RUN   TestGetAllProducts
   === RUN   TestGetProductByID
   === RUN   TestCreateProduct
   === RUN   TestUpdateProduct
   === RUN   TestDeleteProduct
   --- PASS: TestGetAllProducts (0.00s)
   --- PASS: TestGetProductByID (0.00s)
   --- PASS: TestCreateProduct (0.00s)
   --- PASS: TestUpdateProduct (0.00s)
   --- PASS: TestDeleteProduct (0.00s)
   PASS
   ok  	your_module_name	0.002s
   ```

### Writing Your Own Tests
You can extend your test suite by creating custom tests. Here are the steps to write effective tests:

1. **Create a Test File:**
   Add a new file named `*_test.go` in the same package as the code you want to test. For example, if your main file is `main.go`, create a file called `main_test.go`.

2. **Import the Necessary Packages:**
   At the top of your test file, import the required packages:
   ```go
   package main

   import (
       "net/http"
       "net/http/httptest"
       "testing"
   )
   ```

3. **Write a Test Function:**
   Define your test function using the `Test` prefix followed by the name of the function you are testing. Here's an example for testing the `GetAllProducts` handler:
   ```go
   func TestGetAllProducts(t *testing.T) {
       req, err := http.NewRequest("GET", "/products", nil)
       if err != nil {
           t.Fatal(err)
       }
       
       rr := httptest.NewRecorder()
       handler := http.HandlerFunc(GetAllProducts) // Replace with your actual handler
       handler.ServeHTTP(rr, req)

       // Check the status code
       if status := rr.Code; status != http.StatusOK {
           t.Errorf("handler returned wrong status code: got %v want %v",
               status, http.StatusOK)
       }

       // Check the response body (optional)
       expected := `[{"id":1,"name":"keyboard","quantity":100,"price":500.0}]`
       if rr.Body.String() != expected {
           t.Errorf("handler returned unexpected body: got %v want %v",
               rr.Body.String(), expected)
       }
   }
   ```

4. **Run Your Tests:**
   After writing your tests, run them using the command mentioned earlier:
   ```bash
   go test ./...
   ```

### Testing Coverage
To ensure your tests cover a significant portion of your code, use the coverage tool:
```bash
go test -cover ./...
```
This command will display the percentage of code covered by your tests. A higher percentage indicates better test coverage, although it’s important to focus on the quality of tests rather than just the coverage percentage.

### Best Practices for Writing Tests
- **Isolate Tests:** Each test should be independent. Avoid sharing state between tests to ensure they can be run in any order.
- **Use Descriptive Names:** Name your test functions clearly to indicate what functionality they are testing.
- **Test Edge Cases:** Ensure that you cover both expected and unexpected inputs to validate how your application behaves under different scenarios.
- **Keep Tests Simple:** Write tests that are easy to read and understand. Complex tests can lead to confusion and maintenance challenges.

### Example Test Cases
Here’s an additional example of testing the product creation endpoint:
```go
func TestCreateProduct(t *testing.T) {
    req, err := http.NewRequest("POST", "/product", strings.NewReader(`{"name":"desk","quantity":5,"price":250.0}`))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreateProduct) // Replace with your actual handler
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }

    // You may add additional checks for the response body
}
```

### Conclusion
Running tests regularly and writing comprehensive tests are crucial practices for maintaining the quality of your API. They help you catch issues early and ensure that your application behaves as expected, leading to a more robust and reliable codebase.


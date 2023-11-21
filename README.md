# Bank RestAPI using GoLang and MySQL

## Project Overview

This project is a simple Bank RestAPI built using GoLang and MySQL. It provides basic functionalities for managing bank accounts, transactions, and user information. The project uses the Gorilla Mux framework for routing and the go-sql-driver/mysql package for MySQL database connectivity.

## Tools Used

### Software

- **Golang Version:** 1.21.0
- **Thunder Client:** Used for testing and interacting with the RestAPI.

### Databases

- **MySQL**
- **Golang MySQL driver:** github.com/go-sql-driver/mysql v1.7.1

### Frameworks

- **Gorilla Mux Version:** github.com/gorilla/mux v1.8.1
  - Used for routing in the Golang application.

### Libraries

- **Golang Standard Library**
  - Utilized for various functionalities within Golang.

## Project Structure

The project is organized into several modules, each serving a specific purpose.

### Packages

1. **config**
   - `app.go`: Configuration settings for the application.

2. **controllers**
   - `Account_controller.go`: Handles account-related operations.
   - `Transaction_controller.go`: Manages transaction-related functionalities.
   - `Users_controller.go`: Controls user-related actions.

3. **modules**
   - `Account_modules.go`: Implements account-related logic.
   - `Transaction_modules.go`: Provides logic for transaction processing.
   - `Users_modules.go`: Manages user-related operations.

4. **routes**
   - `Bank_routes.go`: Defines the routes and their corresponding handlers.

5. **utils**
   - `utilities.go`: Contains utility functions used across the project.

## Running the Project

1. Ensure you have Golang and MySQL installed.
2. Set up the MySQL database using the `DB.sql` file.
3. Modify the `config/app.go` file to configure the database connection.
4. Open a terminal and navigate to the project directory.
5. Run the following command to start the server:

   ```bash
   go run main.go

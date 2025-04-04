# Singleton & Dependency Injection

## Problem
Implement a Singleton pattern for a database connection pool and use Dependency Injection to provide this connection pool to various services. The system should ensure that only one instance of the connection pool is created, and services should be testable by allowing mock implementations to be injected.

## Requirements
1. Create a DatabaseConnectionPool as a Singleton
2. Implement connection management methods (GetConnection, ReleaseConnection)
3. Create service interfaces for UserService and ProductService
4. Implement concrete services that depend on the connection pool
5. Use Dependency Injection to provide the connection pool to services
6. Write tests using mock implementations

## Examples
```go
// Get the singleton instance of the connection pool
connectionPool := GetDatabaseConnectionPool()

// Create services with dependency injection
userService := NewUserService(connectionPool)
productService := NewProductService(connectionPool)

// Use the services
user, err := userService.GetUserByID("user123")
product, err := productService.GetProductByID("product456")
Expected Output
CopyInitializing database connection pool...
Opening connection to database
Getting user with ID: user123
Getting product with ID: product456
Closing connection to database
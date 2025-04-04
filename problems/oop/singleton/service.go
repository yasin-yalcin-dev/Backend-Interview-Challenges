/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

// problems/oop/singleton/service.go
package singleton

import (
	"fmt"
)

// User represents a user in the system
type User struct {
	ID       string
	Username string
	Email    string
}

// Product represents a product in the system
type Product struct {
	ID    string
	Name  string
	Price float64
}

// UserService defines methods for user management
type UserService interface {
	GetUserByID(id string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

// ProductService defines methods for product management
type ProductService interface {
	GetProductByID(id string) (*Product, error)
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id string) error
}

// UserServiceImpl implements UserService with dependency injection
type UserServiceImpl struct {
	connectionPool ConnectionPool
}

// NewUserService creates a new user service with the given connection pool
func NewUserService(connectionPool ConnectionPool) UserService {
	return &UserServiceImpl{
		connectionPool: connectionPool,
	}
}

// GetUserByID retrieves a user by ID
func (s *UserServiceImpl) GetUserByID(id string) (*User, error) {
	conn, err := s.connectionPool.GetConnection()
	if err != nil {
		return nil, err
	}
	defer s.connectionPool.ReleaseConnection(conn)

	fmt.Printf("Getting user with ID: %s\n", id)

	// In a real implementation, we would query the database here
	// For this example, we'll just return a mock user
	return &User{
		ID:       id,
		Username: "user_" + id,
		Email:    "user_" + id + "@example.com",
	}, nil
}

// CreateUser creates a new user
func (s *UserServiceImpl) CreateUser(user *User) error {
	conn, err := s.connectionPool.GetConnection()
	if err != nil {
		return err
	}
	defer s.connectionPool.ReleaseConnection(conn)

	fmt.Printf("Creating user: %s\n", user.Username)

	// In a real implementation, we would insert into the database here
	return nil
}

// UpdateUser updates an existing user
func (s *UserServiceImpl) UpdateUser(user *User) error {
	conn, err := s.connectionPool.GetConnection()
	if err != nil {
		return err
	}
	defer s.connectionPool.ReleaseConnection(conn)

	fmt.Printf("Updating user with ID: %s\n", user.ID)

	// In a real implementation, we would update the database here
	return nil
}

// DeleteUser deletes a user by ID
func (s *UserServiceImpl) DeleteUser(id string) error {
	conn, err := s.connectionPool.GetConnection()
	if err != nil {
		return err
	}
	defer s.connectionPool.ReleaseConnection(conn)

	fmt.Printf("Deleting user with ID: %s\n", id)

	// In a real implementation, we would delete from the database here
	return nil
}

// ProductServiceImpl implements ProductService with dependency injection
type ProductServiceImpl struct {
	connectionPool ConnectionPool
}

// NewProductService creates a new product service with the given connection pool
func NewProductService(connectionPool ConnectionPool) ProductService {
	return &ProductServiceImpl{
		connectionPool: connectionPool,
	}
}

// GetProductByID retrieves a product by ID
func (s *ProductServiceImpl) GetProductByID(id string) (*Product, error) {
	conn, err := s.connectionPool.GetConnection()
	if err != nil {
		return nil, err
	}
	defer s.connectionPool.ReleaseConnection(conn)

	fmt.Printf("Getting product with ID: %s\n", id)

	// In a real implementation, we would query the database here
	// For this example, we'll just return a mock product
	return &Product{
		ID:    id,
		Name:  "Product " + id,
		Price: 99.99,
	}, nil
}

// CreateProduct creates a new product
func (s *ProductServiceImpl) CreateProduct(product *Product) error {
	conn, err := s.connectionPool.GetConnection()
	if err != nil {
		return err
	}
	defer s.connectionPool.ReleaseConnection(conn)

	fmt.Printf("Creating product: %s\n", product.Name)

	// In a real implementation, we would insert into the database here
	return nil
}

// UpdateProduct updates an existing product
func (s *ProductServiceImpl) UpdateProduct(product *Product) error {
	conn, err := s.connectionPool.GetConnection()
	if err != nil {
		return err
	}
	defer s.connectionPool.ReleaseConnection(conn)

	fmt.Printf("Updating product with ID: %s\n", product.ID)

	// In a real implementation, we would update the database here
	return nil
}

// DeleteProduct deletes a product by ID
func (s *ProductServiceImpl) DeleteProduct(id string) error {
	conn, err := s.connectionPool.GetConnection()
	if err != nil {
		return err
	}
	defer s.connectionPool.ReleaseConnection(conn)

	fmt.Printf("Deleting product with ID: %s\n", id)

	// In a real implementation, we would delete from the database here
	return nil
}

/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

// problems/oop/singleton/singleton_test.go
package singleton

import (
	"testing"
)

// MockConnectionPool implements the ConnectionPool interface for testing
type MockConnectionPool struct {
	getConnectionCalled     int
	releaseConnectionCalled int
}

// GetConnection mock implementation
func (m *MockConnectionPool) GetConnection() (*DatabaseConnection, error) {
	m.getConnectionCalled++
	return &DatabaseConnection{id: 999, isInUse: true}, nil
}

// ReleaseConnection mock implementation
func (m *MockConnectionPool) ReleaseConnection(conn *DatabaseConnection) error {
	m.releaseConnectionCalled++
	return nil
}

// GetStats mock implementation
func (m *MockConnectionPool) GetStats() ConnectionPoolStats {
	return ConnectionPoolStats{
		TotalConnections:  1,
		ActiveConnections: m.getConnectionCalled - m.releaseConnectionCalled,
		IdleConnections:   1 - (m.getConnectionCalled - m.releaseConnectionCalled),
	}
}

// TestSingleton tests that the connection pool is truly a singleton
func TestSingleton(t *testing.T) {
	// Reset the singleton instance before the test
	ResetConnectionPool()

	// Get the singleton instance
	pool1 := GetDatabaseConnectionPool()
	pool2 := GetDatabaseConnectionPool()

	// Both instances should be the same
	if pool1 != pool2 {
		t.Errorf("Expected both instances to be the same")
	}

	// Test connection creation and release
	conn, err := pool1.GetConnection()
	if err != nil {
		t.Errorf("Expected no error getting connection, got: %v", err)
	}

	stats := pool1.GetStats()
	if stats.ActiveConnections != 1 {
		t.Errorf("Expected 1 active connection, got: %d", stats.ActiveConnections)
	}

	err = pool1.ReleaseConnection(conn)
	if err != nil {
		t.Errorf("Expected no error releasing connection, got: %v", err)
	}

	stats = pool1.GetStats()
	if stats.ActiveConnections != 0 {
		t.Errorf("Expected 0 active connections after release, got: %d", stats.ActiveConnections)
	}
}

// TestDependencyInjection tests the dependency injection pattern
func TestDependencyInjection(t *testing.T) {
	// Create a mock connection pool
	mockPool := &MockConnectionPool{}

	// Create services using the mock
	userService := NewUserService(mockPool)

	// Use the service
	user, err := userService.GetUserByID("test123")

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if user == nil {
		t.Errorf("Expected user to be returned")
	}

	// Test that the mock methods were called
	if mockPool.getConnectionCalled != 1 {
		t.Errorf("Expected GetConnection to be called once, called %d times", mockPool.getConnectionCalled)
	}

	if mockPool.releaseConnectionCalled != 1 {
		t.Errorf("Expected ReleaseConnection to be called once, called %d times", mockPool.releaseConnectionCalled)
	}

	// Test product service with the same mock
	productService := NewProductService(mockPool)

	// Use the service
	product, err := productService.GetProductByID("prod123")

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if product == nil {
		t.Errorf("Expected product to be returned")
	}

	// Test that the mock methods were called again
	if mockPool.getConnectionCalled != 2 {
		t.Errorf("Expected GetConnection to be called twice, called %d times", mockPool.getConnectionCalled)
	}

	if mockPool.releaseConnectionCalled != 2 {
		t.Errorf("Expected ReleaseConnection to be called twice, called %d times", mockPool.releaseConnectionCalled)
	}
}

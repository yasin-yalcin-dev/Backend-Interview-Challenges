/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package singleton

import (
	"fmt"
	"sync"
)

// DatabaseConnection represents a connection to the database
type DatabaseConnection struct {
	id      int
	isInUse bool
}

// ConnectionPool interface defines methods for a connection pool
type ConnectionPool interface {
	GetConnection() (*DatabaseConnection, error)
	ReleaseConnection(conn *DatabaseConnection) error
	GetStats() ConnectionPoolStats
}

// ConnectionPoolStats contains statistics about the connection pool
type ConnectionPoolStats struct {
	TotalConnections  int
	ActiveConnections int
	IdleConnections   int
}

// DatabaseConnectionPool implements the Singleton pattern for a database connection pool
type DatabaseConnectionPool struct {
	connections       []*DatabaseConnection
	maxConnections    int
	activeConnections int
	mutex             sync.Mutex
}

// singleton instance with thread-safe initialization
var (
	instance *DatabaseConnectionPool
	once     sync.Once
)

// GetDatabaseConnectionPool returns the singleton instance of the connection pool
func GetDatabaseConnectionPool() *DatabaseConnectionPool {
	once.Do(func() {
		fmt.Println("Initializing database connection pool...")
		instance = &DatabaseConnectionPool{
			connections:    make([]*DatabaseConnection, 0),
			maxConnections: 10, // Default max connections
		}
	})
	return instance
}

// GetConnection gets a connection from the pool or creates a new one if needed
func (p *DatabaseConnectionPool) GetConnection() (*DatabaseConnection, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// First try to find an idle connection
	for _, conn := range p.connections {
		if !conn.isInUse {
			conn.isInUse = true
			p.activeConnections++
			fmt.Println("Opening connection to database")
			return conn, nil
		}
	}

	// If no idle connection is found and we haven't reached max, create a new one
	if len(p.connections) < p.maxConnections {
		conn := &DatabaseConnection{
			id:      len(p.connections) + 1,
			isInUse: true,
		}
		p.connections = append(p.connections, conn)
		p.activeConnections++
		fmt.Println("Opening connection to database")
		return conn, nil
	}

	// If we've reached max connections, return an error
	return nil, fmt.Errorf("connection pool exhausted, max %d connections reached", p.maxConnections)
}

// ReleaseConnection returns a connection to the pool
func (p *DatabaseConnectionPool) ReleaseConnection(conn *DatabaseConnection) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	found := false
	for _, c := range p.connections {
		if c == conn {
			found = true
			if !c.isInUse {
				return fmt.Errorf("connection is already released")
			}
			c.isInUse = false
			p.activeConnections--
			fmt.Println("Closing connection to database")
			break
		}
	}

	if !found {
		return fmt.Errorf("connection not found in the pool")
	}

	return nil
}

// GetStats returns statistics about the connection pool
func (p *DatabaseConnectionPool) GetStats() ConnectionPoolStats {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return ConnectionPoolStats{
		TotalConnections:  len(p.connections),
		ActiveConnections: p.activeConnections,
		IdleConnections:   len(p.connections) - p.activeConnections,
	}
}

// SetMaxConnections sets the maximum number of connections
func (p *DatabaseConnectionPool) SetMaxConnections(max int) {
	if max <= 0 {
		return
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.maxConnections = max
}

// ResetConnectionPool resets the connection pool (for testing purposes)
func ResetConnectionPool() {
	instance = nil
}

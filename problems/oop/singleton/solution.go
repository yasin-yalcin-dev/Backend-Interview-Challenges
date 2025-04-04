/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

// problems/oop/singleton/solution.go
package singleton

import "fmt"

// RunExample demonstrates the singleton and dependency injection patterns
func RunExample() {
	fmt.Println("Singleton & Dependency Injection Example:")
	fmt.Println("----------------------------------------")

	// Get the singleton instance of the connection pool
	connectionPool := GetDatabaseConnectionPool()

	// Show initial connection pool stats
	stats := connectionPool.GetStats()
	fmt.Printf("Initial connection pool stats - Total: %d, Active: %d, Idle: %d\n\n",
		stats.TotalConnections, stats.ActiveConnections, stats.IdleConnections)

	// Create services with dependency injection
	userService := NewUserService(connectionPool)
	productService := NewProductService(connectionPool)

	// Use the user service
	fmt.Println("Using UserService:")
	user, err := userService.GetUserByID("user123")
	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
	} else {
		fmt.Printf("Retrieved user: %s (%s)\n", user.Username, user.Email)
	}

	// Show updated connection pool stats
	stats = connectionPool.GetStats()
	fmt.Printf("Connection pool stats - Total: %d, Active: %d, Idle: %d\n\n",
		stats.TotalConnections, stats.ActiveConnections, stats.IdleConnections)

	// Use the product service
	fmt.Println("Using ProductService:")
	product, err := productService.GetProductByID("product456")
	if err != nil {
		fmt.Printf("Error getting product: %v\n", err)
	} else {
		fmt.Printf("Retrieved product: %s (Price: $%.2f)\n", product.Name, product.Price)
	}

	// Show final connection pool stats
	stats = connectionPool.GetStats()
	fmt.Printf("\nFinal connection pool stats - Total: %d, Active: %d, Idle: %d\n",
		stats.TotalConnections, stats.ActiveConnections, stats.IdleConnections)

	// Demonstrate connection reuse
	fmt.Println("\nDemonstrating connection reuse:")

	// Create multiple connections
	fmt.Println("Opening multiple connections...")
	conns := make([]*DatabaseConnection, 0)
	for i := 0; i < 5; i++ {
		conn, err := connectionPool.GetConnection()
		if err != nil {
			fmt.Printf("Error getting connection: %v\n", err)
			continue
		}
		conns = append(conns, conn)
	}

	// Show stats with multiple active connections
	stats = connectionPool.GetStats()
	fmt.Printf("Connection pool stats - Total: %d, Active: %d, Idle: %d\n",
		stats.TotalConnections, stats.ActiveConnections, stats.IdleConnections)

	// Release connections
	fmt.Println("Releasing connections...")
	for _, conn := range conns {
		connectionPool.ReleaseConnection(conn)
	}

	// Show final stats
	stats = connectionPool.GetStats()
	fmt.Printf("Final connection pool stats - Total: %d, Active: %d, Idle: %d\n",
		stats.TotalConnections, stats.ActiveConnections, stats.IdleConnections)
}

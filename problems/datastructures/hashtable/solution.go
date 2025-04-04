/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package hashtable

import (
	"fmt"
	"time"
)

// RunExample demonstrates the hash table implementation
func RunExample() {
	fmt.Println("Hash Table Example:")
	fmt.Println("------------------")

	// Create a new hash table
	ht := NewHashTable()

	// Check if it's empty
	fmt.Printf("Empty hash table size: %d\n", ht.Size())

	// Basic operations
	fmt.Println("\nBasic Operations:")
	fmt.Println("1. Adding key-value pairs")
	ht.Put("name", "John Doe")
	ht.Put("age", 30)
	ht.Put("email", "john@example.com")

	fmt.Printf("Hash table size: %d\n", ht.Size())
	fmt.Printf("Keys: %v\n", ht.Keys())

	// Retrieving values
	fmt.Println("\n2. Retrieving values")
	name, found := ht.Get("name")
	fmt.Printf("name: %v (found: %v)\n", name, found)

	phone, found := ht.Get("phone")
	fmt.Printf("phone: %v (found: %v)\n", phone, found)

	// Checking if keys exist
	fmt.Printf("\nContains 'email': %v\n", ht.Contains("email"))
	fmt.Printf("Contains 'address': %v\n", ht.Contains("address"))

	// Deleting a key
	fmt.Println("\n3. Deleting a key")
	fmt.Printf("Before deletion - Size: %d, Keys: %v\n", ht.Size(), ht.Keys())
	ht.Delete("age")
	fmt.Printf("After deletion - Size: %d, Keys: %v\n", ht.Size(), ht.Keys())

	// Updating a value
	fmt.Println("\n4. Updating a value")
	ht.Put("email", "johndoe@example.com")
	email, _ := ht.Get("email")
	fmt.Printf("Updated email: %v\n", email)

	// Clearing the hash table
	fmt.Println("\n5. Clearing the hash table")
	ht.Clear()
	fmt.Printf("After clearing - Size: %d, Keys: %v\n", ht.Size(), ht.Keys())

	// Demonstrating resizing
	fmt.Println("\nDemonstrating Resizing:")
	ht = NewHashTableWithCapacity(4) // Start with a small capacity
	fmt.Printf("Initial capacity: %d\n", ht.GetCapacity())

	// Add many key-value pairs to trigger resizing
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("key%d", i)
		ht.Put(key, i)
		// Print status after each insertion to show resizing
		if i%4 == 0 {
			fmt.Printf("After %d insertions - Size: %d, Capacity: %d, Load Factor: %.2f\n",
				i+1, ht.Size(), ht.GetCapacity(), ht.GetLoadFactor())
		}
	}

	// Practical Application: Simple Cache
	fmt.Println("\nPractical Application: Simple Cache")
	cache := NewHashTable()

	// Simulate expensive operation
	fetchData := func(key string) interface{} {
		fmt.Printf("Fetching data for '%s' (expensive operation)...\n", key)
		time.Sleep(100 * time.Millisecond) // Simulate work
		return fmt.Sprintf("Data for %s", key)
	}

	// Get data with caching
	getData := func(key string) interface{} {
		// Check if in cache
		if value, found := cache.Get(key); found {
			fmt.Printf("Cache hit for '%s'\n", key)
			return value
		}

		// Not in cache, fetch data
		value := fetchData(key)
		cache.Put(key, value)
		return value
	}

	// Demo the cache
	keys := []string{"user1", "user2", "user1", "user3", "user2"}
	for _, key := range keys {
		fmt.Printf("Requesting: %s\n", key)
		result := getData(key)
		fmt.Printf("Result: %v\n\n", result)
	}

	fmt.Printf("Cache size: %d\n", cache.Size())
	fmt.Printf("Cache keys: %v\n", cache.Keys())
}

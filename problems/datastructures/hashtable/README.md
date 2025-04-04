# Hash Table

## Problem
Implement a Hash Table data structure that efficiently stores and retrieves key-value pairs using a hash function. A hash table (also known as a hash map) uses a hash function to compute an index into an array of buckets, from which the desired value can be found.

## Requirements
1. Implement a hash table with the following features:
   - Support for string keys and interface{} values
   - A robust hash function to minimize collisions
   - Collision resolution using chaining (linked lists at each bucket)
   - Automatic resizing to maintain performance
   - Basic operations:
     - Put(key, value) - Insert or update a key-value pair
     - Get(key) - Retrieve a value by key
     - Delete(key) - Remove a key-value pair
     - Contains(key) - Check if a key exists
     - Size() - Get the number of key-value pairs
     - Clear() - Remove all key-value pairs
     - Keys() - Get a list of all keys

## Examples
```go
// Create a new hash table
ht := NewHashTable()

// Insert key-value pairs
ht.Put("name", "John Doe")
ht.Put("age", 30)
ht.Put("email", "john@example.com")

// Retrieve values
name, found := ht.Get("name")
fmt.Println(name, found)  // Output: John Doe true

// Check if key exists
fmt.Println(ht.Contains("phone"))  // Output: false

// Get all keys
keys := ht.Keys()
fmt.Println(keys)  // Output: [name age email]

// Delete a key
ht.Delete("age")
fmt.Println(ht.Size())  // Output: 2

// Clear the hash table
ht.Clear()
fmt.Println(ht.Size())  // Output: 0
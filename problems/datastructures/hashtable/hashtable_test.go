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
	"sort"
	"testing"
)

func TestHashTableBasicOperations(t *testing.T) {
	// Create a new hash table
	ht := NewHashTable()

	// Test initial state
	if ht.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ht.Size())
	}

	// Test Put and Get
	ht.Put("key1", "value1")
	ht.Put("key2", 42)
	ht.Put("key3", true)

	if ht.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ht.Size())
	}

	val1, found := ht.Get("key1")
	if !found {
		t.Error("Expected key1 to be found")
	}
	if val1 != "value1" {
		t.Errorf("Expected value1, got %v", val1)
	}

	val2, found := ht.Get("key2")
	if !found {
		t.Error("Expected key2 to be found")
	}
	if val2 != 42 {
		t.Errorf("Expected 42, got %v", val2)
	}

	// Test Contains
	if !ht.Contains("key3") {
		t.Error("Expected key3 to be contained")
	}

	if ht.Contains("nonexistent") {
		t.Error("Expected nonexistent key to not be contained")
	}

	// Test updating existing key
	ht.Put("key1", "updated_value")
	val1, _ = ht.Get("key1")
	if val1 != "updated_value" {
		t.Errorf("Expected updated_value, got %v", val1)
	}

	// Check size didn't change after update
	if ht.Size() != 3 {
		t.Errorf("Expected size to remain 3, got %d", ht.Size())
	}
}

func TestHashTableDelete(t *testing.T) {
	// Create a new hash table
	ht := NewHashTable()

	// Add some key-value pairs
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	// Test deleting existing key
	deleted := ht.Delete("key2")
	if !deleted {
		t.Error("Delete should return true for existing key")
	}

	if ht.Size() != 2 {
		t.Errorf("Expected size 2 after deletion, got %d", ht.Size())
	}

	if ht.Contains("key2") {
		t.Error("key2 should be deleted")
	}

	// Test deleting non-existent key
	deleted = ht.Delete("nonexistent")
	if deleted {
		t.Error("Delete should return false for non-existent key")
	}

	if ht.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", ht.Size())
	}

	// Delete all remaining keys
	ht.Delete("key1")
	ht.Delete("key3")

	if ht.Size() != 0 {
		t.Errorf("Expected size 0 after deleting all keys, got %d", ht.Size())
	}
}

func TestHashTableKeys(t *testing.T) {
	// Create a new hash table
	ht := NewHashTable()

	// Add some key-value pairs
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	// Test Keys method
	keys := ht.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Sort keys for deterministic comparison
	sort.Strings(keys)
	expectedKeys := []string{"key1", "key2", "key3"}
	for i, key := range expectedKeys {
		if keys[i] != key {
			t.Errorf("Expected key %s at index %d, got %s", key, i, keys[i])
		}
	}
}

func TestHashTableClear(t *testing.T) {
	// Create a new hash table
	ht := NewHashTable()

	// Add some key-value pairs
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")

	// Test Clear method
	ht.Clear()

	if ht.Size() != 0 {
		t.Errorf("Expected size 0 after clearing, got %d", ht.Size())
	}

	if ht.Contains("key1") || ht.Contains("key2") {
		t.Error("Hash table should not contain any keys after clearing")
	}
}

func TestHashTableResizing(t *testing.T) {
	// Create a small hash table to trigger resizing
	ht := NewHashTableWithCapacity(4)

	initialCapacity := ht.GetCapacity()
	if initialCapacity != 4 {
		t.Errorf("Expected initial capacity 4, got %d", initialCapacity)
	}

	// Add enough key-value pairs to trigger resizing
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		ht.Put(key, i)
	}

	// Check if capacity increased
	newCapacity := ht.GetCapacity()
	if newCapacity <= initialCapacity {
		t.Errorf("Expected capacity to increase from %d, got %d", initialCapacity, newCapacity)
	}

	// Check if all values are still accessible
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		val, found := ht.Get(key)
		if !found {
			t.Errorf("Key %s should be found after resizing", key)
		}
		if val != i {
			t.Errorf("Expected value %d, got %v", i, val)
		}
	}
}

func TestHashTableCollisionHandling(t *testing.T) {
	// Create a hash table with a custom hash function for testing collisions
	ht := NewHashTableWithCapacity(4)

	// These keys should collide in a table with capacity 4
	// (assuming our hash function works as expected)
	collisionKeys := []string{
		"key0", "key4", "key8",
		"key1", "key5", "key9",
	}

	// Add the keys that should collide
	for i, key := range collisionKeys {
		ht.Put(key, i)
	}

	// Check that all values are retrievable despite collisions
	for i, key := range collisionKeys {
		val, found := ht.Get(key)
		if !found {
			t.Errorf("Key %s should be found despite collisions", key)
		}
		if val != i {
			t.Errorf("Expected value %d, got %v", i, val)
		}
	}

	// Delete a key from a collision chain
	ht.Delete("key4")

	// Verify that key4 is gone but other keys in the same bucket remain
	if ht.Contains("key4") {
		t.Error("key4 should be deleted")
	}

	if !ht.Contains("key0") || !ht.Contains("key8") {
		t.Error("Other keys in the same bucket should remain after deletion")
	}
}

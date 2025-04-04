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
	"hash/fnv"
	"strings"
)

// KeyValuePair represents a key-value pair in the hash table
type KeyValuePair struct {
	Key   string
	Value interface{}
	Next  *KeyValuePair // For chaining in case of collisions
}

// HashTable represents a hash table with chaining for collision resolution
type HashTable struct {
	buckets    []*KeyValuePair
	size       int
	capacity   int
	loadFactor float64
}

// DefaultCapacity is the initial capacity of the hash table
const DefaultCapacity = 16

// DefaultLoadFactor is the load factor threshold for resizing
const DefaultLoadFactor = 0.75

// NewHashTable creates a new hash table with default capacity
func NewHashTable() *HashTable {
	return NewHashTableWithCapacity(DefaultCapacity)
}

// NewHashTableWithCapacity creates a new hash table with the specified capacity
func NewHashTableWithCapacity(capacity int) *HashTable {
	if capacity < 1 {
		capacity = DefaultCapacity
	}

	return &HashTable{
		buckets:    make([]*KeyValuePair, capacity),
		size:       0,
		capacity:   capacity,
		loadFactor: DefaultLoadFactor,
	}
}

// hash computes the hash code for a key
func (ht *HashTable) hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

// getBucketIndex gets the bucket index for a key
func (ht *HashTable) getBucketIndex(key string) int {
	hashCode := ht.hash(key)
	return int(hashCode % uint32(ht.capacity))
}

// Put inserts or updates a key-value pair in the hash table
func (ht *HashTable) Put(key string, value interface{}) {
	// Check if we need to resize
	if float64(ht.size+1)/float64(ht.capacity) > ht.loadFactor {
		ht.resize(ht.capacity * 2)
	}

	index := ht.getBucketIndex(key)

	// Check if key already exists
	current := ht.buckets[index]
	for current != nil {
		if current.Key == key {
			// Key exists, update value
			current.Value = value
			return
		}
		current = current.Next
	}

	// Key doesn't exist, create new entry
	newPair := &KeyValuePair{
		Key:   key,
		Value: value,
		Next:  ht.buckets[index],
	}

	ht.buckets[index] = newPair
	ht.size++
}

// Get retrieves a value by key
func (ht *HashTable) Get(key string) (interface{}, bool) {
	index := ht.getBucketIndex(key)

	current := ht.buckets[index]
	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}

	return nil, false
}

// Contains checks if a key exists in the hash table
func (ht *HashTable) Contains(key string) bool {
	_, found := ht.Get(key)
	return found
}

// Delete removes a key-value pair from the hash table
func (ht *HashTable) Delete(key string) bool {
	index := ht.getBucketIndex(key)

	// Handle case where the key is at the head of the list
	if ht.buckets[index] != nil && ht.buckets[index].Key == key {
		ht.buckets[index] = ht.buckets[index].Next
		ht.size--
		return true
	}

	// Handle case where the key is further down the chain
	current := ht.buckets[index]
	for current != nil && current.Next != nil {
		if current.Next.Key == key {
			current.Next = current.Next.Next
			ht.size--
			return true
		}
		current = current.Next
	}

	return false
}

// Size returns the number of key-value pairs in the hash table
func (ht *HashTable) Size() int {
	return ht.size
}

// Clear removes all key-value pairs from the hash table
func (ht *HashTable) Clear() {
	ht.buckets = make([]*KeyValuePair, ht.capacity)
	ht.size = 0
}

// Keys returns a list of all keys in the hash table
func (ht *HashTable) Keys() []string {
	keys := make([]string, 0, ht.size)

	for _, bucket := range ht.buckets {
		current := bucket
		for current != nil {
			keys = append(keys, current.Key)
			current = current.Next
		}
	}

	return keys
}

// Values returns a list of all values in the hash table
func (ht *HashTable) Values() []interface{} {
	values := make([]interface{}, 0, ht.size)

	for _, bucket := range ht.buckets {
		current := bucket
		for current != nil {
			values = append(values, current.Value)
			current = current.Next
		}
	}

	return values
}

// resize resizes the hash table to the new capacity
func (ht *HashTable) resize(newCapacity int) {
	// Create a new hash table with the new capacity
	newHashTable := NewHashTableWithCapacity(newCapacity)

	// Rehash all key-value pairs
	for _, bucket := range ht.buckets {
		current := bucket
		for current != nil {
			newHashTable.Put(current.Key, current.Value)
			current = current.Next
		}
	}

	// Update this hash table
	ht.buckets = newHashTable.buckets
	ht.capacity = newCapacity
	// size doesn't change
}

// GetLoadFactor returns the current load factor of the hash table
func (ht *HashTable) GetLoadFactor() float64 {
	return float64(ht.size) / float64(ht.capacity)
}

// GetCapacity returns the current capacity of the hash table
func (ht *HashTable) GetCapacity() int {
	return ht.capacity
}

// GetBucketSizes returns the sizes of all buckets (for debugging)
func (ht *HashTable) GetBucketSizes() []int {
	sizes := make([]int, ht.capacity)

	for i, bucket := range ht.buckets {
		count := 0
		current := bucket
		for current != nil {
			count++
			current = current.Next
		}
		sizes[i] = count
	}

	return sizes
}

// String returns a string representation of the hash table
func (ht *HashTable) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("HashTable(size=%d, capacity=%d):\n", ht.size, ht.capacity))

	for i, bucket := range ht.buckets {
		if bucket != nil {
			sb.WriteString(fmt.Sprintf("  Bucket %d: ", i))

			current := bucket
			for current != nil {
				sb.WriteString(fmt.Sprintf("[%s: %v] ", current.Key, current.Value))
				current = current.Next
			}

			sb.WriteString("\n")
		}
	}

	return sb.String()
}

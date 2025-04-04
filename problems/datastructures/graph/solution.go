/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

// problems/datastructures/graph/solution.go
package graph

import (
	"fmt"
	"sort"
)

// RunExample demonstrates the graph implementation
func RunExample() {
	fmt.Println("Graph Example:")
	fmt.Println("--------------")

	// Create an undirected graph
	fmt.Println("\n1. Creating an undirected graph")
	undirectedGraph := NewGraph(false)

	// Add vertices
	vertices := []string{"A", "B", "C", "D", "E"}
	for _, v := range vertices {
		undirectedGraph.AddVertex(v)
	}

	// Add edges
	undirectedGraph.AddEdge("A", "B")
	undirectedGraph.AddEdge("A", "C", 2.5)
	undirectedGraph.AddEdge("B", "D", 1.8)
	undirectedGraph.AddEdge("C", "D")
	undirectedGraph.AddEdge("C", "E")
	undirectedGraph.AddEdge("D", "E", 3.2)

	// Print the graph
	fmt.Println(undirectedGraph)

	// Check if vertex and edge exist
	fmt.Printf("Has vertex 'A': %v\n", undirectedGraph.HasVertex("A"))
	fmt.Printf("Has vertex 'F': %v\n", undirectedGraph.HasVertex("F"))
	fmt.Printf("Has edge A-B: %v\n", undirectedGraph.HasEdge("A", "B"))
	fmt.Printf("Has edge A-E: %v\n", undirectedGraph.HasEdge("A", "E"))

	// Get neighbors
	neighbors, _ := undirectedGraph.GetNeighbors("C")
	fmt.Printf("Neighbors of C: %v\n", neighbors)

	// Check if graph is connected
	fmt.Printf("Is graph connected: %v\n", undirectedGraph.IsConnected())

	// Check if graph has a cycle
	fmt.Printf("Has cycle: %v\n", undirectedGraph.HasCycle())

	// BFS and DFS traversal
	fmt.Println("\n2. Graph Traversal")
	bfsOrder, _ := undirectedGraph.BFS("A")
	fmt.Printf("BFS from A: %v\n", bfsOrder)

	dfsOrder, _ := undirectedGraph.DFS("A")
	fmt.Printf("DFS from A: %v\n", dfsOrder)

	// Shortest path
	fmt.Println("\n3. Finding Shortest Paths")
	path, _ := undirectedGraph.ShortestPath("A", "E")
	fmt.Printf("Shortest path from A to E: %v\n", path)

	// Create a directed graph
	fmt.Println("\n4. Creating a directed graph")
	directedGraph := NewGraph(true)

	// Add vertices
	for _, v := range vertices {
		directedGraph.AddVertex(v)
	}

	// Add edges
	directedGraph.AddEdge("A", "B")
	directedGraph.AddEdge("B", "C")
	directedGraph.AddEdge("C", "D")
	directedGraph.AddEdge("D", "B") // This creates a cycle
	directedGraph.AddEdge("B", "E")

	// Print the graph
	fmt.Println(directedGraph)

	// Check if directed graph has a cycle
	fmt.Printf("Has cycle: %v\n", directedGraph.HasCycle())

	// Remove an edge
	fmt.Println("\n5. Removing an edge")
	directedGraph.RemoveEdge("D", "B")
	fmt.Printf("Has cycle after removal: %v\n", directedGraph.HasCycle())

	// Remove a vertex
	fmt.Println("\n6. Removing a vertex")
	directedGraph.RemoveVertex("C")
	fmt.Println(directedGraph)

	// Practical application: Social Network
	fmt.Println("\n7. Practical Application: Social Network")
	socialNetwork := createSocialNetworkGraph()
	fmt.Println(socialNetwork)

	// Find mutual friends
	person1 := "Alice"
	person2 := "Dave"
	mutualFriends := findMutualFriends(socialNetwork, person1, person2)
	fmt.Printf("Mutual friends between %s and %s: %v\n", person1, person2, mutualFriends)

	// Find friend recommendations based on mutual friends
	recommendations := recommendFriends(socialNetwork, "Alice")
	fmt.Printf("Friend recommendations for Alice: %v\n", recommendations)
}

// createSocialNetworkGraph creates a sample social network graph
func createSocialNetworkGraph() *Graph {
	g := NewGraph(false)

	// Add people
	people := []string{"Alice", "Bob", "Charlie", "Dave", "Eve", "Frank"}
	for _, person := range people {
		g.AddVertex(person)
	}

	// Add friendships (edges)
	g.AddEdge("Alice", "Bob")
	g.AddEdge("Alice", "Charlie")
	g.AddEdge("Alice", "Dave")
	g.AddEdge("Bob", "Charlie")
	g.AddEdge("Charlie", "Dave")
	g.AddEdge("Eve", "Frank")
	g.AddEdge("Dave", "Eve")

	return g
}

// findMutualFriends finds mutual friends between two people
func findMutualFriends(g *Graph, person1, person2 string) []interface{} {
	friends1, _ := g.GetNeighbors(person1)
	friends2, _ := g.GetNeighbors(person2)

	// Find intersection of friends
	mutualFriends := make([]interface{}, 0)

	for _, friend1 := range friends1 {
		for _, friend2 := range friends2 {
			if friend1 == friend2 {
				mutualFriends = append(mutualFriends, friend1)
			}
		}
	}

	return mutualFriends
}

// recommendFriends recommends potential friends based on mutual friends
func recommendFriends(g *Graph, person string) []interface{} {
	recommendations := make(map[interface{}]int) // person -> mutual friend count

	// Get direct friends
	friends, _ := g.GetNeighbors(person)

	// For each friend, get their friends
	for _, friend := range friends {
		friendsOfFriend, _ := g.GetNeighbors(friend)

		// Check if the friend's friend is not already a direct friend
		for _, fof := range friendsOfFriend {
			if fof != person && !contains(friends, fof) {
				recommendations[fof]++
			}
		}
	}

	// Sort recommendations by mutual friend count
	sortedRecs := make([]interface{}, 0, len(recommendations))
	for rec := range recommendations {
		if recommendations[rec] > 0 {
			sortedRecs = append(sortedRecs, rec)
		}
	}

	// Sort by number of mutual friends (descending)
	sort.Slice(sortedRecs, func(i, j int) bool {
		return recommendations[sortedRecs[i]] > recommendations[sortedRecs[j]]
	})

	return sortedRecs
}

// contains checks if a slice contains a value
func contains(slice []interface{}, value interface{}) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

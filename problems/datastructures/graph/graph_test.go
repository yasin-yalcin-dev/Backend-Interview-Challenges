/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

// problems/datastructures/graph/graph_test.go
package graph

import (
	"reflect"
	"testing"
)

func TestUndirectedGraph(t *testing.T) {
	g := NewGraph(false)

	// Test adding vertices
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")

	if !g.HasVertex("A") || !g.HasVertex("B") || !g.HasVertex("C") {
		t.Error("Expected vertices to be added")
	}

	// Test adding duplicate vertex (should return false)
	if g.AddVertex("A") {
		t.Error("Adding duplicate vertex should return false")
	}

	// Test vertex count
	if g.VertexCount() != 3 {
		t.Errorf("Expected 3 vertices, got %d", g.VertexCount())
	}

	// Test adding edges
	g.AddEdge("A", "B")
	g.AddEdge("B", "C", 2.5)

	if !g.HasEdge("A", "B") || !g.HasEdge("B", "A") {
		t.Error("Expected edge A-B to exist in both directions (undirected)")
	}

	if !g.HasEdge("B", "C") || !g.HasEdge("C", "B") {
		t.Error("Expected edge B-C to exist in both directions (undirected)")
	}

	// Test edge weights
	weight, _ := g.GetEdgeWeight("A", "B")
	if weight != 1.0 {
		t.Errorf("Expected default weight 1.0, got %.1f", weight)
	}

	weight, _ = g.GetEdgeWeight("B", "C")
	if weight != 2.5 {
		t.Errorf("Expected weight 2.5, got %.1f", weight)
	}

	// Test edge count
	if g.EdgeCount() != 2 {
		t.Errorf("Expected 2 edges, got %d", g.EdgeCount())
	}

	// Test getting neighbors
	neighbors, _ := g.GetNeighbors("B")
	expectedNeighbors := []interface{}{"A", "C"}
	if !reflect.DeepEqual(neighbors, expectedNeighbors) {
		t.Errorf("Expected neighbors %v, got %v", expectedNeighbors, neighbors)
	}

	// Test removing edge
	g.RemoveEdge("A", "B")
	if g.HasEdge("A", "B") || g.HasEdge("B", "A") {
		t.Error("Expected edge A-B to be removed")
	}

	// Test removing vertex
	g.RemoveVertex("B")
	if g.HasVertex("B") {
		t.Error("Expected vertex B to be removed")
	}

	// Edge B-C should also be removed
	if g.HasEdge("B", "C") || g.HasEdge("C", "B") {
		t.Error("Expected edge B-C to be removed")
	}
}

func TestDirectedGraph(t *testing.T) {
	g := NewGraph(true)

	// Add vertices
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")

	// Add edges

	g.AddEdge("A", "B")
	g.AddEdge("B", "C")

	// In directed graph, edges should only exist in one direction
	if !g.HasEdge("A", "B") {
		t.Error("Expected edge A->B to exist")
	}

	if g.HasEdge("B", "A") {
		t.Error("Edge B->A should not exist in directed graph")
	}

	// Test edge count in directed graph
	if g.EdgeCount() != 2 {
		t.Errorf("Expected 2 edges, got %d", g.EdgeCount())
	}

	// Test removing edge in directed graph
	g.RemoveEdge("A", "B")
	if g.HasEdge("A", "B") {
		t.Error("Expected edge A->B to be removed")
	}
}

func TestGraphTraversal(t *testing.T) {
	g := NewGraph(false)

	// Create test graph
	//      A
	//     / \
	//    B   C
	//   /     \
	//  D-------E

	// Add vertices
	vertices := []string{"A", "B", "C", "D", "E"}
	for _, v := range vertices {
		g.AddVertex(v)
	}

	// Add edges
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")
	g.AddEdge("D", "E")

	// Test BFS
	bfsOrder, _ := g.BFS("A")
	// BFS could be [A, B, C, D, E] or [A, C, B, E, D] depending on implementation
	// Just check that A is first and all vertices are present
	if bfsOrder[0] != "A" || len(bfsOrder) != 5 {
		t.Errorf("Invalid BFS order: %v", bfsOrder)
	}

	// Test DFS
	dfsOrder, _ := g.DFS("A")
	// DFS could have different orders too
	if dfsOrder[0] != "A" || len(dfsOrder) != 5 {
		t.Errorf("Invalid DFS order: %v", dfsOrder)
	}

	// Test connected graph
	if !g.IsConnected() {
		t.Error("Graph should be connected")
	}

	// Disconnect part of the graph
	g.RemoveEdge("A", "C")
	g.RemoveEdge("D", "E")

	// Now graph is not connected
	if g.IsConnected() {
		t.Error("Graph should not be connected")
	}
}

func TestShortestPath(t *testing.T) {
	g := NewGraph(false)

	// Create test graph with weights
	//       A
	//     /   \
	//   1/     \2
	//   B       C
	//   |       |
	//  3|       |1
	//   D-------E
	//       2

	// Add vertices
	vertices := []string{"A", "B", "C", "D", "E"}
	for _, v := range vertices {
		g.AddVertex(v)
	}

	// Add weighted edges
	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 2)
	g.AddEdge("B", "D", 3)
	g.AddEdge("C", "E", 1)
	g.AddEdge("D", "E", 2)

	// Test shortest path
	path, _ := g.ShortestPath("A", "E")
	expectedPath := []interface{}{"A", "C", "E"} // A-C-E (total: 3) is shorter than A-B-D-E (total: 6)

	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Expected shortest path %v, got %v", expectedPath, path)
	}

	// Test path to non-existent vertex
	_, err := g.ShortestPath("A", "F")
	if err == nil {
		t.Error("ShortestPath to non-existent vertex should return error")
	}

	// Test disconnected graph
	// Completely disconnect E from the rest of the graph
	g.RemoveEdge("C", "E")
	g.RemoveEdge("D", "E")

	_, err = g.ShortestPath("A", "E")
	if err == nil {
		t.Error("ShortestPath in disconnected graph should return error")
	}
}

func TestCycleDetection(t *testing.T) {
	// Test in undirected graph
	g1 := NewGraph(false)

	// Add vertices
	for _, v := range []string{"A", "B", "C", "D"} {
		g1.AddVertex(v)
	}

	// No cycle initially
	g1.AddEdge("A", "B")
	g1.AddEdge("B", "C")
	g1.AddEdge("C", "D")

	if g1.HasCycle() {
		t.Error("Undirected graph should not have cycle")
	}

	// Add edge to create cycle
	g1.AddEdge("A", "D")

	if !g1.HasCycle() {
		t.Error("Undirected graph should have cycle")
	}

	// Test in directed graph
	g2 := NewGraph(true)

	// Add vertices
	for _, v := range []string{"A", "B", "C", "D"} {
		g2.AddVertex(v)
	}

	// No cycle initially
	g2.AddEdge("A", "B")
	g2.AddEdge("B", "C")
	g2.AddEdge("C", "D")

	if g2.HasCycle() {
		t.Error("Directed graph should not have cycle")
	}

	// Add edge to create cycle
	g2.AddEdge("D", "B")

	if !g2.HasCycle() {
		t.Error("Directed graph should have cycle")
	}
}

func TestSocialNetworkFunctions(t *testing.T) {
	g := createSocialNetworkGraph()

	// Test mutual friends
	mutualFriends := findMutualFriends(g, "Alice", "Dave")
	expectedFriends := []interface{}{"Charlie"}
	if !reflect.DeepEqual(mutualFriends, expectedFriends) {
		t.Errorf("Expected mutual friends %v, got %v", expectedFriends, mutualFriends)
	}

	// Test friend recommendations
	// Alice is friends with Bob, Charlie, and Dave
	// Friend recommendations should include Eve (friend of Dave)
	recommendations := recommendFriends(g, "Alice")

	if len(recommendations) == 0 || recommendations[0] != "Eve" {
		t.Errorf("Expected Eve as a top recommendation, got %v", recommendations)
	}
}

# Graph

## Problem
Implement a graph data structure that can represent relationships between entities, supporting both directed and undirected graphs, as well as weighted and unweighted edges. The implementation should include standard graph operations and algorithms.

## Requirements
1. Implement a graph with the following features:
   - Support for both directed and undirected graphs
   - Support for weighted and unweighted edges
   - Adjacency list representation for efficient storage
   - Basic operations:
     - AddVertex - Add a new vertex to the graph
     - AddEdge - Add an edge between two vertices
     - RemoveVertex - Remove a vertex and all its edges
     - RemoveEdge - Remove an edge between two vertices
     - GetVertices - Get all vertices in the graph
     - GetNeighbors - Get all neighbors of a vertex
     - HasVertex - Check if a vertex exists
     - HasEdge - Check if an edge exists between two vertices
   - Traversal algorithms:
     - Breadth-First Search (BFS)
     - Depth-First Search (DFS)
   - Path finding and analysis:
     - Find shortest path (Dijkstra's algorithm)
     - Check if the graph is connected
     - Detect cycles

## Examples
```go
// Create a new undirected graph
graph := NewGraph(false)

// Add vertices
graph.AddVertex("A")
graph.AddVertex("B")
graph.AddVertex("C")
graph.AddVertex("D")

// Add edges (unweighted by default)
graph.AddEdge("A", "B")
graph.AddEdge("A", "C")
graph.AddEdge("B", "D")
graph.AddEdge("C", "D")

// Check if vertices and edges exist
fmt.Println(graph.HasVertex("A"))  // Output: true
fmt.Println(graph.HasEdge("A", "B"))  // Output: true

// Get neighbors of a vertex
neighbors := graph.GetNeighbors("A")
fmt.Println(neighbors)  // Output: [B C]

// Perform BFS from a starting vertex
bfsOrder := graph.BFS("A")
fmt.Println(bfsOrder)  // Output: [A B C D]

// Perform DFS from a starting vertex
dfsOrder := graph.DFS("A")
fmt.Println(dfsOrder)  // Output: [A B D C] or [A C D B]

// Find shortest path
path := graph.ShortestPath("A", "D")
fmt.Println(path)  // Output: [A B D] or [A C D]
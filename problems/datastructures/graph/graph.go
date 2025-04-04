/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package graph

import (
	"container/heap"
	"container/list"
	"errors"
	"fmt"
	"math"
	"sort"
)

// Edge represents an edge in the graph
type Edge struct {
	From   interface{}
	To     interface{}
	Weight float64
}

// Graph represents a graph using adjacency list
type Graph struct {
	vertices      map[interface{}]bool
	adjacencyList map[interface{}]map[interface{}]float64
	isDirected    bool
	defaultWeight float64
}

// NewGraph creates a new graph (directed or undirected)
func NewGraph(isDirected bool) *Graph {
	return &Graph{
		vertices:      make(map[interface{}]bool),
		adjacencyList: make(map[interface{}]map[interface{}]float64),
		isDirected:    isDirected,
		defaultWeight: 1.0, // Default weight for unweighted edges
	}
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(vertex interface{}) bool {
	if g.HasVertex(vertex) {
		return false // Vertex already exists
	}

	g.vertices[vertex] = true
	g.adjacencyList[vertex] = make(map[interface{}]float64)
	return true
}

// HasVertex checks if a vertex exists in the graph
func (g *Graph) HasVertex(vertex interface{}) bool {
	_, exists := g.vertices[vertex]
	return exists
}

// AddEdge adds an edge between two vertices
func (g *Graph) AddEdge(from, to interface{}, weight ...float64) error {
	// Check if vertices exist
	if !g.HasVertex(from) {
		return fmt.Errorf("vertex %v does not exist", from)
	}
	if !g.HasVertex(to) {
		return fmt.Errorf("vertex %v does not exist", to)
	}

	// Use default weight if not specified
	w := g.defaultWeight
	if len(weight) > 0 {
		w = weight[0]
		if w <= 0 {
			return errors.New("weight must be positive")
		}
	}

	// Add edge
	g.adjacencyList[from][to] = w

	// If undirected, add the reverse edge
	if !g.isDirected {
		g.adjacencyList[to][from] = w
	}

	return nil
}

// RemoveEdge removes an edge between two vertices
func (g *Graph) RemoveEdge(from, to interface{}) error {
	// Check if vertices exist
	if !g.HasVertex(from) {
		return fmt.Errorf("vertex %v does not exist", from)
	}
	if !g.HasVertex(to) {
		return fmt.Errorf("vertex %v does not exist", to)
	}

	// Check if edge exists
	if !g.HasEdge(from, to) {
		return fmt.Errorf("edge from %v to %v does not exist", from, to)
	}

	// Remove edge
	delete(g.adjacencyList[from], to)

	// If undirected, remove the reverse edge
	if !g.isDirected {
		delete(g.adjacencyList[to], from)
	}

	return nil
}

// HasEdge checks if an edge exists between two vertices
func (g *Graph) HasEdge(from, to interface{}) bool {
	if !g.HasVertex(from) || !g.HasVertex(to) {
		return false
	}

	_, exists := g.adjacencyList[from][to]
	return exists
}

// GetEdgeWeight returns the weight of an edge between two vertices
func (g *Graph) GetEdgeWeight(from, to interface{}) (float64, error) {
	if !g.HasVertex(from) {
		return 0, fmt.Errorf("vertex %v does not exist", from)
	}
	if !g.HasVertex(to) {
		return 0, fmt.Errorf("vertex %v does not exist", to)
	}

	weight, exists := g.adjacencyList[from][to]
	if !exists {
		return 0, fmt.Errorf("edge from %v to %v does not exist", from, to)
	}

	return weight, nil
}

// RemoveVertex removes a vertex and all its edges
func (g *Graph) RemoveVertex(vertex interface{}) error {
	if !g.HasVertex(vertex) {
		return fmt.Errorf("vertex %v does not exist", vertex)
	}

	// Remove all edges to this vertex from other vertices
	for v := range g.vertices {
		if v != vertex {
			delete(g.adjacencyList[v], vertex)
		}
	}

	// Remove vertex and its edges
	delete(g.adjacencyList, vertex)
	delete(g.vertices, vertex)

	return nil
}

// GetVertices returns all vertices in the graph
func (g *Graph) GetVertices() []interface{} {
	vertices := make([]interface{}, 0, len(g.vertices))
	for v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

// GetNeighbors returns all neighbors of a vertex
func (g *Graph) GetNeighbors(vertex interface{}) ([]interface{}, error) {
	if !g.HasVertex(vertex) {
		return nil, fmt.Errorf("vertex %v does not exist", vertex)
	}

	neighbors := make([]interface{}, 0, len(g.adjacencyList[vertex]))
	for n := range g.adjacencyList[vertex] {
		neighbors = append(neighbors, n)
	}

	// Sort neighbors for deterministic output (if they can be compared)
	sort.Slice(neighbors, func(i, j int) bool {
		return fmt.Sprintf("%v", neighbors[i]) < fmt.Sprintf("%v", neighbors[j])
	})

	return neighbors, nil
}

// GetEdges returns all edges in the graph
func (g *Graph) GetEdges() []Edge {
	edges := make([]Edge, 0)

	for from := range g.adjacencyList {
		for to, weight := range g.adjacencyList[from] {
			// For undirected graphs, only add edges where from < to to avoid duplicates
			if g.isDirected || fmt.Sprintf("%v", from) <= fmt.Sprintf("%v", to) {
				edges = append(edges, Edge{From: from, To: to, Weight: weight})
			}
		}
	}

	return edges
}

// VertexCount returns the number of vertices in the graph
func (g *Graph) VertexCount() int {
	return len(g.vertices)
}

// EdgeCount returns the number of edges in the graph
func (g *Graph) EdgeCount() int {
	count := 0

	for from := range g.adjacencyList {
		count += len(g.adjacencyList[from])
	}

	// For undirected graphs, each edge is counted twice
	if !g.isDirected {
		count /= 2
	}

	return count
}

// IsDirected returns true if the graph is directed
func (g *Graph) IsDirected() bool {
	return g.isDirected
}

// BFS performs breadth-first search from a starting vertex
func (g *Graph) BFS(start interface{}) ([]interface{}, error) {
	if !g.HasVertex(start) {
		return nil, fmt.Errorf("vertex %v does not exist", start)
	}

	visited := make(map[interface{}]bool)
	result := make([]interface{}, 0)

	// Create a queue
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		// Dequeue
		element := queue.Front()
		vertex := element.Value
		queue.Remove(element)

		// Add to result
		result = append(result, vertex)

		// Get neighbors
		neighbors, _ := g.GetNeighbors(vertex)

		// Visit neighbors
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}

	return result, nil
}

// dfsUtil is a helper function for DFS
func (g *Graph) dfsUtil(vertex interface{}, visited map[interface{}]bool, result *[]interface{}) {
	// Mark the current vertex as visited
	visited[vertex] = true
	*result = append(*result, vertex)

	// Visit all neighbors
	neighbors, _ := g.GetNeighbors(vertex)
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			g.dfsUtil(neighbor, visited, result)
		}
	}
}

// DFS performs depth-first search from a starting vertex
func (g *Graph) DFS(start interface{}) ([]interface{}, error) {
	if !g.HasVertex(start) {
		return nil, fmt.Errorf("vertex %v does not exist", start)
	}

	visited := make(map[interface{}]bool)
	result := make([]interface{}, 0)

	g.dfsUtil(start, visited, &result)

	return result, nil
}

// Item is a queue item for Dijkstra's algorithm
type Item struct {
	vertex   interface{}
	priority float64
	index    int // The index of the item in the heap
}

// A PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// ShortestPath finds the shortest path between two vertices using Dijkstra's algorithm
func (g *Graph) ShortestPath(start, end interface{}) ([]interface{}, error) {
	if !g.HasVertex(start) {
		return nil, fmt.Errorf("start vertex %v does not exist", start)
	}
	if !g.HasVertex(end) {
		return nil, fmt.Errorf("end vertex %v does not exist", end)
	}

	// Initialize distances and previous vertices
	dist := make(map[interface{}]float64)
	prev := make(map[interface{}]interface{})

	// Set distances to infinity for all vertices except start
	for v := range g.vertices {
		dist[v] = math.Inf(1)
	}
	dist[start] = 0

	// Initialize priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Add all vertices to the priority queue
	vertexToItem := make(map[interface{}]*Item)
	for v := range g.vertices {
		item := &Item{
			vertex:   v,
			priority: dist[v],
		}
		heap.Push(&pq, item)
		vertexToItem[v] = item
	}

	// Dijkstra's algorithm
	for pq.Len() > 0 {
		// Get vertex with minimum distance
		item := heap.Pop(&pq).(*Item)
		u := item.vertex

		// If we reached the end, we can stop
		if u == end {
			break
		}

		// For each neighbor of u
		for v, weight := range g.adjacencyList[u] {
			// Calculate potential new distance
			alt := dist[u] + weight

			// If we found a shorter path to v
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u

				// Update priority queue
				item := vertexToItem[v]
				item.priority = alt
				heap.Fix(&pq, item.index)
			}
		}
	}

	// If there's no path to the end vertex
	if _, ok := prev[end]; !ok && start != end {
		return nil, fmt.Errorf("no path from %v to %v", start, end)
	}

	// Reconstruct path
	path := make([]interface{}, 0)
	curr := end

	for curr != start {
		path = append([]interface{}{curr}, path...)
		curr = prev[curr]
	}
	path = append([]interface{}{start}, path...)

	return path, nil
}

// IsConnected checks if the graph is connected
func (g *Graph) IsConnected() bool {
	// If graph is empty, it's considered connected
	if len(g.vertices) == 0 {
		return true
	}

	// Get a starting vertex
	var start interface{}
	for v := range g.vertices {
		start = v
		break
	}

	// Perform BFS from start
	visited, _ := g.BFS(start)

	// If all vertices are visited, the graph is connected
	return len(visited) == len(g.vertices)
}

// HasCycle checks if the graph has a cycle
func (g *Graph) HasCycle() bool {
	if g.isDirected {
		// For directed graph
		visited := make(map[interface{}]bool)
		recStack := make(map[interface{}]bool)

		for v := range g.vertices {
			if !visited[v] {
				if g.hasCycleDirected(v, visited, recStack) {
					return true
				}
			}
		}
	} else {
		// For undirected graph
		visited := make(map[interface{}]bool)

		for v := range g.vertices {
			if !visited[v] {
				if g.hasCycleUndirected(v, nil, visited) {
					return true
				}
			}
		}
	}

	return false
}

// hasCycleDirected checks if a directed graph has a cycle using DFS
func (g *Graph) hasCycleDirected(vertex interface{}, visited, recStack map[interface{}]bool) bool {
	visited[vertex] = true
	recStack[vertex] = true

	for neighbor := range g.adjacencyList[vertex] {
		if !visited[neighbor] {
			if g.hasCycleDirected(neighbor, visited, recStack) {
				return true
			}
		} else if recStack[neighbor] {
			return true
		}
	}

	recStack[vertex] = false
	return false
}

// hasCycleUndirected checks if an undirected graph has a cycle using DFS
func (g *Graph) hasCycleUndirected(vertex, parent interface{}, visited map[interface{}]bool) bool {
	visited[vertex] = true

	for neighbor := range g.adjacencyList[vertex] {
		// If the neighbor is not visited, then recurse on it
		if !visited[neighbor] {
			if g.hasCycleUndirected(neighbor, vertex, visited) {
				return true
			}
		} else if neighbor != parent {
			// If the neighbor is visited and not the parent, then there's a cycle
			return true
		}
	}

	return false
}

// String returns a string representation of the graph
func (g *Graph) String() string {
	var result string

	if g.isDirected {
		result = "Directed Graph:\n"
	} else {
		result = "Undirected Graph:\n"
	}

	result += fmt.Sprintf("Vertices: %d, Edges: %d\n", g.VertexCount(), g.EdgeCount())

	for vertex := range g.adjacencyList {
		result += fmt.Sprintf("Vertex %v: ", vertex)

		edges := make([]string, 0)
		for neighbor, weight := range g.adjacencyList[vertex] {
			edges = append(edges, fmt.Sprintf("%v(%.1f)", neighbor, weight))
		}

		sort.Strings(edges)
		for i, edge := range edges {
			if i > 0 {
				result += ", "
			}
			result += edge
		}

		result += "\n"
	}

	return result
}

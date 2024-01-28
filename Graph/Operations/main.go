package main

import (
	"fmt"
)

type graph struct {
	vertices []*vertex
}
type vertex struct {
	data     int
	adjacent []*vertex
}

func main() {
	g := graph{}
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)
	g.addVertex(5)

	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(3, 4)
	g.addEdge(5, 3)
	g.addEdge(2, 4)

	g.print()

	g.deleteVertex(5)
	g.deleteEdge(1, 3)
	g.print()

}
func (g *graph) addVertex(data int) {
	if !contains(g.vertices, data) {
		g.vertices = append(g.vertices, &vertex{data: data})
	}
}
func contains(vertex []*vertex, data int) bool {
	for _, v := range vertex {
		if v.data == data {
			return true
		}
	}
	return false
}
func (g *graph) print() {
	for _, v := range g.vertices {
		fmt.Print("\n vertex", v.data, ":")
		for _, v := range v.adjacent {
			fmt.Print(" ", v.data)
		}

	}
}
func (g *graph) getVertex(data int) *vertex {
	for _, v := range g.vertices {
		if v.data == data {
			return v
		}
	}
	return nil
}
func (g *graph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		fmt.Println("Error: Vertex not found.")
		return
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
}
func (v *vertex) deleteAdjacent(data int) {
	for i, adj := range v.adjacent {
		if adj.data == data {
			v.adjacent = append(v.adjacent[:i], v.adjacent[i+1:]...)
		}
		return
	}
}
func (g *graph) deleteVertex(data int) {
	for i, v := range g.vertices {
		if v.data == data {
			g.vertices = append(g.vertices[:i], g.vertices[i+1:]...)
			for _, vertex := range g.vertices {
				vertex.deleteAdjacent(data)
			}
		}
	}
	return
}
func (g *graph) deleteEdge(from, to int) {
	Vfrom := g.getVertex(from)
	Vto := g.getVertex(to)
	if Vfrom == nil || Vto == nil {
		fmt.Println("Error")
		return
	}
	Vfrom.deleteAdjacent(to)
	Vto.deleteAdjacent(from)
}

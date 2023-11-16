package ktools

import (
	"time"
)

// InvoiceStatusVertex is a vertex (node) of the status graph
type vertex struct {
	class       string
	id          string
	name        string
	description string
}

// InvoiceStatusEdge is an edge of the status graph
type edge struct {
	class       string
	vertexFrom  string
	vertexTo    string
	description string
}

type edgeQueueItem struct {
	id          string
	class       string
	vertexFrom  string
	vertexTo    string
	description string
	author      string
	createdAt   time.Time
}

// KGraph applys graph theory to keel system
type KGraph struct {
	vertexMap map[string]*vertex
	edgeMap   map[string]*edge
	edgeQueue map[string][]*edgeQueueItem
}

// NewKGraph creates a new graph of the status of an invoice
func NewKGraph() *KGraph {
	return &KGraph{
		vertexMap: make(map[string]*vertex),
		edgeMap:   make(map[string]*edge),
		edgeQueue: make(map[string][]*edgeQueueItem),
	}
}

// AddVertex adds a new vertex to the graph
func (g *KGraph) AddVertex(class, id, name, description string) {
	vertex := vertex{
		class:       class,
		id:          id,
		name:        name,
		description: description,
	}
	g.vertexMap[class+id] = &vertex
}

// AddEdge adds a new edge to the graph
func (g *KGraph) AddEdge(class, vertexFrom, vertexTo, description string) {
	edge := edge{
		class:       class,
		vertexFrom:  vertexFrom,
		vertexTo:    vertexTo,
		description: description,
	}
	g.edgeMap[class+vertexFrom+vertexTo] = &edge
}

// CheckVertex checks if a vertex exists
func (g *KGraph) CheckEdge(class, vertexFrom, vertexTo string) bool {
	_, ok := g.edgeMap[class+vertexFrom+vertexTo]
	return ok
}

func (g *KGraph) EnqueueEdge(id, class, vertexFrom, vertexTo, description, author string) {
	edgeLog := &edgeQueueItem{
		id:          id,
		class:       class,
		vertexFrom:  vertexFrom,
		vertexTo:    vertexTo,
		description: description,
		author:      author,
		createdAt:   time.Now(),
	}
	g.edgeQueue[class] = append(g.edgeQueue[class], edgeLog)
}

// GetEdgeLog
func (g *KGraph) DequeueEdge(class string) (bool, string, string, string, string, string, time.Time) {
	if len(g.edgeQueue[class]) == 0 {
		return false, "", "", "", "", "", time.Time{}
	}
	item := g.edgeQueue[class][0]
	g.edgeQueue[class] = g.edgeQueue[class][1:]
	return true, item.id, item.vertexFrom, item.vertexTo, item.description, item.author, item.createdAt
}

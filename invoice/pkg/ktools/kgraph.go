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

type edgeLog struct {
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
	edgeLog   map[string][]*edgeLog
}

// NewKGraph creates a new graph of the status of an invoice
func NewKGraph() *KGraph {
	return &KGraph{}
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

func (g *KGraph) EnqueueEdge(class, vertexFrom, vertexTo, description, author string) {
	edgeLog := &edgeLog{
		class:       class,
		vertexFrom:  vertexFrom,
		vertexTo:    vertexTo,
		description: description,
		author:      author,
		createdAt:   time.Now(),
	}
	g.edgeLog[class] = append(g.edgeLog[class], edgeLog)
}

// GetEdgeLog
func (g *KGraph) DequeueEdge(class string) (bool, string, string, string, string, time.Time) {
	if len(g.edgeLog[class]) == 0 {
		return false, "", "", "", "", time.Time{}
	}
	edgeLog := g.edgeLog[class][0]
	g.edgeLog[class] = g.edgeLog[class][1:]
	return true, edgeLog.class, edgeLog.vertexFrom, edgeLog.vertexTo, edgeLog.description, edgeLog.createdAt
}

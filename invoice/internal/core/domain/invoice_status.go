package domain

// InvoiceStatusVertex is a vertex (node) of the status graph
type invoiceStatusVertex struct {
	id          int
	internalName string
	name        string
	description string
}

// InvoiceStatusEdge is an edge of the status graph
type invoiceStatusEdge struct {
	id          int
	vertexFrom  int
	vertexTo    int
	description string
}

// InvoiceStatusGraph is a graph of the status of an invoice
type InvoiceStatusGraph struct {
	verticesById    map[int]*invoiceStatusVertex
	verticesByName  map[string]*invoiceStatusVertex
	edgesById       map[int]*invoiceStatusEdge
	edgesByTo       map[int][]*invoiceStatusEdge
}

// NewInvoiceStatusGraph creates a new graph of the status of an invoice
func NewInvoiceStatusGraph() *InvoiceStatusGraph {
	return &InvoiceStatusGraph{}
}

// AddVertex adds a new vertex to the graph
func (g *InvoiceStatusGraph) AddVertex(id int, internalName, name, description string) {
	vertex := invoiceStatusVertex{
		id:           id,
		internalName: internalName,
		name:         name,
		description:  description,
	}
	g.verticesById[id] = &vertex
	g.verticesByName[name] = &vertex
}

// AddEdge adds a new edge to the graph
func (g *InvoiceStatusGraph) AddEdge(id, vertexFrom, vertexTo int, description string) {
	edge := invoiceStatusEdge{
		id:          id,
		vertexFrom:  vertexFrom,
		vertexTo:    vertexTo,
		description: description,
	}
	g.edgesById[id] = &edge
	g.edgesByTo[vertexTo] = append(g.edgesByTo[vertexTo], &edge)
}


// GetVertexIdByInternalName returns a vertex by its internal name
func (g *InvoiceStatusGraph) GetVertexIdByInternalName(internalName string) int {
	return g.verticesByName[internalName].id
}

// GetEdgeIdByVertexTo returns an edge by its vertex to
func (g *InvoiceStatusGraph) GetEdgeIdByVertexTo(vertexTo int) int {
	return g.edgesByTo[vertexTo][0].id
}

// CheckEdge checks if an edge exists
func (g *InvoiceStatusGraph) CheckEdge(vertexFrom, vertexTo int) bool {
	edges := g.edgesByTo[vertexTo]
	for _, edge := range edges {
		if edge.vertexFrom == vertexFrom {
			return true
		}
	}
	return false
}

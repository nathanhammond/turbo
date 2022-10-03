package dagextend

import (
	"fmt"

	"github.com/pyr-sh/dag"
)

// NamedEdge is a basic implementation of Edge that has the source and
// target vertex.
type NamedEdge struct {
	Name string
	S    dag.Vertex
	T    dag.Vertex
	dag.Edge
}

func (e *NamedEdge) Hashcode() interface{} {
	return fmt.Sprintf("%p-%s-%p", e.S, e.Name, e.T)
}

func (e *NamedEdge) Source() dag.Vertex {
	return e.S
}

func (e *NamedEdge) Target() dag.Vertex {
	return e.T
}

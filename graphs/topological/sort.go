package topological

import (
	"errors"

	"github.com/Zaba505/adhoc/graphs/model"
)

var ErrUnsortable = errors.New("topological: graph is not sortable")

// Sort
func Sort(g *model.Graph) ([]*model.Subject, error) {
	triples := toSet(g.Triples)

	var L []*model.Subject
	S := getStartingNodes(triples)

	for len(S) > 0 {
		n := popFirst(S)
		L = append(L, n)

		nEdges := getOutgoingEdges(n, triples)
		for _, e := range nEdges {
			delete(triples, e)

			o := e.Object
			os, isSubject := o.Value.(*model.Object_Subject)
			if !isSubject {
				continue
			}

			m := os.Subject
			mEdges := getIncomingEdges(m, triples)
			if len(mEdges) > 1 {
				continue
			}

			S[m] = struct{}{}
		}
	}

	if len(triples) > 0 {
		return nil, ErrUnsortable
	}

	return L, nil
}

func toSet[T comparable](ts []T) map[T]struct{} {
	m := make(map[T]struct{})
	for _, t := range ts {
		m[t] = struct{}{}
	}
	return m
}

func getStartingNodes(triples map[*model.Triple]struct{}) map[*model.Subject]struct{} {
	m := make(map[*model.Subject]struct{})

	// Identify all subjects
	for t := range triples {
		s := t.Subject
		m[s] = struct{}{}

		o := t.Object
		os, isSubject := o.Value.(*model.Object_Subject)
		if !isSubject {
			continue
		}
		s = os.Subject
		m[s] = struct{}{}
	}

	// Remove all subjects which have an edge pointing towards them
	for t := range triples {
		val := t.Object.Value
		os, isSubject := val.(*model.Object_Subject)
		if !isSubject {
			continue
		}

		s := os.Subject
		delete(m, s)
	}
	return m
}

func popFirst[T comparable](set map[T]struct{}) T {
	var val T
	for x := range set {
		val = x
		break
	}
	delete(set, val)
	return val
}

// getIncomingEdges returns all the incoming edges to the given node
func getIncomingEdges(n *model.Subject, triples map[*model.Triple]struct{}) []*model.Triple {
	var nTriples []*model.Triple
	for t := range triples {
		o := t.Object
		os, isSubject := o.Value.(*model.Object_Subject)
		if !isSubject {
			continue
		}

		s := os.Subject
		if n != s {
			continue
		}
		nTriples = append(nTriples, t)
	}
	return nTriples
}

// getOutgoingEdges returns all the outgoing edges from the given node
func getOutgoingEdges(n *model.Subject, triples map[*model.Triple]struct{}) []*model.Triple {
	var nTriples []*model.Triple
	for t := range triples {
		if t.Subject != n {
			continue
		}
		nTriples = append(nTriples, t)
	}
	return nTriples
}

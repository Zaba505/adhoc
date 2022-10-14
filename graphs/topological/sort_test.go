package topological

import (
	"testing"

	"github.com/Zaba505/adhoc/graphs/model"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	t.Run("will fail to sort", func(t *testing.T) {
		t.Run("if the graph contains a cycle", func(t *testing.T) {
			subA := node("Node", "A")
			subB := node("Node", "B")

			a := withSubject(subA)
			b := withSubject(subB)

			g := &model.Graph{
				Triples: []*model.Triple{
					a(pred("type"), str(subA.Type)),
					a(pred("type_id"), str(subA.Tuid)),
					a(pred("friend"), subject(subB)),
					b(pred("type"), str(subB.Type)),
					b(pred("type_id"), str(subB.Tuid)),
					b(pred("friend"), subject(subA)),
				},
			}

			_, err := Sort(g)
			if !assert.Equal(t, ErrUnsortable, err) {
				return
			}
		})
	})

	t.Run("will succeed to sort", func(t *testing.T) {
		t.Run("if the graph is made of only attribute edges", func(t *testing.T) {
			nodeA := node("Node", "A")

			a := withSubject(nodeA)

			g := &model.Graph{
				Triples: []*model.Triple{
					a(pred("type"), str(nodeA.Type)),
					a(pred("type_id"), str(nodeA.Tuid)),
				},
			}

			sortedNodes, err := Sort(g)
			if !assert.Nil(t, err) {
				return
			}
			if !assert.Equal(t, 1, len(sortedNodes)) {
				return
			}
		})

		t.Run("if the graph contains multiple nodes", func(t *testing.T) {
			nodeA := node("Node", "A")
			nodeB := node("Node", "B")
			nodeC := node("Node", "C")
			nodeD := node("Node", "D")

			a := withSubject(nodeA)
			b := withSubject(nodeB)
			c := withSubject(nodeC)
			d := withSubject(nodeD)

			g := &model.Graph{
				Triples: []*model.Triple{
					a(pred("type"), str(nodeA.Type)),
					a(pred("type_id"), str(nodeA.Tuid)),
					a(pred("friend"), subject(nodeB)),
					a(pred("friend"), subject(nodeC)),

					b(pred("type"), str(nodeB.Type)),
					b(pred("type_id"), str(nodeB.Tuid)),
					b(pred("friend"), subject(nodeD)),

					c(pred("type"), str(nodeC.Type)),
					c(pred("type_id"), str(nodeC.Tuid)),
					c(pred("friend"), subject(nodeD)),

					d(pred("type"), str(nodeD.Type)),
					d(pred("type_id"), str(nodeD.Tuid)),
				},
			}

			sortedNodes, err := Sort(g)
			if !assert.Nil(t, err) {
				return
			}
			if !assert.Equal(t, 4, len(sortedNodes)) {
				t.Log(sortedNodes)
				return
			}
		})
	})
}

func node(typ, tuid string) *model.Subject {
	return &model.Subject{
		Type: typ,
		Tuid: tuid,
	}
}

func withSubject(s *model.Subject) func(*model.Predicate, *model.Object) *model.Triple {
	return func(p *model.Predicate, o *model.Object) *model.Triple {
		return &model.Triple{
			Subject:   s,
			Predicate: p,
			Object:    o,
		}
	}
}

func pred(name string) *model.Predicate {
	return &model.Predicate{
		Name: name,
	}
}

func str(s string) *model.Object {
	return &model.Object{
		Value: &model.Object_String_{
			String_: s,
		},
	}
}

func subject(s *model.Subject) *model.Object {
	return &model.Object{
		Value: &model.Object_Subject{
			Subject: s,
		},
	}
}

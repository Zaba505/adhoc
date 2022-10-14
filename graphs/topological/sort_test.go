package topological

import (
	"testing"

	"github.com/Zaba505/adhoc/graphs/model"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	t.Run("will fail to sort", func(t *testing.T) {
		t.Run("if the graph is not a directed acyclic graph", func(t *testing.T) {
			subA := &model.Subject{
				Type: "Node",
				Tuid: "A",
			}
			subB := &model.Subject{
				Type: "Node",
				Tuid: "B",
			}

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

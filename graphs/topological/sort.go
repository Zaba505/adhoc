package topological

import (
	"errors"

	"github.com/Zaba505/adhoc/graphs/model"
)

var ErrUnsortable = errors.New("topological: graph is not sortable")

func Sort(g *model.Graph) ([]*model.Subject, error) {
	return nil, nil
}

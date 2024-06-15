package vec

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Float | constraints.Integer
}

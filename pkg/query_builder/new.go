package query_builder

import (
	"cloud.google.com/go/spanner/spansql"
	"sync"
)

func New() *QueryBuilder {
	return &QueryBuilder{
		query: &spansql.Query{
			Select: spansql.Select{},
		},
		mutex:    new(sync.RWMutex),
		distinct: false,
		_select:  make([]spansql.Expr, 0),
		table:    make([]spansql.SelectFromTable, 0),
		join:     make([]SelectFromJoin, 0),
		where:    make([]spansql.BoolExpr, 0),
		order:    make([]spansql.Order, 0),
		limit:    nil,
		offset:   nil,
	}
}

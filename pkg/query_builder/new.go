package query_builder

import (
	"cloud.google.com/go/spanner/spansql"
	"sync"
)

// New initalize/setup a new query builder
//
//	func main(){
//		var queryBuilder = query_builder.New()
//	}
func New() IQueryBuilder {
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
		param:    make(map[string]interface{}, 0),
	}
}

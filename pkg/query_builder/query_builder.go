package query_builder

import (
	"cloud.google.com/go/spanner/spansql"
	"sync"
)

type QueryBuilder struct {
	query *spansql.Query
	mutex *sync.RWMutex

	distinct      bool
	_select       []spansql.Expr
	table         []spansql.SelectFromTable
	join          []SelectFromJoin
	where         []spansql.BoolExpr
	order         []spansql.Order
	limit, offset spansql.LiteralOrParam
}

type IQueryBuilder interface {
	Distinct() *QueryBuilder
	Table(table spansql.SelectFromTable) *QueryBuilder
	Select(_select []spansql.Expr) *QueryBuilder
	Join(join SelectFromJoin) *QueryBuilder
	Where(expr spansql.BoolExpr) *QueryBuilder
	Order(order spansql.Order) *QueryBuilder
	Limit(limit spansql.LiteralOrParam) *QueryBuilder
	Offset(offset spansql.LiteralOrParam) *QueryBuilder
	Query() *spansql.Query
	SQL() string
}

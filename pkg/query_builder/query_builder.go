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
	param         map[string]interface{}
}

type IQueryBuilder interface {
	Distinct() IQueryBuilder
	Table(table spansql.SelectFromTable) IQueryBuilder
	Select(_select []spansql.Expr) IQueryBuilder
	Join(join SelectFromJoin) IQueryBuilder
	Where(expr spansql.BoolExpr) IQueryBuilder
	Order(order spansql.Order) IQueryBuilder
	Limit(limit spansql.LiteralOrParam) IQueryBuilder
	Offset(offset spansql.LiteralOrParam) IQueryBuilder
	Param(name string, value interface{}) spansql.Param
	Query() *spansql.Query
	SQL() (string, map[string]interface{})
}

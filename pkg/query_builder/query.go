package query_builder

import "cloud.google.com/go/spanner/spansql"

func (queryBuilder *QueryBuilder) Query() *spansql.Query {
	var distinct = queryBuilder.distinctResolve()
	var list = queryBuilder.selectResolve()
	var from = queryBuilder.selectFromResolve()
	var where = queryBuilder.whereResolve()
	var order = queryBuilder.orderResolve()
	var limit = queryBuilder.limitResolve()
	var offset = queryBuilder.offsetResolve()

	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.query.Select.Distinct = distinct
	queryBuilder.query.Select.List = list
	queryBuilder.query.Select.From = from
	queryBuilder.query.Select.Where = where
	queryBuilder.query.Order = order
	queryBuilder.query.Limit = limit
	queryBuilder.query.Offset = offset

	return queryBuilder.query
}

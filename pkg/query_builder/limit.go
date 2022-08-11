package query_builder

import "cloud.google.com/go/spanner/spansql"

func (queryBuilder *QueryBuilder) Limit(limit spansql.LiteralOrParam) *QueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.limit = limit
	return queryBuilder
}

func (queryBuilder *QueryBuilder) limitResolve() spansql.LiteralOrParam {
	queryBuilder.mutex.RLock()
	queryBuilder.mutex.RUnlock()

	return queryBuilder.limit
}

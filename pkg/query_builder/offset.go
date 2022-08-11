package query_builder

import "cloud.google.com/go/spanner/spansql"

func (queryBuilder *QueryBuilder) Offset(offset spansql.LiteralOrParam) *QueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.offset = offset
	return queryBuilder
}

func (queryBuilder *QueryBuilder) offsetResolve() spansql.LiteralOrParam {
	queryBuilder.mutex.RLock()
	queryBuilder.mutex.RUnlock()

	return queryBuilder.offset
}

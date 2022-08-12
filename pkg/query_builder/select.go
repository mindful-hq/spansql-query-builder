package query_builder

import "cloud.google.com/go/spanner/spansql"

func (queryBuilder *QueryBuilder) Select(_select []spansql.Expr) IQueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder._select = append(queryBuilder._select, _select...)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) selectResolve() []spansql.Expr {
	queryBuilder.mutex.RLock()
	queryBuilder.mutex.RUnlock()

	if len(queryBuilder._select) == 0 {
		return nil
	}

	return queryBuilder._select
}

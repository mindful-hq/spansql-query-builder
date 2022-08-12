package query_builder

func (queryBuilder *QueryBuilder) Distinct() IQueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.distinct = true
	return queryBuilder
}

func (queryBuilder *QueryBuilder) distinctResolve() bool {
	queryBuilder.mutex.RLock()
	queryBuilder.mutex.RUnlock()

	return queryBuilder.distinct
}

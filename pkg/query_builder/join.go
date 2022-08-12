package query_builder

func (queryBuilder *QueryBuilder) Join(join SelectFromJoin) IQueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.join = append(queryBuilder.join, join)
	return queryBuilder
}

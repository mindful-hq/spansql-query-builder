package query_builder

func (queryBuilder *QueryBuilder) Join(join SelectFromJoin) *QueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.join = append(queryBuilder.join, join)
	return queryBuilder
}

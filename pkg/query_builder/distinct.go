package query_builder

// Distinct discards duplicate rows and returns only the remaining rows.
// For more informations please see [Google Standard SQL distinct specification].
//
//	func main(){
//		var queryBuilder = query_builder.New().
//			Distinct()
//
//		// SELECT DISTINCT
//	}
//
// [Google Standard SQL distinct specification]: https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax
func (queryBuilder *QueryBuilder) Distinct() IQueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.distinct = true
	return queryBuilder
}

// distinctResolve resolves the spansql.Query.Select.Distinct field.
func (queryBuilder *QueryBuilder) distinctResolve() bool {
	queryBuilder.mutex.RLock()
	queryBuilder.mutex.RUnlock()

	return queryBuilder.distinct
}

package query_builder

import "cloud.google.com/go/spanner/spansql"

// Limit specifies a non-negative count of type INT64, and no more than count rows will be returned.
// For more informations please see [Google Standard SQL limit specification].
//
//	func main(){
//		var queryBuilder = query_builder.New().
//			Select([]spansql.Expr{
//				spansql.ID("Id"),
//			}).
//			Table(spansql.SelectFromTable{
//				Table: "Todos",
//			}).
//			Limit(spansql.IntegerLiteral(10))
//
//		// SELECT Id FROM Todos LIMIT 10
//	}
//
// [Google Standard SQL limit specification]: https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#limit_and_offset_clause
func (queryBuilder *QueryBuilder) Limit(limit spansql.LiteralOrParam) IQueryBuilder {
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

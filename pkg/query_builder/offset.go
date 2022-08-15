package query_builder

import "cloud.google.com/go/spanner/spansql"

// Offset specifies a non-negative number of rows to skip before applying Limit.
// Limit works without Offset, while Offset gets ignored without Limit
// For more informations please see [Google Standard SQL offset specification].
//
//	func main(){
//		var queryBuilder = query_builder.New().
//			Select([]spansql.Expr{
//				spansql.ID("Id"),
//			}).
//			Table(spansql.SelectFromTable{
//				Table: "Todos",
//			}).
//			Limit(spansql.IntegerLiteral(10)).
//			Offset(spansql.IntegerLiteral(10))
//
//		// SELECT Id FROM Todos LIMIT 10 OFFSET 10
//	}
//
// [Google Standard SQL offset specification]: https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#limit_and_offset_clause
func (queryBuilder *QueryBuilder) Offset(offset spansql.LiteralOrParam) IQueryBuilder {
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

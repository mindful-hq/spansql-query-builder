package query_builder

// Join merges two from_items so that the SELECT clause can query them as one source.
// For more informations please see [Google Standard SQL join specification].
//
//	func main(){
//		var queryBuilder = query_builder.New().
//			Select([]spansql.Expr{
//				spansql.ID("Todos.Id"),
//				spansql.ID("Places.Name"),
//			}).
//			Table(spansql.SelectFromTable{
//				Table: "Todos",
//			}).
//			Join(query_builder.SelectFromJoin{
//				Type: spansql.InnerJoin,
//				LHS:  spansql.SelectFromTable{Table: "Places"},
//				On:   spansql.ComparisonOp{Op: spansql.Eq, LHS: spansql.PathExp{"Todos", "Id"}, RHS: spansql.PathExp{"Places.TodoId"}},
//	   		})
//
//		// SELECT Todos.Id, Places.Name FROM Todos INNER JOIN Places ON Todos.Id = Places.TodoId
//	}
//
// [Google Standard SQL join specification]: https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#join_types
func (queryBuilder *QueryBuilder) Join(join SelectFromJoin) IQueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.join = append(queryBuilder.join, join)
	return queryBuilder
}

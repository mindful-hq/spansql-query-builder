package query_builder

import "cloud.google.com/go/spanner/spansql"

// Table indicates the table or tables from which to retrieve rows.
// For more informations please see [Google Standard SQL from specification].
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
//			Table(spansql.SelectFromTable{
//				Table: "Places",
//			})
//
//		// SELECT Todos.Id, Places.Name FROM Todos, Places
//	}
//
// [Google Standard SQL from specification]: https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#from_clause
func (queryBuilder *QueryBuilder) Table(table spansql.SelectFromTable) IQueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.table = append(queryBuilder.table, table)
	return queryBuilder
}

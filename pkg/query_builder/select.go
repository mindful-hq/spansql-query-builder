package query_builder

import "cloud.google.com/go/spanner/spansql"

// Select defines the columns that the query will return.
// For more informations please see [Google Standard SQL select specification].
//
//	func main(){
//		var queryBuilder = query_builder.New().
//			Select([]spansql.Expr{
//				spansql.ID("Id"),
//				spansql.ID("Name"),
//				spansql.ID("Content"),
//			}).
//			Table(spansql.SelectFromTable{
//				Table: "Todos",
//			})
//
//		// SELECT Id FROM Todos
//	}
//
// [Google Standard SQL select specification]: https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#select_list
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

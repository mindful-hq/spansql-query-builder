package query_builder

import "cloud.google.com/go/spanner/spansql"

// Param sets and store query parameters.
//
//	func main(){
//		var queryBuilder = query_builder.New().
//			Select([]spansql.Expr{
//				spansql.ID("Id"),
//			}).
//			Table(spansql.SelectFromTable{
//				Table: "Todos",
//			}).
//			Where(spansql.ComparisonOp{Op: spansql.Eq, LHS: spansql.ID("Id"), RHS: query_builder.Param("Id", spansql.IntegerLiteral(1))})
//
//		// SELECT Id FROM Todos WHERE Id = @Id
//	}
func (queryBuilder *QueryBuilder) Param(name string, value interface{}) spansql.Param {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.param[name] = value
	return spansql.Param(name)
}

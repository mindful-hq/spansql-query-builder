package query_builder

import (
	"cloud.google.com/go/spanner/spansql"
	"encoding/hex"
	"github.com/google/uuid"
)

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

// ParamUnique sets and store query parameters with a unique name (uuid suffix).
//
//	func main(){
//		var queryBuilder = query_builder.New().
//			Select([]spansql.Expr{
//				spansql.ID("Id"),
//			}).
//			Table(spansql.SelectFromTable{
//				Table: "Todos",
//			}).
//			Where(spansql.ComparisonOp{Op: spansql.Eq, LHS: spansql.ID("Id"), RHS: query_builder.ParamUnique("Id", spansql.IntegerLiteral(1))})
//
//		// SELECT Id FROM Todos WHERE Id = @Id92730f9147e24499b4b4b02a0939bc06
//	}
func (queryBuilder *QueryBuilder) ParamUnique(name string, value interface{}) spansql.Param {
	var uniqueId = uuid.New()
	return queryBuilder.Param(name+hex.EncodeToString(uniqueId[:]), value)
}

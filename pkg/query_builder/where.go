package query_builder

import "cloud.google.com/go/spanner/spansql"

// Where filters the results of the FROM clause.
// For more informations please see [Google Standard SQL where specification].
//
//	func main(){
//		var queryBuilder = query_builder.New().
//			Select([]spansql.Expr{
//				spansql.ID("Todos.Id"),
//				spansql.ID("Places.Name"),
//			}).
//			Where(spansql.ComparisonOp{Op: spansql.Eq, LHS: spansql.ID("Todos.Id"), RHS: qb.Param("TodosId", spansql.IntegerLiteral(1))})
//
//		// SELECT Todos.Id, Places.Name FROM Todos, Places WHERE Todos.Id = @TodosId
//	}
//
// [Google Standard SQL where specification]: https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#where_clause
func (queryBuilder *QueryBuilder) Where(expr spansql.BoolExpr) IQueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.where = append(queryBuilder.where, expr)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) whereResolve() spansql.BoolExpr {
	queryBuilder.mutex.RLock()
	defer queryBuilder.mutex.RUnlock()

	if len(queryBuilder.where) == 0 {
		return nil
	}

	if len(queryBuilder.where) == 1 {
		return queryBuilder.where[0]
	}

	op := &spansql.LogicalOp{
		Op: spansql.And,
	}

	for _, expr := range queryBuilder.where {
		if op.LHS == nil {
			op.LHS = expr
			continue
		}

		if op.RHS == nil {
			op.RHS = expr
			continue
		}

		op.RHS = &spansql.LogicalOp{
			Op:  spansql.And,
			LHS: op.RHS,
			RHS: expr,
		}
	}

	return *op
}

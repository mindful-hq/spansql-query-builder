package query_builder

import "cloud.google.com/go/spanner/spansql"

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

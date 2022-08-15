package query_builder

import "cloud.google.com/go/spanner/spansql"

// Order specifies a column or expression as the sort criterion for the result set.
// For more informations please see [Google Standard SQL order specification].
//
//	func main(){
//		var queryBuilder = query_builder.New().
//			Select([]spansql.Expr{
//				spansql.ID("Id"),
//			}).
//			Table(spansql.SelectFromTable{
//				Table: "Todos",
//			}).
//			Order(spansql.Order{Expr: spansql.ID("Name"), Desc: true})
//
//		// SELECT Id FROM Todos ORDER BY Name
//	}
//
// [Google Standard SQL order specification]: https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#order_by_clause
func (queryBuilder *QueryBuilder) Order(order spansql.Order) IQueryBuilder {
	queryBuilder.mutex.Lock()
	defer queryBuilder.mutex.Unlock()

	queryBuilder.order = append(queryBuilder.order, order)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) orderResolve() []spansql.Order {
	queryBuilder.mutex.RLock()
	defer queryBuilder.mutex.RUnlock()

	if len(queryBuilder.order) == 0 {
		return nil
	}

	return queryBuilder.order
}

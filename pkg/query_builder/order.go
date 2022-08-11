package query_builder

import "cloud.google.com/go/spanner/spansql"

func (queryBuilder *QueryBuilder) Order(order spansql.Order) *QueryBuilder {
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

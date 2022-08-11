package query_builder

import "cloud.google.com/go/spanner/spansql"

func (queryBuilder *QueryBuilder) selectFromResolve() []spansql.SelectFrom {
	queryBuilder.mutex.RLock()
	queryBuilder.mutex.RUnlock()

	if len(queryBuilder.table) == 0 && len(queryBuilder.join) == 0 {
		return nil
	}

	var selectFrom = make([]spansql.SelectFrom, 0)

	for _, t := range queryBuilder.table {
		selectFrom = append(selectFrom, t)
	}

	for _, j := range queryBuilder.join {
		selectFrom = append(selectFrom, j)
	}

	return selectFrom
}

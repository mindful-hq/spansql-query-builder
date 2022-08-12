package query_builder

import (
	"regexp"
)

// SQL returns the query sql statement and all named paramters.
func (queryBuilder *QueryBuilder) SQL() (string, map[string]interface{}) {
	var query = queryBuilder.Query()

	queryBuilder.mutex.RLock()
	var param = queryBuilder.param
	queryBuilder.mutex.RUnlock()

	if len(queryBuilder.join) > 0 {
		regex := regexp.MustCompile(`(,\s)(\w+\sJOIN)`)
		return regex.ReplaceAllString(query.SQL(), " $2"), param
	}

	return query.SQL(), param
}

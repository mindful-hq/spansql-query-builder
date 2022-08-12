package query_builder

import (
	"regexp"
)

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

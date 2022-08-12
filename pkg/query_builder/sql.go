package query_builder

import (
	"regexp"
)

func (queryBuilder *QueryBuilder) SQL() string {
	var query = queryBuilder.Query()

	if len(queryBuilder.join) > 0 {
		regex := regexp.MustCompile(`(,\s)(\w+\sJOIN)`)
		return regex.ReplaceAllString(query.SQL(), " $2")
	}

	return query.SQL()
}

package query_builder

import (
	"cloud.google.com/go/spanner/spansql"
	"github.com/mindful-hq/spansql-query-builder/pkg/query_builder"
	"regexp"
)

func (queryBuilder *QueryBuilder) SQL() string {
	var sql = query_builder.New().
		Select([]spansql.Expr{
			spansql.ID("Todos.Id"),
			spansql.ID("Places.Name"),
		}).
		Table(spansql.SelectFromTable{
			Table: "Todos",
		}).
		Join(query_builder.SelectFromJoin{
			Type: spansql.InnerJoin,
			LHS:  spansql.SelectFromTable{Table: "Places"},
			On:   spansql.ComparisonOp{Op: spansql.Eq, LHS: spansql.PathExp{"Todos", "Id"}, RHS: spansql.PathExp{"Places.TodoId"}},
		}).
		Where(spansql.ComparisonOp{Op: spansql.Eq, LHS: spansql.ID("Id"), RHS: spansql.IntegerLiteral(1)}).
		Where(spansql.ComparisonOp{Op: spansql.Like, LHS: spansql.ID("Name"), RHS: spansql.StringLiteral("%test%")}).
		SQL()
	var query = queryBuilder.Query()

	if len(queryBuilder.join) > 0 {
		regex := regexp.MustCompile(`(,\s)(\w+\sJOIN)`)
		return regex.ReplaceAllString(query.SQL(), " $2")
	}

	return query.SQL()
}

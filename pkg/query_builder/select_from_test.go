package query_builder

import (
	"cloud.google.com/go/spanner/spansql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectFromResolve(t *testing.T) {
	var tests = []struct {
		name         string
		queryBuilder *QueryBuilder
		table        []spansql.SelectFromTable
		join         []SelectFromJoin
		expr         []spansql.SelectFrom
		sql          string
		err          assert.ErrorAssertionFunc
	}{
		{
			name:         "no table, no join",
			queryBuilder: New(),
			table:        nil,
			join:         nil,
			expr:         nil,
			sql:          "SELECT ",
			err:          assert.NoError,
		},
		{
			name:         "one table, no join",
			queryBuilder: New().Select([]spansql.Expr{spansql.Star}),
			table: []spansql.SelectFromTable{
				{Table: "Test1"},
			},
			join: nil,
			expr: []spansql.SelectFrom{
				spansql.SelectFromTable{Table: "Test1"},
			},
			sql: `SELECT * FROM Test1`,
			err: assert.NoError,
		},
		{
			name:         "two table, no join",
			queryBuilder: New().Select([]spansql.Expr{spansql.Star}),
			table: []spansql.SelectFromTable{
				{Table: "Test1"},
				{Table: "Test2"},
			},
			join: nil,
			expr: []spansql.SelectFrom{
				spansql.SelectFromTable{Table: "Test1"},
				spansql.SelectFromTable{Table: "Test2"},
			},
			sql: `SELECT * FROM Test1, Test2`,
			err: assert.NoError,
		},
		{
			name:         "two table, one join",
			queryBuilder: New().Select([]spansql.Expr{spansql.Star}),
			table: []spansql.SelectFromTable{
				{Table: "Test1"},
				{Table: "Test2"},
			},
			join: []SelectFromJoin{
				{
					Type: spansql.InnerJoin,
					LHS:  spansql.SelectFromTable{Table: "Test3"},
					On: spansql.ComparisonOp{
						Op:  spansql.Eq,
						LHS: spansql.ID("Id"),
						RHS: spansql.StringLiteral("test3"),
					},
				},
			},
			expr: []spansql.SelectFrom{
				spansql.SelectFromTable{Table: "Test1"},
				spansql.SelectFromTable{Table: "Test2"},
				SelectFromJoin{
					Type: spansql.InnerJoin,
					LHS:  spansql.SelectFromTable{Table: "Test3"},
					On: spansql.ComparisonOp{
						Op:  spansql.Eq,
						LHS: spansql.ID("Id"),
						RHS: spansql.StringLiteral("test3"),
					},
				},
			},
			sql: `SELECT * FROM Test1, Test2 INNER JOIN Test3 ON Id = "test3"`,
			err: assert.NoError,
		},
		{
			name:         "two table, two join",
			queryBuilder: New().Select([]spansql.Expr{spansql.Star}),
			table: []spansql.SelectFromTable{
				{Table: "Test1"},
				{Table: "Test2"},
			},
			join: []SelectFromJoin{
				{
					Type: spansql.InnerJoin,
					LHS:  spansql.SelectFromTable{Table: "Test3"},
					On: spansql.ComparisonOp{
						Op:  spansql.Eq,
						LHS: spansql.ID("Id"),
						RHS: spansql.StringLiteral("test3"),
					},
				},
				{
					Type: spansql.InnerJoin,
					LHS:  spansql.SelectFromTable{Table: "Test4"},
					On: spansql.ComparisonOp{
						Op:  spansql.Eq,
						LHS: spansql.ID("Id"),
						RHS: spansql.StringLiteral("test4"),
					},
				},
			},
			expr: []spansql.SelectFrom{
				spansql.SelectFromTable{Table: "Test1"},
				spansql.SelectFromTable{Table: "Test2"},
				SelectFromJoin{
					Type: spansql.InnerJoin,
					LHS:  spansql.SelectFromTable{Table: "Test3"},
					On: spansql.ComparisonOp{
						Op:  spansql.Eq,
						LHS: spansql.ID("Id"),
						RHS: spansql.StringLiteral("test3"),
					},
				},
				SelectFromJoin{
					Type: spansql.InnerJoin,
					LHS:  spansql.SelectFromTable{Table: "Test4"},
					On: spansql.ComparisonOp{
						Op:  spansql.Eq,
						LHS: spansql.ID("Id"),
						RHS: spansql.StringLiteral("test4"),
					},
				},
			},
			sql: `SELECT * FROM Test1, Test2 INNER JOIN Test3 ON Id = "test3" INNER JOIN Test4 ON Id = "test4"`,
			err: assert.NoError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.table != nil {
				for _, selectFrom := range test.table {
					test.queryBuilder.Table(selectFrom)
				}
			}

			if test.join != nil {
				for _, selectFrom := range test.join {
					test.queryBuilder.Join(selectFrom)
				}
			}

			var expr = test.queryBuilder.selectFromResolve()
			assert.Equal(t, test.expr, expr)

			if expr != nil {
				assert.Equal(t, test.sql, test.queryBuilder.SQL())
			}
		})
	}
}

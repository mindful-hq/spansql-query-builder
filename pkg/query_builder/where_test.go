package query_builder

import (
	"cloud.google.com/go/spanner/spansql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func comparisonOp(name string) spansql.ComparisonOp {
	return spansql.ComparisonOp{
		Op:   spansql.Eq,
		LHS:  spansql.ID("Id"),
		RHS:  spansql.StringLiteral(fmt.Sprintf("test-%s", name)),
		RHS2: nil,
	}
}

func logicalOp(name string) spansql.LogicalOp {
	return spansql.LogicalOp{
		Op: spansql.And,
		LHS: spansql.ComparisonOp{
			Op:   spansql.Eq,
			LHS:  spansql.ID("Id"),
			RHS:  spansql.StringLiteral(fmt.Sprintf("test-%s-1", name)),
			RHS2: nil,
		},
		RHS: spansql.ComparisonOp{
			Op:   spansql.Eq,
			LHS:  spansql.ID("Id"),
			RHS:  spansql.StringLiteral(fmt.Sprintf("test-%s-2", name)),
			RHS2: nil,
		},
	}
}

func logicalOpParen(name string) spansql.Paren {
	return spansql.Paren{
		Expr: spansql.LogicalOp{
			Op: spansql.Or,
			LHS: spansql.ComparisonOp{
				Op:   spansql.Eq,
				LHS:  spansql.ID("Id"),
				RHS:  spansql.StringLiteral(fmt.Sprintf("test-%s-1", name)),
				RHS2: nil,
			},
			RHS: spansql.ComparisonOp{
				Op:   spansql.Eq,
				LHS:  spansql.ID("Id"),
				RHS:  spansql.StringLiteral(fmt.Sprintf("test-%s-2", name)),
				RHS2: nil,
			},
		},
	}
}

func TestWhereResolve(t *testing.T) {
	var tests = []struct {
		name         string
		queryBuilder IQueryBuilder
		where        []spansql.BoolExpr
		expr         spansql.BoolExpr
		sql          string
	}{
		{
			name:         "no where",
			queryBuilder: New(),
			where:        []spansql.BoolExpr{},
			expr:         nil,
			sql:          "",
		},
		{
			name:         "one where",
			queryBuilder: New(),
			where: []spansql.BoolExpr{
				comparisonOp("1"),
			},
			expr: spansql.ComparisonOp{
				Op:   spansql.Eq,
				LHS:  spansql.ID("Id"),
				RHS:  spansql.StringLiteral("test-1"),
				RHS2: nil,
			},
			sql: `Id = "test-1"`,
		},
		{
			name:         "two where",
			queryBuilder: New(),
			where: []spansql.BoolExpr{
				comparisonOp("1"),
				comparisonOp("2"),
			},
			expr: spansql.LogicalOp{
				Op: spansql.And,
				LHS: spansql.ComparisonOp{
					Op:   spansql.Eq,
					LHS:  spansql.ID("Id"),
					RHS:  spansql.StringLiteral("test-1"),
					RHS2: nil,
				},
				RHS: spansql.ComparisonOp{
					Op:   spansql.Eq,
					LHS:  spansql.ID("Id"),
					RHS:  spansql.StringLiteral("test-2"),
					RHS2: nil,
				},
			},
			sql: `Id = "test-1" AND Id = "test-2"`,
		},
		{
			name:         "three where",
			queryBuilder: New(),
			where: []spansql.BoolExpr{
				comparisonOp("1"),
				comparisonOp("2"),
				comparisonOp("3"),
			},
			expr: spansql.LogicalOp{
				Op: spansql.And,
				LHS: spansql.ComparisonOp{
					Op:   spansql.Eq,
					LHS:  spansql.ID("Id"),
					RHS:  spansql.StringLiteral("test-1"),
					RHS2: nil,
				},
				RHS: &spansql.LogicalOp{
					Op: spansql.And,
					LHS: spansql.ComparisonOp{
						Op:   spansql.Eq,
						LHS:  spansql.ID("Id"),
						RHS:  spansql.StringLiteral("test-2"),
						RHS2: nil,
					},
					RHS: spansql.ComparisonOp{
						Op:   spansql.Eq,
						LHS:  spansql.ID("Id"),
						RHS:  spansql.StringLiteral("test-3"),
						RHS2: nil,
					},
				},
			},
			sql: `Id = "test-1" AND Id = "test-2" AND Id = "test-3"`,
		},
		{
			name:         "four where",
			queryBuilder: New(),
			where: []spansql.BoolExpr{
				comparisonOp("1"),
				comparisonOp("2"),
				comparisonOp("3"),
				comparisonOp("4"),
			},
			expr: spansql.LogicalOp{
				Op: spansql.And,
				LHS: spansql.ComparisonOp{
					Op:   spansql.Eq,
					LHS:  spansql.ID("Id"),
					RHS:  spansql.StringLiteral("test-1"),
					RHS2: nil,
				},
				RHS: &spansql.LogicalOp{
					Op: spansql.And,
					LHS: &spansql.LogicalOp{
						Op: spansql.And,
						LHS: spansql.ComparisonOp{
							Op:   spansql.Eq,
							LHS:  spansql.ID("Id"),
							RHS:  spansql.StringLiteral("test-2"),
							RHS2: nil,
						},
						RHS: spansql.ComparisonOp{
							Op:   spansql.Eq,
							LHS:  spansql.ID("Id"),
							RHS:  spansql.StringLiteral("test-3"),
							RHS2: nil,
						},
					},
					RHS: spansql.ComparisonOp{
						Op:   spansql.Eq,
						LHS:  spansql.ID("Id"),
						RHS:  spansql.StringLiteral("test-4"),
						RHS2: nil,
					},
				},
			},
			sql: `Id = "test-1" AND Id = "test-2" AND Id = "test-3" AND Id = "test-4"`,
		},
		{
			name:         "four where with logicalop",
			queryBuilder: New(),
			where: []spansql.BoolExpr{
				comparisonOp("1"),
				comparisonOp("2"),
				comparisonOp("3"),
				logicalOp("4"),
			},
			expr: spansql.LogicalOp{
				Op: spansql.And,
				LHS: spansql.ComparisonOp{
					Op:   spansql.Eq,
					LHS:  spansql.ID("Id"),
					RHS:  spansql.StringLiteral("test-1"),
					RHS2: nil,
				},
				RHS: &spansql.LogicalOp{
					Op: spansql.And,
					LHS: &spansql.LogicalOp{
						Op: spansql.And,
						LHS: spansql.ComparisonOp{
							Op:   spansql.Eq,
							LHS:  spansql.ID("Id"),
							RHS:  spansql.StringLiteral("test-2"),
							RHS2: nil,
						},
						RHS: spansql.ComparisonOp{
							Op:   spansql.Eq,
							LHS:  spansql.ID("Id"),
							RHS:  spansql.StringLiteral("test-3"),
							RHS2: nil,
						},
					},
					RHS: spansql.LogicalOp{
						Op: spansql.And,
						LHS: spansql.ComparisonOp{
							Op:   spansql.Eq,
							LHS:  spansql.ID("Id"),
							RHS:  spansql.StringLiteral("test-4-1"),
							RHS2: nil,
						},
						RHS: spansql.ComparisonOp{
							Op:   spansql.Eq,
							LHS:  spansql.ID("Id"),
							RHS:  spansql.StringLiteral("test-4-2"),
							RHS2: nil,
						},
					},
				},
			},
			sql: `Id = "test-1" AND Id = "test-2" AND Id = "test-3" AND Id = "test-4-1" AND Id = "test-4-2"`,
		},
		{
			name:         "four where with logicalop paren",
			queryBuilder: New(),
			where: []spansql.BoolExpr{
				comparisonOp("1"),
				comparisonOp("2"),
				comparisonOp("3"),
				logicalOpParen("4"),
			},
			expr: spansql.LogicalOp{
				Op: spansql.And,
				LHS: spansql.ComparisonOp{
					Op:   spansql.Eq,
					LHS:  spansql.ID("Id"),
					RHS:  spansql.StringLiteral("test-1"),
					RHS2: nil,
				},
				RHS: &spansql.LogicalOp{
					Op: spansql.And,
					LHS: &spansql.LogicalOp{
						Op: spansql.And,
						LHS: spansql.ComparisonOp{
							Op:   spansql.Eq,
							LHS:  spansql.ID("Id"),
							RHS:  spansql.StringLiteral("test-2"),
							RHS2: nil,
						},
						RHS: spansql.ComparisonOp{
							Op:   spansql.Eq,
							LHS:  spansql.ID("Id"),
							RHS:  spansql.StringLiteral("test-3"),
							RHS2: nil,
						},
					},
					RHS: spansql.Paren{
						Expr: spansql.LogicalOp{
							Op: spansql.Or,
							LHS: spansql.ComparisonOp{
								Op:   spansql.Eq,
								LHS:  spansql.ID("Id"),
								RHS:  spansql.StringLiteral("test-4-1"),
								RHS2: nil,
							},
							RHS: spansql.ComparisonOp{
								Op:   spansql.Eq,
								LHS:  spansql.ID("Id"),
								RHS:  spansql.StringLiteral("test-4-2"),
								RHS2: nil,
							},
						},
					},
				},
			},
			sql: `Id = "test-1" AND Id = "test-2" AND Id = "test-3" AND (Id = "test-4-1" OR Id = "test-4-2")`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expr := range test.where {
				test.queryBuilder.Where(expr)
			}
			var expr = test.queryBuilder.(*QueryBuilder).whereResolve()
			assert.Equal(t, test.expr, expr)

			if expr != nil {
				assert.Equal(t, test.sql, expr.SQL())
			}
		})
	}
}

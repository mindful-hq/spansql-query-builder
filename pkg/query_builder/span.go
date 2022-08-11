package query_builder

import (
	"cloud.google.com/go/spanner/spansql"
	"strings"
)

// SelectFromJoin is a custom SelectFromJoin implementation
// It does not include the original LHS because it conflicts with the SelectFromTable implementation
// RHS is renamed to LHS
type SelectFromJoin struct {
	spansql.SelectFrom
	Type spansql.JoinType
	LHS  spansql.SelectFrom

	// Join condition.
	// At most one of {On,Using} may be set.
	On    spansql.BoolExpr
	Using []spansql.ID

	// Hints are suggestions for how to evaluate a join.
	// https://cloud.google.com/spanner/docs/query-syntax#join-hints
	Hints map[string]string
}

type selectFromEmpty struct {
	spansql.SelectFrom
}

func (selectFromEmpty selectFromEmpty) SQL() string {
	return ""
}

func (selectFromJoin SelectFromJoin) SQL() string {
	var clone = spansql.SelectFromJoin{
		Type:  selectFromJoin.Type,
		LHS:   selectFromEmpty{},
		RHS:   selectFromJoin.LHS,
		On:    selectFromJoin.On,
		Using: selectFromJoin.Using,
		Hints: selectFromJoin.Hints,
	}

	return strings.TrimSpace(clone.SQL())
}

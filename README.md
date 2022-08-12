<img width="100" src="https://raw.githubusercontent.com/mindful-hq/spansql-query-builder/main/assets/images/spanner.png" alt="spanner-logo">

# Google Spanner SpanSQL Query Builder

![License](https://img.shields.io/npm/l/@ls-lint/ls-lint.svg?sanitize=true)
[![Build Status](https://drone.mindful.com/api/badges/mindful-hq/spansql-query-builder/status.svg)](https://drone.mindful.com/mindful-hq/spansql-query-builder)
[![Go Reference](https://pkg.go.dev/badge/cloud.google.com/go/spanner.svg)](https://pkg.go.dev/github.com/mindful-hq/spansql-query-builder)
[![Go Report Card](https://goreportcard.com/badge/github.com/mindful-hq/spansql-query-builder)](https://goreportcard.com/report/github.com/mindful-hq/spansql-query-builder)

High level lightweight [spansql](https://pkg.go.dev/cloud.google.com/go/spanner/spansql) query builder.
This is not an officially supported Google product.

## Example

```go
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
```

```sql 
SELECT Todos.Id, Places.Name FROM Todos INNER JOIN Places ON Todos.Id = Places.TodoId WHERE Id = 1 AND Name LIKE "%test%"
```

Play with it: [Go Playground](https://go.dev/play/p/Ih2IOS8UCJn)
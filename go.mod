module github.com/mindful-hq/spansql-query-builder

go 1.19

replace golang.org/x/tools => golang.org/x/tools v0.1.12 // https://github.com/bazelbuild/rules_go/issues/3230#issuecomment-1216728711

require (
	cloud.google.com/go/spanner v1.36.0
	github.com/stretchr/testify v1.8.0
)

require (
	cloud.google.com/go v0.103.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220811182439-13a9a731de15 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/tools v0.1.12 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

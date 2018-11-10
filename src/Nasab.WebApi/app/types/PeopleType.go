package types

import "github.com/graphql-go/graphql"

var PeopleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "People",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

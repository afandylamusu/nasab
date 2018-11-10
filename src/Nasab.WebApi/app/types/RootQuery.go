package types

import (
	"github.com/afandylamusu/nasab/src/Nasab.WebApi/app/db"
	"github.com/graphql-go/graphql"
)

// RootQuery for graphql endpoint
var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* POST (read) peopeles by paging
			   query body:

			   {
				   peoples(page:0, page_size:25)
				   	{
						names,
						birthOfDate,
						birthOfPlace
					}
				}

			*/
			"peoples": &graphql.Field{
				Type:        graphql.NewList(PeopleType),
				Description: "Get People list",
				Args: graphql.FieldConfigArgument{
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"page_size": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// page, ok := p.Args["page"].(int)
					// pageSize, ok := p.Args["page_size"].(int)

					var peoples []db.People

					return peoples, nil
				},
			},
		},
	},
)

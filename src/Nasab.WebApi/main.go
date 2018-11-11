package main

import (
	"github.com/afandylamusu/nasab/src/Nasab.WebApi/app/db"
	"github.com/graphql-go/graphql"
	"github.com/kataras/iris"

	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// NewUser will call insert User into table
func NewUser(people *db.People) {
	connection := db.Connect()

	// Create
	connection.Create(people)
}

var peoples []db.People

var peopleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "People",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var queryType = graphql.NewObject(
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
				Type:        graphql.NewList(peopleType),
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

					return peoples, nil
				},
			},
		},
	},
)

/**
 * sfsdfds
 * dsdfsdf
 * sdfds
 */
func main() {
	app := iris.Default()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", graphiql.ServeGraphiQL)

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// Method:   POST
	// Resource: /api/graphql This is to handle graphql query and mutations
	app.Handle("POST", "/graphql", func(ctx context.Context) {

		// params := &User{}
		// err := ctx.ReadJSON(params)
		// if err != nil {
		// 	ctx.JSON(context.Map{"response": err.Error()})
		// } else {
		// 	params.LastUpdate = time.Now()
		// 	err := c.Insert(params)
		// 	if err != nil {
		// 		ctx.JSON(context.Map{"response": err.Error()})
		// 	} else {
		// 		fmt.Println("Successfully inserted into database")
		// 		result := User{}
		// 		err = c.Find(bson.M{"msisdn": params.Msisdn}).One(&result)
		// 		if err != nil {
		// 			ctx.JSON(context.Map{"response": err.Error()})
		// 		}
		// 		ctx.JSON(context.Map{"response": "User succesfully created", "message": result})
		// 	}
		// }

	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

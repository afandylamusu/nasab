package main

import (
	"encoding/json"
	"net/http"

	"github.com/afandylamusu/nasab/src/Nasab.WebApi/app/db"
	"github.com/afandylamusu/nasab/src/Nasab.WebApi/app/types"

	"github.com/graphql-go/graphql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"github.com/afandylamusu/go-graphiql"
)

// NewUser will call insert User into table
func NewUser(people *db.People) {
	connection := db.Connect()

	// Create
	connection.Create(people)
}

func serveGraphQL(s graphql.Schema, w http.ResponseWriter, r *http.Request) {

	sendError := func(err error) {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	req := &graphiql.Request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		sendError(err)
		return
	}

	res := graphql.Do(graphql.Params{
		Schema:        s,
		RequestString: req.Query,
	})

	if err := json.NewEncoder(w).Encode(res); err != nil {
		sendError(err)
	}
}

/**
 * sfsdfds
 * dsdfsdf
 * sdfds
 */
func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		graphiql.ServeGraphiQL(ctx.ResponseWriter(), ctx.Request())
	})

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
		schema, err := graphql.NewSchema(graphql.SchemaConfig{
			Query: types.RootQuery,
			// Mutation: types.RootMutation,
		})

		if err != nil {
			panic(err)
		}

		serveGraphQL(schema, ctx.ResponseWriter(), ctx.Request())

	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8086"), iris.WithoutServerError(iris.ErrServerClosed))
}

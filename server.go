package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/sudo-NithishKarthik/hackernews/graph"
	"github.com/sudo-NithishKarthik/hackernews/graph/generated"
	"github.com/sudo-NithishKarthik/hackernews/internal/pkg/db/mysql"
	"github.com/sudo-NithishKarthik/hackernews/internal/auth"
)

const defaultPort = "8080"
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
  router := chi.NewRouter()
  router.Use(auth.Middleware())
  database.InitDB()
	database.Migrate()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

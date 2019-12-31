package main

import (
	"flag"
	"github.com/deni1688/bookingql/graphql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	setServerENV()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := graphql.Config{Resolvers: &graphql.Resolver{}}
	queryHandler := handler.GraphQL(graphql.NewExecutableSchema(config))

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", graphql.DataloaderMiddleware(queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setServerENV() {
	serverURL := flag.String("server", "https://my.fleetster.de", "Target fleetster server")
	dumpReq := flag.String("dump-req", "false", "Dumps the outgoing request to the terminal")
	flag.Parse()

	os.Setenv("SERVER", *serverURL)
	os.Setenv("DUMPREQ", *dumpReq)
}

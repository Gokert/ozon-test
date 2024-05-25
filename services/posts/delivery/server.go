package main

import (
	"net/http"
	"ozon-test/configs"
	"ozon-test/configs/logger"
	graph2 "ozon-test/services/posts/delivery/graph"
	"ozon-test/services/posts/usecase"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := defaultPort

	log := logger.GetLogger()
	err := configs.InitEnv()
	if err != nil {
		log.Errorf("Init env error: %s", err.Error())
		return
	}

	grpcCfg, err := configs.GetGrpcConfig()
	if err != nil {
		log.Errorf("failed to parse grpc configs file: %s", err.Error())
		return
	}

	psxCfg, err := configs.GetPostsPsxConfig()
	if err != nil {
		log.Errorf("Create psx config error: %s", err.Error())
		return
	}

	core, err := usecase.GetCore(grpcCfg, psxCfg, log)
	if err != nil {
		log.Errorf("Create core error: %s", err.Error())
		return
	}

	resolver := &graph2.Resolver{
		Core: core,
		Log:  log,
	}

	srv := handler.NewDefaultServer(graph2.NewExecutableSchema(graph2.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

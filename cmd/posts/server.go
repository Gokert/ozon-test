package main

import (
	"flag"
	"net/http"
	"os"
	"ozon-test/configs"
	"ozon-test/configs/logger"
	"ozon-test/pkg/middleware"
	graph2 "ozon-test/services/posts/delivery/graph"
	"ozon-test/services/posts/usecase"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	var core *usecase.Core

	log := logger.GetLogger()
	err := configs.InitEnv()
	if err != nil {
		log.Errorf("Init env error: %s", err.Error())
		return
	}

	port := os.Getenv("POSTS_APP_PORT")
	option := flag.String("database", "postgresql", "выбор БД для записи данных (postgresql | redis)")
	flag.Parse()

	grpcCfg, err := configs.GetGrpcConfig()
	if err != nil {
		log.Errorf("failed to parse grpc configs file: %s", err.Error())
		return
	}

	if *option == "postgresql" {
		psxCfg, err := configs.GetPostsPsxConfig()
		if err != nil {
			log.Errorf("Create psx config error: %s", err.Error())
			return
		}

		core, err = usecase.GetPostgresCore(grpcCfg, psxCfg, log)
		if err != nil {
			log.Errorf("Create core error: %s", err.Error())
			return
		}
	}

	if *option == "redis" {
		redisCfg, err := configs.GetRedisConfig()
		if err != nil {
			log.Errorf("Create redis config error: %s", err.Error())
			return
		}

		core, err = usecase.GetRedisCore(grpcCfg, redisCfg, log)
		if err != nil {
			log.Errorf("Create core error: %s", err.Error())
			return
		}
	}

	resolver := &graph2.Resolver{
		Core: core,
		Log:  log,
	}

	srv := handler.NewDefaultServer(graph2.NewExecutableSchema(graph2.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/api/v1/query"))
	http.Handle("/api/v1/query", middleware.AuthCheck(srv, core, log))

	log.Printf("Server Post with GraphQL running on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

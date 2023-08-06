package main

import (
	"log"
	"net/http"

	"github.com/jumpei00/grpc/grpcback/pkg/application"
	"github.com/jumpei00/grpc/grpcback/pkg/infrastructure"
	"github.com/jumpei00/grpc/grpcback/pkg/interfaces"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()

	mysql, err := infrastructure.NewMySQL()
	if err != nil {
		log.Fatalf("failed to connect mysql: %v", err)
	}

	// infrastructure setup
	todoRepository := infrastructure.NewTodoRepository(mysql)

	// application setup
	todoService := application.NewTodoApplication(todoRepository)

	// interface setup
	todoServer := interfaces.NewTodoServer(todoService)

	// handle server
	todoServer.HandleServer(mux)

	// cors
	corsConfig := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders: []string{
			"Accept-Encoding",
			"Content-Encoding",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
		},
	})
	corsHandler := corsConfig.Handler(h2c.NewHandler(mux, &http2.Server{}))

	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

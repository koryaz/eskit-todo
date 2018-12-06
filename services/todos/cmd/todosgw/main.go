package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/koryaz/eskit-todo/generated/grpc/go/todos"
	"google.golang.org/grpc"
)

var (
	endpoint = flag.String("grpc_endpoint", "localhost:9090", "Grpc Service")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := todos.RegisterTodoServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		log.Fatalf("starting gw failed : %v", err)
	}

	log.Printf("gw listening on 8080 and connected to %s", *endpoint)
	http.ListenAndServe(":8080", mux)
}

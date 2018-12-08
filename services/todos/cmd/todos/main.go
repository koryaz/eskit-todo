package main

import (
	"log"
	"net"
	"net/http"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/koryaz/eskit-todo/generated/grpc/go/todos"
	"github.com/koryaz/eskit-todo/services/todos/provider"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
)

type TodoStoreConfig struct {
	ListenAddr        string `json:"listenAddr" mapstructure:"listenAddr"`
	CrudStoreEndpoint string `json:"crudStoreEndpoint" mapstructure:"crudStoreEndpoint"`
}

func (c TodoStoreConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.CrudStoreEndpoint, validation.Required),
	)
}

func main() {

	viper.SetDefault("listenAddr", ":9090")
	viper.BindEnv("crudStoreEndpoint", "CRUDSTORE_ENDPOINT")

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/todostore")
	viper.AddConfigPath(".")

	var config TodoStoreConfig

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	if err := config.Validate(); err != nil {
		log.Fatalf("config validation : %v", err)
	}

	log.Println("Going to listen on : ", config.ListenAddr)
	lis, err := net.Listen("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	todoProvider, err := provider.NewTodoServiceProvider(config.CrudStoreEndpoint)
	if err != nil {
		log.Fatalf("todo provider failed initializing : %v", err)
	}

	grpc_prometheus.Register(s)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(":8888", nil)
	}()

	reflection.Register(s)
	todos.RegisterTodoServiceServer(s, todoProvider)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

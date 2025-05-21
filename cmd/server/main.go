package main

import (
	"log"
	"net"
	"net/http"

	"calculator-service/api"
	"calculator-service/internal/app"
	grpcHandler "calculator-service/internal/transport/grpc"
	httpHandler "calculator-service/internal/transport/http_handler"
	"google.golang.org/grpc"

	_ "calculator-service/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc/reflection"
)

// @title Calculator API
// @version 1.0
// @description Service for processing arithmetic instructions
// @host localhost:8080
// @BasePath /
func main() {
	calc := app.NewCalculator()

	handler := httpHandler.NewHandler(calc)
	http.HandleFunc("/calculate", handler.Calculate)
	
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.json"),
	))

	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	go func() {
		log.Println("HTTP server starting on :8080...")
		log.Println("Swagger UI: http://localhost:8080/swagger/index.html")
		log.Println("Raw JSON: http://localhost:8080/docs/swagger.json")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterCalculatorServer(grpcServer, grpcHandler.NewServer(calc))

	reflection.Register(grpcServer)
	log.Println("gRPC server starting on :50051...")
	log.Fatal(grpcServer.Serve(lis))
}
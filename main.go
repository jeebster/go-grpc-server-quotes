package main

import (
	"database/sql"
	"net"
	"os"
	"strings"

	"github.com/hashicorp/go-hclog"
	protos "github.com/jeebster/go-grpc-server-quotes/protos"
	"github.com/jeebster/go-grpc-server-quotes/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := hclog.Default()

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", "error", err)
		os.Exit(1)
	}

	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	defer db.Close()

	if err != nil {
		logger.Error("Error connecting to the dataabase", "error", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	quoteServer := server.NewQuoteServer(logger, db)

	protos.RegisterQuoteServiceServer(grpcServer, quoteServer)

	// this would be omitted from a production environment to prevent exposure of RPC definitions
	reflection.Register(grpcServer)

	// super over-engineered efficient string concatenation
	envPortString := os.Getenv("PORT")
	var portStringBuilder strings.Builder
	portStringBuilder.Grow(5)
	portStringBuilder.WriteByte(':')
	portStringBuilder.WriteString(envPortString)

	listen, err := net.Listen("tcp", portStringBuilder.String())
	if err != nil {
		logger.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	grpcServer.Serve(listen)
}

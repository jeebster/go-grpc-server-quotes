package server

import (
	"context"
	"database/sql"

	"github.com/hashicorp/go-hclog"
	protos "github.com/jeebster/go-grpc-server-quotes/protos"
	"github.com/jeebster/go-grpc-server-quotes/queries"
)

type QuoteServer struct {
	Logger   hclog.Logger
	Database *sql.DB
}

func NewQuoteServer(logger hclog.Logger, database *sql.DB) *QuoteServer {
	return &QuoteServer{
		Logger:   logger,
		Database: database,
	}
}

func (qs *QuoteServer) GetQuotes(ctx context.Context, req *protos.QuotesRequest) (*protos.QuotesResponse, error) {
	qs.Logger.Info("Handle GetQuotes", "limit", req.GetLimit())

	quotes, err := queries.GetQuotes(qs.Database, req.GetLimit())

	if err != nil {
		// TODO: handle error according to gRPC convention
	}

	return &protos.QuotesResponse{Quotes: quotes}, nil
}

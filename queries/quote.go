package queries

import (
	"database/sql"

	protos "github.com/jeebster/go-grpc-server-quotes/protos"
)

// GetQuotes queries database rows and returns structured data relative to protocol buffer definitions
func GetQuotes(database *sql.DB, limit int32) ([]*protos.Quote, error) {
	var quotes []*protos.Quote
	rows, err := database.Query("SELECT * FROM quotes LIMIT $1 ORDER BY updated_at", limit)

	if err != nil {
		return quotes, err
	}

	defer rows.Close()
	for rows.Next() {
		var quote protos.Quote
		err = rows.Scan(&quote.Id, &quote.Body, &quote.CreatedAt, &quote.UpdatedAt)
		if err != nil {
			return quotes, err
		}
		quotes = append(quotes, &quote)
	}

	if err := rows.Err(); err != nil {
		return quotes, err
	}

	return quotes, nil
}

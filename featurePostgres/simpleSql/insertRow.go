package simplesql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	userInput taskDTO,
	createdAt time.Time,
) error {
	sqlQuery := `
		INSERT INTO tasks (title, description, completed, created_at)
		VALUES ($1, $2, $3, $4);
	`
	_, err := conn.Exec(
		ctx, 
		sqlQuery, 
		userInput.Title, 
		userInput.Description, 
		userInput.Completed, 
		createdAt,
	)

	return err
}

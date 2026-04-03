package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		INSERT INTO tasks (title, description, completed, created_at)
		VALUES ('asdf', 'sdfasdf', 'FALSE', '2025-11-11 18:00:05');
	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err
}

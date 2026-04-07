package simpleconnection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName"
func CreateConnect(ctx context.Context) (*pgx.Conn, error) {
	connStr := os.Getenv("CONN_STR")
	//conn указатель на подключение к БД
	return pgx.Connect(ctx, connStr)
}

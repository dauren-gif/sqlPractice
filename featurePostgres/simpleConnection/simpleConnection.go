package simpleconnection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName"
func CreateConnect(ctx context.Context) (*pgx.Conn, error) {
	//conn указатель на подключение к БД
	return pgx.Connect(ctx, "postgres://postgres:1598752Losaasim@localhost:5432/postgres")
}

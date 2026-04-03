package simpleconnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourUserName:YourPassword@YourHostname:5431/YourDatabaseName"
func CheckConnect() {
	ctx := context.Background()
	//conn указатель на подключение к БД
	conn, err := pgx.Connect(ctx, "postgres://postgres:1598752Losaasim@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	//проверка подключения
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("connected")
}

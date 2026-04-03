package main

import (
	"context"
	"fmt"
	simpleconnection "sqlLesson/featurePostgres/simpleConnection"
	simplesql "sqlLesson/featurePostgres/simpleSql"
) 

func main() {
	ctx := context.Background()

	conn, err := simpleconnection.CreateConnect(ctx)
	if err != nil {
		panic(err)
	}

	if err := simplesql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	if err := simplesql.InsertRow(ctx, conn); err != nil {
		panic(err)
	}
	
	fmt.Println("succeed")
}

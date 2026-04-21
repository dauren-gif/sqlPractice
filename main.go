package main

import (
	"context"
	simpleconnection "sqlLesson/featurePostgres/simpleConnection"
	simplesql "sqlLesson/featurePostgres/simpleSql"
)

func main() {

	// if _, err := os.Create("out/newfile.txt"); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("сервер запущен")

	// if err := httpserver.StartHttpServer(); err != nil {
	// 	fmt.Println("ошибка запуска сервера")
	// }

	ctx := context.Background()

	conn, err := simpleconnection.CreateConnect(ctx)
	if err != nil {
		panic(err)
	}

	if err := simplesql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	tasks, err := simplesql.SelectTable(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.Id == "6" {
			task.Title = simplesql.taskDTO //надо создать http запрос
			//и получить тело JSON
			task.
			if err := simplesql.UpdateTask(ctx, conn, task); err != nil{
				panic(err)
			}
			break
		}
	}

}

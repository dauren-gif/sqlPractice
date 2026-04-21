package main

import (
	"fmt"
	httpserver "sqlLesson/http_server"
)

func main() {

	// if _, err := os.Create("out/newfile.txt"); err != nil {
	// 	panic(err)
	// }

	fmt.Println("сервер запущен")

	if err := httpserver.StartHttpServer(); err != nil {
		fmt.Println("ошибка запуска сервера")
	}

	// ctx := context.Background()

	// conn, err := simpleconnection.CreateConnect(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// if err := simplesql.CreateTable(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := simplesql.InsertRow(
	// 	ctx,
	// 	conn,
	// 	"wwwww",
	// 	"wwwwwwwwww",
	// 	false,
	// 	time.Now(),
	// ); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("succeed")

	// tasks, err := simplesql.SelectTable(ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, task := range tasks {
	// 	if task.Id == "6" {
	// 		task.Title = simplesql.taskDTO //надо создать http запрос
	// 		//и не получить тело JSON
	// 		task.
	// 		if err := simplesql.UpdateTask(ctx, conn, task); err != nil{
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// }

}

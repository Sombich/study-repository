package main

import (
	"context"
	"fmt"
	simpleconnection "git-training/simple_connection"
	simplesql "git-training/simple_sql"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	conn, err := simpleconnection.CreateConnection(ctx)

	if err != nil {
		panic(err)
	}

	if err := simplesql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	fmt.Println("Таблица в БД была успешно создана!")

}

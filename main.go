package main

import (
	"fmt"
	"git-training/feature1"
	"git-training/feature2"
	simpleconnection "git-training/simple_connection"
)

func main() {
	fmt.Println("main func")

	feature1.Feature()
	feature2.Feature2()

	simpleconnection.CheckConnection()
}

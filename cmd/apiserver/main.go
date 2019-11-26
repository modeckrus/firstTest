package main

import (
	"fmt"
	"myrestapi/internal/app/apiserver"
)

func main() {
	fmt.Println("Hello world")
	apiserver.NewDB()
}

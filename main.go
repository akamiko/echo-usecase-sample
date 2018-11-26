package main

import (
	"fmt"

	"github.com/akamiko/echo-usecase-sample/route"
)

func main() {
	fmt.Println("aaa")
	e := route.Init()
	e.Logger.Fatal(e.Start(":1323"))
}

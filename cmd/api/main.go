package main

import (
	"fmt"
	"github.com/Sharykhin/sp-event/storage/mongo"
)

func main() {
	fmt.Println("web api")

	session := mongo.NewSession()
	fmt.Println(session.Ping())
}

package main

import (
	"fmt"
	"github.com/Sharykhin/sp-event/storage/mongo"
)

func main() {
	fmt.Println("web api")

	_ := mongo.NewSession()
	fmt.Printf("haha %s")
}

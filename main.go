package main

import (
	"log"

	"github.com/golangExample/handler"

	"github.com/gin-gonic/gin"
	"github.com/golangExample/db"
)

func main() {
	e := gin.Default()
	//e.GET("/ping", handler.Ping)
	e.POST("/login", handler.Login)
	err := db.AddData()
	if err != nil {
		log.Fatalln(err)
	}
	// var check bool
	// check = handler.Login()
	// if check == true {
	// 	log.Printf("perfect")
	// }
	//var a map[string]db.User
	// err := db.GetData()
	// log.Printf("User: %v", err)
	// err := db.ChecConn()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	e.Run(":9090")
}

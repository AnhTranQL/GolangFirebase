package main

import (
	"github.com/golangExample/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	//e.GET("/ping", handler.Ping)
	e.POST("/login", handler.Login)
	e.POST("/logup", handler.Logup)

	//err := db.AddData()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
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

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golangExample/db"
)

func main() {
	e := gin.Default()
	//e.GET("/ping", handler.Ping)
	err := db.ConnectFirebase()
	if err != nil {
		log.Fatalln(err)
	}

	// err := db.ChecConn()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	e.Run(":9090")
}

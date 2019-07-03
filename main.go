package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golangExample/db"
	"github.com/golangExample/handler"
)

func main() {
	rs, err := db.CheckID("1562142687")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)
	e := gin.Default()
	e.POST("/login", handler.Login)
	e.POST("/logup", handler.Logup)
	e.POST("/getuser", handler.GetUserByEmail)
	e.PUT("/updateuser", handler.UpdateUserPhoneNumber)
	e.DELETE("/deleteuser", handler.DeleteUser)

	// if !db.CheckPhoneNumber() {
	// 	fmt.Printf("Full name Wrong")
	// } else {
	// 	fmt.Printf("Full name true")

	// }

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

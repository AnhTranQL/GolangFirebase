package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golangExample/db"
	"github.com/golangExample/handler"
)

func main() {
	//  err := db.CheckEmail()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(rs)

	if !db.CheckEmail() {
		fmt.Printf("Email Wrong")
	} else {
		fmt.Printf("Email true")

	}

	e := gin.Default()
	e.POST("/login", handler.Login)
	e.POST("/signup", handler.Signup)
	e.POST("/getuser", handler.GetUserByEmail)
	e.PUT("/updateuser", handler.UpdateUserPhoneNumber)
	e.PUT("/disableduser", handler.Disabled)
	e.PUT("/undisableduser", handler.Undisabled)
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

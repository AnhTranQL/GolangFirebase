package handler

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golangExample/db"
)

type AccLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	person := new(AccLogin)
	err := c.Bind(person)
	//Không đọc được dữ liệu
	if err != nil {
		c.JSON(404, map[string]string{
			"fault":   "Not found",
			"message": "Không thể đọc dữ liệu",
		})
		return
	}
	results, err1 := db.GlobalUsersRef.OrderByKey().GetOrdered(context.Background())
	if err1 != nil {
		log.Fatalln("Error querying database:", err)
	}
	var d db.User

	for _, r := range results {
		if err := r.Unmarshal(&d); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		if d.Email == person.Email {
			if d.Password == person.Password {
				c.JSON(200, map[string]string{
					"message": "Log in thanh cong",
				})
				return
			}
			c.JSON(400, map[string]string{
				"message": "Password khong dung",
			})
			return
		}
	}
	c.JSON(400, map[string]string{
		"message": "Email khong ton tai trong he thong",
	})
	return

}

// func Ping(c *gin.Context) {
// 	c.JSON(http.StatusOK, map[string]string{
// 		"message": "pong",
// 	})
// }

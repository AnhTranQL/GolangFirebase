package handler

type Person struct {
	Email       string `json:"email"`
	First       string `json:"first"`
	Last        string `json:"last"`
	PhoneNumber string `json:"phonenumber"`
	Password    string `json:"password"`
}

// func init() {
// 	List[0] = Person{Email: "1610131@hcmut.edu.vn", First: "Anh", Last: "Tran", PhoneNumber: "0123456789", Password: "123456"}
// 	List[1] = Person{Email: "thotranthi@gmail.com", First: "Tho", Last: "Tran", PhoneNumber: "0123456789", Password: "123456"}
// 	List[2] = Person{Email: "@hcmut.edu.vn", First: "Kiem", Last: "Tran", PhoneNumber: "0123456789", Password: "123456"}
// 	List[3] = Person{Email: "1610131@hcmut.edu.vn", First: "Binh", Last: "Tran", PhoneNumber: "0123456789", Password: "123456"}
// }

// func Ping(c *gin.Context) {
// 	c.JSON(http.StatusOK, map[string]string{
// 		"message": "pong",
// 	})
// }

package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golangExample/db"
)

type AccLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//func Logup là func dùng để đăng ký tài khoản
func Logup(c *gin.Context) {
	person := new(db.User)
	err := c.Bind(person)
	//Không đọc được dữ liệu
	if err != nil {
		c.JSON(404, map[string]string{
			"fault":   "Not found",
			"message": "Không thể đọc dữ liệu",
		})
		return
	}

	//Check các trường dữ liệu
	if person.FullName == "" {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Trường FullName là bắt buộc!",
		})
		return
	}
	if person.Email == "" {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Trường email là bắt buộc!",
		})
		return
	}
	//Check email có đúng định dạng
	// re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	// if re.MatchString(person.Email) {
	// 	c.JSON(400, map[string]string{
	// 		"fault":   "Bad request",
	// 		"message": "Email không đúng!",
	// 	})
	// 	return
	// }

	if person.PhoneNumber == "" {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Trường Phone Number là bắt buộc!",
		})
		return
	}

	if len(person.PhoneNumber) != 10 {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Phone number phải gồm 10 số!",
		})
		return
	}

	//Check phone number chứa chữ cái
	for _, v := range person.PhoneNumber {
		if v == '.' || v < '0' || v > '9' {
			c.JSON(400, map[string]string{
				"fault":   "Bad request",
				"message": "Phone number chỉ gồm số!",
			})
			return
		}
	}

	if person.Password == "" {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Trường password  là bắt buộc!",
		})
		return
	}

	if len(person.Password) < 6 {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Password của bạn quá ngắn!",
		})
		return
	}

	//Check email đăng ký đã tồn tại trong tài khoản hay chưa
	results, err1 := db.GlobalUsersRef.OrderByChild("email").GetOrdered(context.Background())
	if err1 != nil {
		log.Fatalln("Error querying database:", err1)
	}

	for _, r := range results {
		var d db.User
		if err := r.Unmarshal(&d); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		if d.Email == person.Email {

			c.JSON(400, map[string]string{
				"message": "Email đã tồn tại trong hệ thống",
			})
			return
		}
	}
	code := fmt.Sprintf("%d", time.Now().Unix())
	//Thêm account vào hệ thống
	_, err = db.GlobalUsersRef.Push(context.Background(), &db.User{

		FullName:    person.FullName,
		Email:       person.Email,
		PhoneNumber: person.PhoneNumber,
		Password:    person.Password,
		ID:          code,
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	person.ID = code
	c.JSON(http.StatusOK, person)
	return
}

//func là dunc dùng để log in vào hệ thống
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

	if person.Email == "" || person.Password == "" {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Trường email vaf password  là bắt buộc!",
		})
		return
	}

	//Retrieve data from firebase
	results, err1 := db.GlobalUsersRef.OrderByKey().GetOrdered(context.Background())
	if err1 != nil {
		log.Fatalln("Error querying database:", err1)
	}
	var d db.User

	for _, r := range results {
		if err := r.Unmarshal(&d); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		if d.Email == person.Email {
			if d.Password == person.Password {
				// c.JSON(200, map[string]string{
				// 	"message": "Log in thanh cong",
				// })
				c.JSON(http.StatusOK, d)
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

type UserEmail struct {
	Email string `json:"email"`
}

func GetUserByEmail(c *gin.Context) {
	person := new(UserEmail)
	err := c.Bind(person)
	//Không đọc được dữ liệu
	if err != nil {
		c.JSON(404, map[string]string{
			"fault":   "Not found",
			"message": "Không thể đọc dữ liệu",
		})
		return
	}

	if person.Email == "" {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Trường email vaf password  là bắt buộc!",
		})
		return
	}

	//Check email cần tìm kiếm có tồn tại trong hệ thống không, có thì lấy tài khoản
	results, err1 := db.GlobalUsersRef.OrderByKey().GetOrdered(context.Background())
	if err1 != nil {
		log.Fatalln("Error querying database:", err1)
	}

	for _, r := range results {
		var d db.User
		if err := r.Unmarshal(&d); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		if d.Email == person.Email {
			c.JSON(http.StatusOK, d)
			return
		}
	}
	//Email không tồn tại trong hệ thống
	c.JSON(400, map[string]string{
		"fault":   "Bad request",
		"message": "Email cần tìm không tồn tại trong hệ thống!",
	})
}

type userUpdate struct {
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phonenumber,omitempty"`
}

//func là func dùng để update lại số điện thoại của một tài khoản theo email
func UpdateUserPhoneNumber(c *gin.Context) {
	person := new(userUpdate)
	err := c.Bind(person)
	//Không đọc được dữ liệu
	if err != nil {
		c.JSON(404, map[string]string{
			"fault":   "Not found",
			"message": "Không thể đọc dữ liệu",
		})
		return
	}

	if person.Email == "" {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Trường email là bắt buộc!",
		})
		return
	}

	//Check email đúng hay không chưa làm dc nhé
	// valid := govalidator.IsEmail(person.Email)
	// if !valid {
	// 	c.JSON(400, map[string]string{
	// 		"fault":   "Bad request",
	// 		"message": "Email không đúng!",
	// 	})
	// 	return
	// }

	if person.PhoneNumber == "" {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Trường Phone Number là bắt buộc!",
		})
		return
	}
	//Check phone number k đủ 10 số
	if len(person.PhoneNumber) != 10 {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Phone number phải gồm 10 số!",
		})
		return
	}
	//Check phone number chứa chữ cái
	for _, v := range person.PhoneNumber {
		if v == '.' || v < '0' || v > '9' {
			c.JSON(400, map[string]string{
				"fault":   "Bad request",
				"message": "Phone number chỉ gồm số!",
			})
			return
		}
	}

	//Check email của tài khoản cần update có tồn tại trong hệ thống không,
	results, err1 := db.GlobalUsersRef.OrderByKey().GetOrdered(context.Background())
	if err1 != nil {
		log.Fatalln("Error querying database:", err1)
	}

	for _, r := range results {
		var d db.User
		if err := r.Unmarshal(&d); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		//Email tồn tại => Update phonenumber
		if d.Email == person.Email {
			hopperRef := db.GlobalUsersRef.Child(r.Key())
			if err := hopperRef.Update(context.Background(), map[string]interface{}{
				"phonenumber": person.PhoneNumber,
			}); err != nil {
				log.Fatalln("Error updating child:", err)
			}
			//Lấy lại thông tin sau khi update
			resultUpdate, err1 := db.GlobalUsersRef.OrderByKey().GetOrdered(context.Background())
			if err1 != nil {
				log.Fatalln("Error querying database:", err1)
			}

			for _, r := range resultUpdate {
				var dUpdate db.User
				if err := r.Unmarshal(&dUpdate); err != nil {
					log.Fatalln("Error unmarshaling result:", err)
				}
				if dUpdate.Email == person.Email {
					c.JSON(http.StatusOK, dUpdate)
					return
				}
			}
		}

	}
	//Email không tồn tại trong hệ thống
	c.JSON(400, map[string]string{
		"fault":   "Bad request",
		"message": "Email cần update không tồn tại trong hệ thống!",
	})
}

func DeleteUser(c *gin.Context) {
	person := new(UserEmail)
	err := c.Bind(person)
	//Không đọc được dữ liệu
	if err != nil {
		c.JSON(404, map[string]string{
			"fault":   "Not found",
			"message": "Không thể đọc dữ liệu",
		})
		return
	}

	if person.Email == "" {
		c.JSON(400, map[string]string{
			"fault":   "Bad request",
			"message": "Trường email vaf password  là bắt buộc!",
		})
		return
	}

	//Check email cần tìm kiếm có tồn tại trong hệ thống không, có thì xóa tài khoản
	results, err1 := db.GlobalUsersRef.OrderByKey().GetOrdered(context.Background())
	if err1 != nil {
		log.Fatalln("Error querying database:", err1)
	}

	for _, r := range results {
		var d db.User
		if err := r.Unmarshal(&d); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		//Email tooonftaij trong hệ thống
		if d.Email == person.Email {
			hopperRef := db.GlobalUsersRef.Child(r.Key())
			if err := hopperRef.Delete(context.Background()); err != nil {
				log.Fatalln("Error delete child:", err)
			}
			c.JSON(200, map[string]string{
				"message": "Xóa tài khoản thành công!",
			})
			return
		}

	}
	//Email không tồn tại trong hệ thống
	c.JSON(400, map[string]string{
		"fault":   "Bad request",
		"message": "Email cần xóa không tồn tại trong hệ thống!",
	})
}

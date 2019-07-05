package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golangExample/db"

	"github.com/golangExample/handler"
	"github.com/stretchr/testify/assert"
)

type User struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phonenumber,omitempty"`
	Password    string `json:"password,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	Disabled    bool   `json:"disabled"` //false là không khóa, true bị khóa

}

//func TestValidSignup dùng để test cho chức năng log up
func TestValidSignup(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/signup", handler.Signup)
	data := User{FullName: "Alan Turing22", Email: "nana@gmail.com", PhoneNumber: "0123456789", Password: "123456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/signup", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))
	assert.NotEmpty(t, respData.FullName, fmt.Sprintf("Bad request. Fullname null. Handler returned:  %v ", w.Code))
	assert.NotEmpty(t, respData.Email, fmt.Sprintf("Bad request. Email null. Handler returned:  %v", w.Code))
	assert.NotEmpty(t, respData.PhoneNumber, fmt.Sprintf("Bad request. PhoneNumber null. Handler returned:  %v ", w.Code))
	assert.NotEmpty(t, respData.Password, fmt.Sprintf("Bad request. Password null. Handler returned:  %v ", w.Code))

	//Check from firebase
	rs, err := db.GetOneItem(respData.ID)
	assert.NoError(t, err, fmt.Sprintf("Lỗi get one item. Handler returned:  %v ", w.Code))
	assert.Equal(t, rs.FullName, data.FullName, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))
	assert.Equal(t, rs.Email, data.Email, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))
	assert.Equal(t, rs.Password, data.Password, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))
	assert.Equal(t, rs.PhoneNumber, data.PhoneNumber, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))

}

//func TestInvalidSignupFullName dùng để teesst full name log up sai
func TestInvalidSignupFullName(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/signup", handler.Signup)
	data := User{FullName: "", Email: "thotranthinana@gmail.com", PhoneNumber: "0123456789", Password: "123456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/signup", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:   %v", w.Code))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: got %v ", w.Code))
	assert.Empty(t, respData.FullName, fmt.Sprintf("Bad request. Fullname null. Handler returned:  %v ", w.Code))
}

//func TestInvalidSignupEmail dùng để teesst email log up sai, email đã tồn tại
func TestInvalidSignEmail(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/signup", handler.Signup)
	data := User{FullName: "Best seller", Email: "nana22@gmail.com", PhoneNumber: "0123456789", Password: "123456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/signup", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: got %v ", w.Code))
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email null. Handler returned:  %v ", w.Code))
}

//func TestInvalidSignupPhoneNumber dùng để teesst phone number log up sai, phone number k đủ 10 số
func TestInvalidSignupPhoneNumber(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/signup", handler.Signup)
	data := User{FullName: "Best seller", Email: "", PhoneNumber: "123456789", Password: "123456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/signup", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned: %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned: %v  ", w.Code))
	assert.Empty(t, respData.PhoneNumber, fmt.Sprintf("Bad request. Phone number null.Handler returned: %v ", w.Code))
}

//func TestInvalidSignupPassword dùng để test password log up , password quá ngắn dưới 6 ký tự
func TestInvalidSignupPassword(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/signup", handler.Signup)
	data := User{FullName: "Best seller", Email: "aaa@gmail.com", PhoneNumber: "123456789", Password: "3456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/signup", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: got %v want %v", w.Code, 400))
	assert.Empty(t, respData.Password, fmt.Sprintf("Bad request. Password null. Handler returned:  %v ", w.Code))
}

//func TestValidLogIn dùng để test cho chức năng log in
func TestValidLogIn(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/login", handler.Login)
	data := User{Email: "nana22@gmail.com", Password: "123456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/login", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code:%v ", w.Code))
	assert.NotEmpty(t, respData.Email, fmt.Sprintf("Bad request. Email null. Handler returned:  %v ", w.Code))
	assert.NotEmpty(t, respData.Password, fmt.Sprintf("Bad request. Password null. Handler returned:  %v ", w.Code))

	//Check from firebase
	rs, err := db.GetOneItem(respData.ID)
	//value := rs.(struct User)
	assert.NoError(t, err, fmt.Sprintf("Lỗi get one item. Handler returned:  %v ", w.Code))
	assert.Equal(t, rs.Email, data.Email, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))
	assert.Equal(t, rs.Password, data.Password, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))

}

//func TestInvalidLoginEmail dùng để teesst email log in sai, email không tồn tại
func TestInvalidLoginEmail(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/login", handler.Login)
	data := User{Email: "nanalovely@gmail.com", Password: "123456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/login", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: %v", w.Code))
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email null. Handler returned:  %v ", w.Code))
}

//func TestInvalidLoginPassword dùng để teesst password log in sai, password sai
func TestInvalidLoginPassword(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/login", handler.Login)
	data := User{Email: "35ngocanh@gmail.com", Password: "12qq456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/login", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: %v", w.Code))
	assert.Empty(t, respData.Password, fmt.Sprintf("Bad request. Password sai. Handler returned:  %v ", w.Code))
}

//func TestValidGetUserByEmail dùng để test cho chức năng log in
func TestValidGetUserByEmail(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/getuser", handler.GetUserByEmail)
	data := User{Email: "nana@gmail.com"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/getuser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.NotEmpty(t, respData.Email, fmt.Sprintf("Bad request. Email null. Handler returned: %v", w.Code))

	//Check from firebase
	rs, err := db.GetOneItem(respData.ID)
	//value := rs.(struct User)
	assert.NoError(t, err, fmt.Sprintf("Lỗi get one item. Handler returned:  %v ", w.Code))
	assert.Equal(t, rs.Email, data.Email, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))

}

//func TestInvalidGetUserByEmailWithEmail dùng để test get user by email sai, vì email không tồn tại trong hệ thống
func TestInvalidGetUserByEmailWithEmail(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/getuser", handler.GetUserByEmail)
	data := User{Email: "ngocanh@gmail.com"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/getuser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: %v", w.Code))
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email không tồn tại trong hệ thống. Handler returned:  %v ", w.Code))
}

//func TestValidUpdateUserPhoneNumber dùng để test cho chức năng update phone number
func TestValidUpdateUserPhoneNumber(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/updateuser", handler.UpdateUserPhoneNumber)
	data := User{Email: "nana@gmail.com", PhoneNumber: "2390077844"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("PUT", "/updateuser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.NotEmpty(t, respData.Email, fmt.Sprintf("Bad request. Email null. Handler returned: %v", w.Code))
	assert.NotEmpty(t, respData.PhoneNumber, fmt.Sprintf("Bad request. Phone number rỗng. Handler returned:  %v", w.Code))

	//Check from firebase
	rs, err := db.GetOneItem(respData.ID)
	assert.NoError(t, err, fmt.Sprintf("Lỗi get one item. Handler returned:  %v ", w.Code))
	assert.Equal(t, rs.Email, data.Email, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))
	assert.Equal(t, rs.PhoneNumber, data.PhoneNumber, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))

}

//func TestInvalidUpdateUserWithEmail dùng để test update user by email sai, vì email không tồn tại trong hệ thống
func TestInvalidUpdateUserWithEmail(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/updateuser", handler.UpdateUserPhoneNumber)
	data := User{Email: "ngocanh@gmail.com", PhoneNumber: "0981726345"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("PUT", "/updateuser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code:  %v", w.Code))
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email không tồn tại trong hệ thống. Handler returned:  %v ", w.Code))
}

//func TestInvalidUpdateUserWithPhoneNumber dùng để test update user by email sai, vì phone number không đúng
func TestInvalidUpdateUserWithPhoneNumber(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/updateuser", handler.UpdateUserPhoneNumber)
	data := User{Email: "nana22@gmail.com", PhoneNumber: "12adf45678"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("PUT", "/updateuser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: %v", w.Code))
	assert.Empty(t, respData.PhoneNumber, fmt.Sprintf("Bad request. Phone number sai. Handler returned:  %v ", w.Code))
}

//func TestValidDeleteUser dùng để test delete user thành công
func TestValidDeleteUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.DELETE("/deleteuser", handler.DeleteUser)
	data := User{Email: "nana@gmail.com"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("DELETE", "/deleteuser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))
	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))

	//Check from firebase
	_, err := db.GetOneItem(respData.ID)
	assert.Error(t, err, fmt.Sprintf("Lỗi get one item. Handler returned:  %v ", w.Code)) //check xóa thành công
	// assert.Empty(t, rs.Email, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))
}

// func TestValidDisabledUser dùng để test khóa tài khoản của một user
func TestValidDisabledUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/disableduser", handler.Disabled)
	data := User{Email: "nana22@gmail.com"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("PUT", "/disableduser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.Equal(t, 200, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: %v", w.Code))
	assert.NotEmpty(t, respData.Email, fmt.Sprintf("Bad request. Email không tồn tại trong hệ thống. Handler returned: %v", w.Code))

	//Check từ firebase
	rs, err := db.GetOneItem(respData.ID)
	assert.NoError(t, err, fmt.Sprintf("Lỗi get one item. Handler returned:  %v ", w.Code))
	assert.Equal(t, rs.Email, data.Email, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))
	assert.Equal(t, rs.Disabled, true, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))
}

// func TestInValidDisabledUser dùng để test khóa tài khoản không thành công , Email không tồn tại trong hệ thống
func TestInValidDisabledUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/disableduser", handler.Disabled)
	data := User{Email: "baba@gmail.com"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("PUT", "/disableduser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: %v", w.Code))
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email không tồn tại trong hệ thống. Handler returned: %v", w.Code))
}

// func TestInValidDisabledUser dùng để test mở khóa tài khoản thành công

func TestValidUnDisabledUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/undisableduser", handler.Undisabled)
	data := User{Email: "nana22@gmail.com"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("PUT", "/undisableduser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.Equal(t, 200, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: %v", w.Code))
	assert.NotEmpty(t, respData.Email, fmt.Sprintf("Bad request. Email không tồn tại trong hệ thống. Handler returned: %v", w.Code))

	//Check từ firebase
	rs, err := db.GetOneItem(respData.ID)
	assert.NoError(t, err, fmt.Sprintf("Lỗi get one item. Handler returned:  %v ", w.Code))
	assert.Equal(t, rs.Email, data.Email, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))
	assert.Equal(t, rs.Disabled, false, fmt.Sprintf("Not acceptable. Handler returned %v", w.Code))
}

// func TestInValidDisabledUser dùng để test khóa tài khoản không thành công , Email không tồn tại trong hệ thống
func TestInValidUnDisabledUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/undisableduser", handler.Undisabled)
	data := User{Email: "baba@gmail.com"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("PUT", "/undisableduser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: %v", w.Code))
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email không tồn tại trong hệ thống. Handler returned: %v", w.Code))
}

//func TestInvalidDeleteUserWithEmail dùng để test delete user sai, vì email không tồn tại
func TestInvalidDeleteUserWithEmail(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.DELETE("/deleteuser", handler.DeleteUser)
	data := User{Email: "hello@gmail.com"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("DELETE", "/deleteuser", (reqData))
	assert.NoError(t, err0, fmt.Sprintf("Không thể tạo new request. Handler returned:  %v mong muốn %v", w.Code, 200))

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	resp, err1 := ioutil.ReadAll(w.Body)
	assert.NoError(t, err1, fmt.Sprintf("Không thể đọc được body. Handler returned:  %v mong muốn %v", w.Code, 200))

	var respData User
	err2 := json.Unmarshal(resp, &respData)
	assert.NoError(t, err2, fmt.Sprintf("Lỗi unmarshal data. Handler returned:  %v mong muốn %v", w.Code, http.StatusOK))

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: %v", w.Code))
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email không tồn tại trong hệ thống. Handler returned: %v", w.Code))
}

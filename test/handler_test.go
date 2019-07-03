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
	"github.com/golangExample/handler"
	"github.com/stretchr/testify/assert"
)

type User struct {
	FullName    string `json:"full_name,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phonenumber,omitempty"`
	Password    string `json:"password,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
}

//func TestValidLogup dùng để test cho chức năng log up
func TestValidLogup(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/logup", handler.Logup)
	data := User{FullName: "Alan Turing", Email: "nana@gmail.com", PhoneNumber: "0123456789", Password: "123456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/logup", (reqData))
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
	assert.NotEmpty(t, respData.Email, fmt.Sprintf("Bad request. Email null. Handler returned:  %v mong muốn %v", w.Code, 400))
	assert.NotEmpty(t, respData.PhoneNumber, fmt.Sprintf("Bad request. PhoneNumber null. Handler returned:  %v mong muốn %v", w.Code, 400))
	assert.NotEmpty(t, respData.Password, fmt.Sprintf("Bad request. Password null. Handler returned:  %v mong muốn %v", w.Code, 400))

}

//func TestInvalidLogupFullName dùng để teesst full name log up sai
func TestInvalidLogupFullName(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/logup", handler.Logup)
	data := User{FullName: "", Email: "thotranthinana@gmail.com", PhoneNumber: "0123456789", Password: "123456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/logup", (reqData))
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
	assert.Empty(t, respData.FullName, fmt.Sprintf("Bad request. Fullname null. Handler returned:  %v ", w.Code))
}

//func TestInvalidLogupEmail dùng để teesst email log up sai, email đã tồn tại
func TestInvalidLogupEmail(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/logup", handler.Logup)
	data := User{FullName: "Best seller", Email: "nana@gmail.com", PhoneNumber: "0123456789", Password: "123456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/logup", (reqData))
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
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email null. Handler returned:  %v ", w.Code))
}

//func TestInvalidLogupPhoneNumber dùng để teesst phone number log up sai, phone number k đủ 10 số
func TestInvalidLogupPhoneNumber(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/logup", handler.Logup)
	data := User{FullName: "Best seller", Email: "", PhoneNumber: "123456789", Password: "123456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/logup", (reqData))
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
	assert.Empty(t, respData.PhoneNumber, fmt.Sprintf("Bad request. Phone number null. Handler returned:  %v ", w.Code))
}

//func TestInvalidLogupPassword dùng để test password log up , password quá ngắn dưới 6 ký tự
func TestInvalidLogupPassword(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/logup", handler.Logup)
	data := User{FullName: "Best seller", Email: "aaa@gmail.com", PhoneNumber: "123456789", Password: "3456"}

	reqData := new(bytes.Buffer)
	json.NewEncoder(reqData).Encode(data)

	req, err0 := http.NewRequest("POST", "/logup", (reqData))
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
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: got %v want %v", w.Code, 400))
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

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: got %v want %v", w.Code, 400))
	assert.Empty(t, respData.Password, fmt.Sprintf("Bad request. Password sai. Handler returned:  %v ", w.Code))
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

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: got %v want %v", w.Code, 400))
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email không tồn tại trong hệ thống. Handler returned:  %v ", w.Code))
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

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: got %v want %v", w.Code, 400))
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email không tồn tại trong hệ thống. Handler returned:  %v ", w.Code))
}

//func TestInvalidUpdateUserWithPhoneNumber dùng để test update user by email sai, vì phone number không đúng
func TestInvalidUpdateUserWithPhoneNumber(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/updateuser", handler.UpdateUserPhoneNumber)
	data := User{Email: "35ngocanh@gmail.com", PhoneNumber: "12adf45678"}

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

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: got %v want %v", w.Code, 400))
	assert.Empty(t, respData.PhoneNumber, fmt.Sprintf("Bad request. Phone number sai. Handler returned:  %v ", w.Code))
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

	assert.NotNil(t, resp, fmt.Sprintf("Not found. Body  null. Handler returned wrong status code: got %v ", w.Code))
	assert.Equal(t, 400, w.Code, fmt.Sprintf("Bad request. Handler returned wrong status code: got %v want %v", w.Code, 400))
	assert.Empty(t, respData.Email, fmt.Sprintf("Bad request. Email không tồn tại trong hệ thống. Handler returned:  %v ", w.Code))
}

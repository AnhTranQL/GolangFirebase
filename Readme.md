Cài đặt: go get github.com/gin-gonic/gin
go get -u firebase.google.com/go
go get -u google.golang.org/api/option
Sau đó Run : go run main.go

Test: http://127.0.0.1:9090/logup
Input:
{
	"email": "tran.kute@gmail.com",
	"full_name": "Hehe",
	"password": "ilikeit",
	"phonenumber": "1238906548"
}
Output:
{
    "id": "1562149317",
    "full_name": "Hehe",
    "email": "tran.kute@gmail.com",
    "phonenumber": "1238906548",
    "password": "ilikeit"
}

Test http://127.0.0.1:9090/login
Input:
{
	"email": "nana22@gmail.com",
	"password": "123456"
}
Output:
{
    "id": "1562147146",
    "full_name": "Alan Turing22",
    "email": "nana22@gmail.com",
    "phonenumber": "2390077865",
    "password": "123456"
}

Test http://127.0.0.1:9090/getuser
Input:
{
	"email": "nana22@gmail.com"
}
Output:
{
    "id": "1562147146",
    "full_name": "Alan Turing22",
    "email": "nana22@gmail.com",
    "phonenumber": "2390077865",
    "password": "123456"
}

Test http://127.0.0.1:9090/updateuser
Input:
{
	"email": "nana22@gmail.com",
	"phonenumber": "1231211167"
}
Output:
{
    "id": "1562147146",
    "full_name": "Alan Turing22",
    "email": "nana22@gmail.com",
    "phonenumber": "2390077865",
    "password": "123456"
}

Test http://127.0.0.1:9090/deleteuser
Input:
{
	"email": "hhhhhhh@gmail.com"
}
Output:
{
    "message": "Xóa tài khoản thành công!"
}

Test http://127.0.0.1:9090/disableduser
Input
{
	"email": "abc@gmail.com"
}
Output
{
    "id": "1562211090",
    "full_name": "Hehe",
    "email": "abc@gmail.com",
    "phonenumber": "1238901238",
    "password": "ilikeit",
    "disabled": true
}

Test http://127.0.0.1:9090/undisableduser
Input
{
	"email": "abc@gmail.com"
}
Output
{
    "id": "1562211090",
    "full_name": "Hehe",
    "email": "abc@gmail.com",
    "phonenumber": "1238901238",
    "password": "ilikeit",
    "disabled": false
}


Để test toàn bộ func: go test test/handler_test.go
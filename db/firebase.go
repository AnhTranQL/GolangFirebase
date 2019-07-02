package db

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	db "firebase.google.com/go/db"

	"google.golang.org/api/option"
)

type User struct {
	FullName    string `json:"full_name,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phonenumber,omitempty"`
	Password    string `json:"password,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
}

var GlobalClient *db.Client

var GlobalUsersRef *db.Ref

func init() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("db/demofirebasego.json")
	config := &firebase.Config{DatabaseURL: "https://demofirebase-3d6aa.firebaseio.com"}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		panic(fmt.Sprintf("error initializing app: %v", err))
	}

	global, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
	GlobalClient = global
	log.Println("Global client init: %v", GlobalClient)

}

func getGlobal() *db.Client {
	if GlobalClient == nil {
		log.Fatalln("Error get global:", nil)
	}
	return GlobalClient
}

// Seed or put data
func AddData() error {
	conn := getGlobal()

	ref := conn.NewRef("fireblog")
	GlobalUsersRef = ref.Child("users")
	err := GlobalUsersRef.Set(context.Background(), map[string]*User{
		"alanisawesome": {
			FullName:    "Alan Turing",
			Email:       "thotranthinana@gmail.com",
			PhoneNumber: "123456789",
			Password:    "123456",
		},
		"gracehop": {
			FullName:    "Grace Beauty",
			Email:       "35ngocanh@gmail.com",
			PhoneNumber: "0987654321",
			Password:    "123456",
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	return nil
}

// func GetData() error {
// 	err := GlobalUsersRef.Get(context.Background(), map[string]*User)
// 	if err != nil {
// 		log.Fatalln("Error getting value:", err)
// 	}
// 	return err
// }

// func addData(db, table string, data interface{}) {
// 	// new table
// 	// set data
// }
// func CheckUserName() {
// 	conn := getConn()
// 	conn.CheckIsExist
// }

// func connectFirebase() {
// 	// load file cfg
// 	// connect firebase
// 	// return conn
// }
// func ChecConn() error {
// 	// to do sth
// 	return nil
// }

// connect firebase
// ctx := context.Background()
// opt := option.WithCredentialsFile("db/demofirebasego.json")
// config := &firebase.Config{DatabaseURL: "https://demofirebase-3d6aa.firebaseio.com"}

// app, err := firebase.NewApp(ctx, config, opt)
// if err != nil {
// 	panic(fmt.Sprintf("error initializing app: %v", err))
// }

// client, err := app.Database(ctx)
// if err != nil {
// 	log.Fatalln("Error initializing database client:", err)
// }

// ref := client.NewRef("fireblog")

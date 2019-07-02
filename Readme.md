func ConnectFirebase() error {
	// connect firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile("db/demofirebasego.json")
	config := &firebase.Config{DatabaseURL: "https://demofirebase-3d6aa.firebaseio.com"}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		panic(fmt.Sprintf("error initializing app: %v", err))
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	ref := client.NewRef("fireblog")

	usersRef := ref.Child("users")
	err = usersRef.Set(ctx, map[string]*User{
		"alanisawesome": {
			FullName:    "Alan Turing",
			Email:       "thotranthi@gmail.com",
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

	log.Println("Done Connect firebase")
	return (nil)
}
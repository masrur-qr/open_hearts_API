package createadmin

import (
	"context"
	env "docs/app/Env"
	"docs/app/hashedpasswod"
	"docs/app/mongoconnect"
	"docs/app/structs"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Createadmin() {
	client, ctx := mongoconnect.DBConnection()
	Connections := client.Database(env.Data_Name).Collection("Users")
	// if admin exist not insert
	found := Connections.FindOne(ctx, bson.M{
		"permission": "Admin",
	})
	var shablon structs.UserStruct
	found.Decode(&shablon)

	if shablon.Permission == "Admin"  {
		fmt.Println("admin is exsist")
	} else {

		ID := primitive.NewObjectID().Hex()
		Password :="Admin"
		Hashed,_ := hashedpasswod.HashPassword(Password)
		Connections.InsertOne(context.Background(), structs.UserStruct{
			Id: ID,
			Phone:      "123",
			Password:  Hashed ,
			Permission: "Admin",
			Email:      "broimshoevmurtazom4@gmail.com",
			Ru: structs.RussianUser{
				Name:       "Admin",
				Surname:    "Admin",
			},
			En: structs.EnglishUser{
				Name: "Admin",
				Surname:    "Admin",
			},
		})
	}
}

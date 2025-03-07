package createadmin

import (
	"context"
	"math/rand"

	env "docs/app/Env"
	"docs/app/baner"
	"docs/app/createimagephoto"
	"docs/app/emptyfieldcheker"
	"docs/app/hashedpasswod"
	"docs/app/mongoconnect"
	returnjwt "docs/app/returnJwt"
	"docs/app/structs"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Createadmin() {
	client, ctx := mongoconnect.DBConnection()
	Connections := client.Database(env.Data_Name).Collection("Users")
	// if admin exist not insert
	found := Connections.FindOne(ctx, bson.M{
		"permission": "MainAdmin",
	})
	var shablon structs.UserStruct
	found.Decode(&shablon)

	if shablon.Permission == "MainAdmin" {
		fmt.Println("MainAdmin is exsist")
	} else {
		ID := primitive.NewObjectID().Hex()
		Password := "Admin"
		Hashed, _ := hashedpasswod.HashPassword(Password)
		Connections.InsertOne(context.Background(), structs.UserStruct{
			Id:         ID,
			Photo:      "",
			Phone:      "123",
			Password:   Hashed,
			Permission: "MainAdmin",
			Email:      "murtazobroimshoevm4@gmail.com",
			Ru: structs.LangForUser{
				Name: "Admin",
			},
			En: structs.LangForUser{
				Name: "Admin",
			},
		})
	}
}
func AdminRegistration(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(401, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(403, "error")
		} else {
			var AdminData structs.UserStruct
			c.ShouldBindJSON(&AdminData)
			Emptyfield, err := emptyfieldcheker.EmptyField(AdminData, "Id", "Permission")
			if Emptyfield {
				c.JSON(400, err)
			} else {
				folderName := "AdminPhoto"
				FolderError := createimagephoto.CreateFolder(folderName)
				if FolderError != nil {
					fmt.Printf("FolderError: %v\n", FolderError)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)

				AdminData.Photo = baner.ImageFunc(AdminData.Photo, ForImage, folderName)

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("Users")
				ID := primitive.NewObjectID().Hex()
				Hashed, _ := hashedpasswod.HashPassword(AdminData.Password)
				_, inserterror := createDB.InsertOne(ctx, bson.M{

					"_id":        ID,
					"photo":      folderName + "/" + AdminData.Photo,
					"phone":      AdminData.Phone,
					"password":   Hashed,
					"permission": "Admin",
					"email":      AdminData.Email,
					"ru": structs.LangForUser{
						Name: AdminData.Ru.Name,
					},
					"en": structs.LangForUser{
						Name: AdminData.En.Name,
					},
				})
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(200, "succes")
				}
			}
		}
	}
}

func UpdateAdmin(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(401, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "MainAdmin" && SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(403, "error")
		} else {
			var Update_Admin structs.UserStruct
			c.ShouldBindJSON(&Update_Admin)
			Emptyfield, err := emptyfieldcheker.EmptyField(Update_Admin, "Permission")
			if Emptyfield {
				c.JSON(400, err)
			} else {
				folder_Name := "AdminPhoto"
				FolderError := createimagephoto.CreateFolder(folder_Name)
				if FolderError != nil {
					fmt.Printf("FolderError: %v\n", FolderError)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Update_Admin.Photo = baner.ImageFunc(Update_Admin.Photo, ForImage, folder_Name)

				client, ctx := mongoconnect.DBConnection()
				collection := client.Database(env.Data_Name).Collection("Users")

				Hashed, _ := hashedpasswod.HashPassword(Update_Admin.Password)
				// Метод findAndUpdate: найти документ и обновить его
				result := collection.FindOneAndUpdate(
					ctx,
					bson.D{
						{"_id", Update_Admin.Id}, // Находим документ по ID
					},
					bson.D{
						{"$set", bson.D{
							{"email", Update_Admin.Email},
							{"phone", Update_Admin.Phone},
							{"photo", folder_Name + "/" + Update_Admin.Photo},
							{"password", Hashed},
							{"ru", Update_Admin.Ru.Name},
							{"en", Update_Admin.En.Name},
						}},
					},
				)
				var Dbdata structs.UserStruct
				result.Decode(&Dbdata)
				if Dbdata.Photo != "" {
					deleteeror := os.RemoveAll("./Statics/" + Dbdata.Photo)
					if deleteeror != nil {
						fmt.Printf("err: %v\n", deleteeror)
					}
				}

				c.JSON(200, "succes")
			}
		}
	}
}

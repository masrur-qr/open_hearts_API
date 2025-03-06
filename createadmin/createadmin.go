package createadmin

import (
	"context"
	"math/rand"

	env "docs/app/Env"
	"docs/app/baner"
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
		if SecretKeyData.Permission != "MainAdmin"  && isvalid {
			c.JSON(403, "error")
		} else {
			var AdminData structs.UserStruct
			c.ShouldBindJSON(&AdminData)
			Emptyfield, err := emptyfieldcheker.EmptyField(AdminData, "Id", "Permission")
			if Emptyfield {
				c.JSON(400, err)
			} else {
				folderName := "AdminPhoto"
				err := os.Mkdir("Statics"+"/"+folderName, os.ModePerm)
				if err != nil {
					if os.IsExist(err) {
						fmt.Println(" Папка уже существует.")
					} else {
						fmt.Println("Ошибка при создании папки")
						return
					}
				} else {
					fmt.Println("Папка успешно создана.")
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
				client, ctx := mongoconnect.DBConnection()
				Connections := client.Database(env.Data_Name).Collection("Users")
				FindRezult := Connections.FindOne(ctx, bson.M{
					"_id": Update_Admin.Id,
				})
				var Dbdata structs.UserStruct
				FindRezult.Decode(&Dbdata)
				if Dbdata.Id != "" {
					if Dbdata.Photo != "" {
						deleteeror := os.RemoveAll("./Statics/" + Dbdata.Photo)
						if deleteeror != nil {
							fmt.Printf("err: %v\n", deleteeror)
						}
					}

					folder_Name := "AdminPhoto"
					err := os.Mkdir("Statics"+"/"+folder_Name, os.ModePerm)
					if err != nil {
						if os.IsExist(err) {
							fmt.Println(" Папка уже существует.")
						} else {
							fmt.Println("Ошибка при создании папки")
							return
						}
					} else {
						fmt.Println("Папка успешно создана.")
					}
					rndName := rand.Intn(10000)
					ForImage := fmt.Sprintf("image_%v.png", rndName)
					Update_Admin.Photo = baner.ImageFunc(Update_Admin.Photo, ForImage, folder_Name)

					Hashed, _ := hashedpasswod.HashPassword(Update_Admin.Password)
					_, err2 := Connections.UpdateOne(ctx,
						bson.M{
							"_id": Update_Admin.Id,
						},
						bson.D{
							{Key: "$set", Value: bson.M{
								"photo":    folder_Name + "/" + Update_Admin.Photo,
								"phone":    Update_Admin.Phone,
								"password": Hashed,
								"email":    Update_Admin.Email,
								"ru": structs.LangForUser{
									Name: Update_Admin.Ru.Name,
								},
								"en": structs.LangForUser{
									Name: Update_Admin.En.Name,
								},
							},
							},
						},
					)
					if err2 != nil {
						fmt.Printf("err: %v\n", err)
					}
					c.JSON(200, "succes")
				} else {
					c.JSON(404, "User 88not founded")
				}
			}
		}
	}

}

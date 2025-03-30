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
	Connections := client.Database(env.Data_Name).Collection("users")
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
			Emptyfield, err := emptyfieldcheker.EmptyField(AdminData, "Id", "Permission", "Code")
			if Emptyfield {
				c.JSON(400, err)
			} else {
				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("users")
				findrezult := createDB.FindOne(ctx, bson.M{
					"phone": AdminData.Phone,
				})
				var Dbdata structs.UserStruct
				findrezult.Decode(&Dbdata)
				if Dbdata.Email == "" && Dbdata.Phone == "" {
					folderName := "AdminPhoto"
					FolderError := createimagephoto.CreateFolder(folderName)
					if FolderError != nil {
						fmt.Printf("FolderError: %v\n", FolderError)
					}
					rndName := rand.Intn(10000)
					ForImage := fmt.Sprintf("image_%v.png", rndName)

					Photo := baner.ImageFunc(AdminData.Photo, ForImage, folderName)
					AdminData.Photo = folderName + "/" + Photo

					AdminData.Id = primitive.NewObjectID().Hex()
					Hashed, _ := hashedpasswod.HashPassword(AdminData.Password)
					AdminData.Password = Hashed
					_, inserterror := createDB.InsertOne(ctx, structs.UserStruct{
						Id:         AdminData.Id,
						Photo:      AdminData.Photo,
						Phone:      AdminData.Phone,
						Password:   Hashed,
						Permission: "Admin",
						Email:      AdminData.Email,
						Ru: structs.LangForUser{
							Name: AdminData.Ru.Name,
						},
						En: structs.LangForUser{
							Name: AdminData.En.Name,
						},
					})
					if inserterror != nil {
						fmt.Printf("inserterror: %v\n", inserterror)
					} else {
						c.JSON(200, "succes")
					}
				} else {
					c.JSON(400, "error Admin alrady exsist")
				}
			}
		}
	}
}

func UpdateAdmin(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(401, "error: No Cookie found")
		return
	}
	SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
	if isvalid || (SecretKeyData.Permission != "MainAdmin" && SecretKeyData.Permission != "Admin") {
		c.JSON(403, "error: Unauthorized access")
		return
	}
	var Update_Admin structs.UserStruct
	if err := c.ShouldBindJSON(&Update_Admin); err != nil {
		c.JSON(400, "error: Invalid JSON data")
		return
	}
	Emptyfield, err := emptyfieldcheker.EmptyField(Update_Admin, "Permission", "Code")
	if Emptyfield {
		c.JSON(400, err)
		return
	}

	folderName := "AdminPhoto"
	FolderError := createimagephoto.CreateFolder(folderName)
	if FolderError != nil {
		fmt.Printf("Folder error: %v\n", FolderError)
	}
	rndName := rand.Intn(10000)
	ForImage := fmt.Sprintf("image_%v.png", rndName)
	Photo := baner.ImageFunc(Update_Admin.Photo, ForImage, folderName)
	Update_Admin.Photo = folderName + "/" + Photo

	client, ctx := mongoconnect.DBConnection()
	collection := client.Database(env.Data_Name).Collection("users")


	Hashed, _ := hashedpasswod.HashPassword(Update_Admin.Password)

	result := collection.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "_id", Value: Update_Admin.Id}},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "photo", Value: Update_Admin.Photo},
				{Key: "email", Value: Update_Admin.Email},
				{Key: "phone", Value: Update_Admin.Phone},
				{Key: "password", Value: Hashed},
				{Key: "ru", Value: structs.LangForProjectStatistic{Name: Update_Admin.Ru.Name}},
				{Key: "en", Value: structs.LangForProjectStatistic{Name: Update_Admin.En.Name}},
			}},
		},
	)


	var Dbdata structs.UserStruct
	if err := result.Decode(&Dbdata); err != nil {
		c.JSON(404, "Admin not found or failed to update")
		return
	}

	if Dbdata.Photo != "" {
		deleteError := os.RemoveAll("./Statics/" + Dbdata.Photo)
		if deleteError != nil {
			fmt.Printf("error removing old photo: %v\n", deleteError)
		}
	}
	c.JSON(200,"Succes")
}

package controlers

import (
	env "docs/app/Env"
	"docs/app/baner"
	"docs/app/createimagephoto"
	"docs/app/emptyfieldcheker"
	"docs/app/mongoconnect"
	returnjwt "docs/app/returnJwt"
	"docs/app/structs"
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddStatistic(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "error:only admin have ecses to add")
		} else {
			client, ctx := mongoconnect.DBConnection()
			createDB := client.Database(env.Data_Name).Collection("Statistic")
			count, err := createDB.CountDocuments(ctx, bson.M{})
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			if count >= 4 {
				c.JSON(400, "You can't add another statistic")
			} else {
				var statistic_shablon structs.AddStatistic
				c.ShouldBindJSON(&statistic_shablon)
				Emptyfield, err := emptyfieldcheker.EmptyField(statistic_shablon, "Id")
				if Emptyfield {
					c.JSON(404, err)
				} else {

					statistic_shablon.Id= primitive.NewObjectID().Hex()
					_, inserterror := createDB.InsertOne(ctx, structs.AddStatistic{
						Id:      statistic_shablon.Id,
						Quantity: statistic_shablon.Quantity,
						Ru: structs.LangForStatistic{
							Description: statistic_shablon.Ru.Description,
						},
						En:  structs.LangForStatistic{
							Description: statistic_shablon.En.Description,
						},
					})
					if inserterror != nil {
						fmt.Printf("inserterror: %v\n", inserterror)
					} else {
						c.JSON(201, "succes")
					}
				}
			}
		}
	}
}


func AddPatientStory(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "error:only admin have ecses to add")
		} else {

			var Patient_data structs.Patient_story
			c.ShouldBindJSON(&Patient_data)

			Emptyfield, err := emptyfieldcheker.EmptyField(Patient_data, "Id",)
			if Emptyfield {
				c.JSON(404, err)
			} else {
				folderName := "PatientProfilPhoto"
				folderName2 := "PatientBagroundPhoto"

				Foldererror:=createimagephoto.CreateFolder(folderName)
				if Foldererror!=nil{
					fmt.Printf("Foldererror: %v\n", Foldererror)
				}
				Foldererror2:=createimagephoto.CreateFolder(folderName2)
				if Foldererror2!=nil{
					fmt.Printf("Foldererror: %v\n", Foldererror2)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Patient_data.Photo = baner.ImageFunc(Patient_data.Photo, ForImage, folderName)
				Patient_data.Smallphoto = baner.ImageFunc(Patient_data.Smallphoto, ForImage, folderName2)

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("PatientStory")

				Patient_data.Id = primitive.NewObjectID().Hex()
				
				_, inserterror := createDB.InsertOne(ctx, structs.Patient_story{
					Id:   Patient_data.Id,
					Photo: folderName + "/" + Patient_data.Photo,
					Smallphoto:folderName2 +"/"+ Patient_data.Smallphoto,
					Ru: structs.LangForPatient{
						Full_Name:        Patient_data.Ru.Full_Name,
						Description:Patient_data.Ru.Description,
						Quot:      Patient_data.Ru.Quot,
					},
					En:structs.LangForPatient{
						Full_Name:        Patient_data.En.Full_Name,
						Description:Patient_data.En.Description,
						Quot:      Patient_data.En.Quot,
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
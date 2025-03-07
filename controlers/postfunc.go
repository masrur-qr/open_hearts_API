package controlers

import (
	"context"
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
			createDB := client.Database(env.Data_Name).Collection("statistic")
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

					statistic_shablon.Id = primitive.NewObjectID().Hex()
					_, inserterror := createDB.InsertOne(ctx, structs.AddStatistic{
						Id:       statistic_shablon.Id,
						Quantity: statistic_shablon.Quantity,
						Ru: structs.LangForStatistic{
							Description: statistic_shablon.Ru.Description,
						},
						En: structs.LangForStatistic{
							Description: statistic_shablon.En.Description,
						},
					})
					if inserterror != nil {
						fmt.Printf("inserterror: %v\n", inserterror)
					} else {
						c.JSON(201, "success")
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

			Emptyfield, err := emptyfieldcheker.EmptyField(Patient_data, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				folderName := "PatientProfilPhoto"
				folderName2 := "PatientBagroundPhoto"

				Foldererror := createimagephoto.CreateFolder(folderName)
				if Foldererror != nil {
					fmt.Printf("Foldererror: %v\n", Foldererror)
				}
				Foldererror2 := createimagephoto.CreateFolder(folderName2)
				if Foldererror2 != nil {
					fmt.Printf("Foldererror: %v\n", Foldererror2)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Photo := baner.ImageFunc(Patient_data.Photo, ForImage, folderName)
				Smallphoto := baner.ImageFunc(Patient_data.Smallphoto, ForImage, folderName2)

				Patient_data.Photo = folderName + "/" + Photo
				Patient_data.Smallphoto = folderName2 + "/" + Smallphoto

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("PatientStory")

				Patient_data.Id = primitive.NewObjectID().Hex()

				_, inserterror := createDB.InsertOne(ctx, Patient_data)
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(200, "succes")
				}
			}
		}

	}
}

func AddPartner(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "error")
		} else {

			var PartnerData structs.Partner
			c.ShouldBindJSON(&PartnerData)
			Emptyfield, err := emptyfieldcheker.EmptyField(PartnerData, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {

				folderName := "Partners"
				Foldererror := createimagephoto.CreateFolder(folderName)
				if Foldererror != nil {
					fmt.Printf("Foldererror: %v\n", Foldererror)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Logo := baner.ImageFunc(PartnerData.Logo, ForImage, folderName)
				PartnerData.Logo = folderName + "/" + Logo
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database(env.Data_Name).Collection("partners")
				PartnerData.Id = primitive.NewObjectID().Hex()
				_, inserterror := createDB.InsertOne(ctx, PartnerData)
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(200, "success")
				}
			}
		}
	}
}

func Add_statistic_for_project(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" &&isvalid {
			c.JSON(404, "error:only admin have ecses to add")
		} else {
			var statistic_shablon structs.AddStatisticForCenter
			c.ShouldBindJSON(&statistic_shablon)

			Emptyfield, err := emptyfieldcheker.EmptyField(statistic_shablon, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database(env.Data_Name).Collection("statistic_for_project")
				count, err := createDB.CountDocuments(ctx, bson.M{})
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}
				if count >= 3 {
					c.JSON(400, "You can't add another statistic")
				} else {
					statistic_shablon.Id = primitive.NewObjectID().Hex()
					_, inserterror := createDB.InsertOne(ctx, statistic_shablon)
					if inserterror != nil {
						fmt.Printf("inserterror: %v\n", inserterror)
					} else {
						c.JSON(201, "success")
					}
				}
			}

		}
	}
}


func AddTeamMambers(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error")
		} else {

			var TeamData structs.Team
			c.ShouldBindJSON(&TeamData)
			Emptyfield, err := emptyfieldcheker.EmptyField(TeamData, "Id")

			if Emptyfield {
				c.JSON(400,err)
			} else {
				folderName := "Team"

				FolderError:= createimagephoto.CreateFolder(folderName)
				if FolderError!= nil{
					fmt.Printf("FolderError: %v\n", FolderError)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Photo := baner.ImageFunc(TeamData.Photo, ForImage, folderName)
				client, ctx := mongoconnect.DBConnection()
				TeamData.Photo=folderName+"/"+Photo
				var createDB = client.Database(env.Data_Name).Collection("team")
				TeamData.Id= primitive.NewObjectID().Hex()
				insertrezult, inserterror := createDB.InsertOne(ctx,TeamData)
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(200, "success")
					fmt.Printf("insertrezult: %v\n", insertrezult)
				}
			}
		}

	}
}

func AddProgram(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" &&SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "error")
		} else {

			var Services structs.Program
			c.ShouldBindJSON(&Services)
			Emptyfield, err := emptyfieldcheker.EmptyField(Services, "Id","Servisec","LastDescription")

			if Emptyfield {
					c.JSON(404, err)
				} else {
				folderName := "program"

				FolderError:= createimagephoto.CreateFolder(folderName)
				if FolderError!= nil{
					fmt.Printf("FolderError: %v\n", FolderError)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Photo := baner.ImageFunc(Services.Photo, ForImage, folderName)
				Services.Photo=folderName+"/"+Photo
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database(env.Data_Name).Collection("programs")
				Services.Id = primitive.NewObjectID().Hex()
				insertrezult, inserterror := createDB.InsertOne(ctx, Services)
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(200, "success")
					fmt.Printf("insertrezult: %v\n", insertrezult)
				}
			}
		}

	}
}

func AddServices(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "error")
		} else {

			var Services structs.Services
			c.ShouldBindJSON(&Services)
			Emptyfield, err := emptyfieldcheker.EmptyField(Services, "Id")

			if Emptyfield {
					c.JSON(404, err)
				} else {
				folderName := "servisec"

				FolderError:= createimagephoto.CreateFolder(folderName)
				if FolderError!= nil{
					fmt.Printf("FolderError: %v\n", FolderError)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Photo := baner.ImageFunc(Services.Photo, ForImage, folderName)
				Services.Photo=folderName+"/"+Photo
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database(env.Data_Name).Collection("servisec")
				fmt.Printf("Patient_data: %v\n", Services)
				Services.Id= primitive.NewObjectID().Hex()
				_, inserterror := createDB.InsertOne(ctx,Services)
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(200, "success")
	
				}
			}
		}

	}
}

func AddStatisticForCenter()  {
	client, ctx := mongoconnect.DBConnection()
	Connections := client.Database(env.Data_Name).Collection("center_statistic")
	// if admin exist not insert
	found := Connections.FindOne(ctx, bson.M{
		"permission": "MainAdmin",
	})
	var shablon structs.UserStruct
	found.Decode(&shablon)

	if shablon.Permission == "MainAdmin" {
		fmt.Println("MainAdmin is exsist")
	} else {
		for i := 0; i < 3; i++ {
			
			ID := primitive.NewObjectID().Hex()
			Connections.InsertOne(context.Background(), structs.ChangNumber{
				Id:ID,
				Quantity:0,
			})
		}
	}
}
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



func AddPatientStory(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "Error: No cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error: Only admins have access to add")
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
					fmt.Printf("Folder error: %v\n", Foldererror)
				}
				Foldererror2 := createimagephoto.CreateFolder(folderName2)
				if Foldererror2 != nil {
					fmt.Printf("Folder error: %v\n", Foldererror2)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Photo := baner.ImageFunc(Patient_data.Photo, ForImage, folderName)
				Smallphoto := baner.ImageFunc(Patient_data.Smallphoto, ForImage, folderName2)

				Patient_data.Photo = folderName + "/" + Photo
				Patient_data.Smallphoto = folderName2 + "/" + Smallphoto

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("patientstory")

				Patient_data.Id = primitive.NewObjectID().Hex()

				_, inserterror := createDB.InsertOne(ctx, Patient_data)
				if inserterror != nil {
					fmt.Printf("Insert error: %v\n", inserterror)
				} else {
					c.JSON(200, "Success")
				}
			}
		}

	}
}

func AddPartner(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "Error: No cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error: Only admins have access to add")
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
					fmt.Printf("Folder error: %v\n", Foldererror)
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
					fmt.Printf("Insert error: %v\n", inserterror)
				} else {
					c.JSON(200, "Success")
				}
			}
		}
	}
}


func AddTeamMembers(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		fmt.Printf("Cookie error: %v\n", cookieerror)
		c.JSON(404, "Error: No cookie found")
		fmt.Printf("Cookie data: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error: Only admins have access")
		} else {

			var TeamData structs.Team
			c.ShouldBindJSON(&TeamData)
			Emptyfield, err := emptyfieldcheker.EmptyField(TeamData, "Id","Expirence")

			if Emptyfield {
				c.JSON(400, err)
			} else {
				folderName := "Team"

				FolderError := createimagephoto.CreateFolder(folderName)
				if FolderError != nil {
					fmt.Printf("Folder error: %v\n", FolderError)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Photo := baner.ImageFunc(TeamData.Photo, ForImage, folderName)
				client, ctx := mongoconnect.DBConnection()
				TeamData.Photo = folderName + "/" + Photo
				var createDB = client.Database(env.Data_Name).Collection("team")
				TeamData.Id = primitive.NewObjectID().Hex()
				insertresult, inserterror := createDB.InsertOne(ctx, TeamData)
				if inserterror != nil {
					fmt.Printf("Insert error: %v\n", inserterror)
				} else {
					c.JSON(200, "Success")
					fmt.Printf("Insert result: %v\n", insertresult)
				}
			}
		}

	}
}

func AddProgram(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "Error: No cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error")
		} else {

			var Services structs.Program
			c.ShouldBindJSON(&Services)
			Emptyfield, err := emptyfieldcheker.EmptyField(Services, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				folderName := "Program"

				FolderError := createimagephoto.CreateFolder(folderName)
				if FolderError != nil {
					fmt.Printf("Folder error: %v\n", FolderError)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Photo := baner.ImageFunc(Services.Photo, ForImage, folderName)
				Services.Photo = folderName + "/" + Photo
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database(env.Data_Name).Collection("programs")
				Services.Id = primitive.NewObjectID().Hex()
				insertresult, inserterror := createDB.InsertOne(ctx, Services)
				if inserterror != nil {
					fmt.Printf("Insert error: %v\n", inserterror)
				} else {
					c.JSON(200, "Success")
					fmt.Printf("Insert result: %v\n", insertresult)
				}
			}
		}

	}
}

func AddServices(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "Error: No cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error")
		} else {

			var Services structs.Services
			c.ShouldBindJSON(&Services)
			Emptyfield, err := emptyfieldcheker.EmptyField(Services, "Id")

			if Emptyfield {
				c.JSON(404, err)
			} else {
				folderName := "Services"

				FolderError := createimagephoto.CreateFolder(folderName)
				if FolderError != nil {
					fmt.Printf("Folder error: %v\n", FolderError)
				}
				rndName := rand.Intn(10000)
				ForImage := fmt.Sprintf("image_%v.png", rndName)
				Photo := baner.ImageFunc(Services.Photo, ForImage, folderName)
				Services.Photo = folderName + "/" + Photo
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database(env.Data_Name).Collection("services")
				fmt.Printf("Services data: %v\n", Services)
				Services.Id = primitive.NewObjectID().Hex()
				_, inserterror := createDB.InsertOne(ctx, Services)
				if inserterror != nil {
					fmt.Printf("Insert error: %v\n", inserterror)
				} else {
					c.JSON(200, "Success")
				}
			}
		}

	}
}

func AddStatisticForCenter() {
	client, ctx := mongoconnect.DBConnection()
	Connections := client.Database(env.Data_Name).Collection("center_statistic")
	count, err := Connections.CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	if count >= 3 {
		fmt.Println("too many documents")
	} else {
		for i := 0; i < 3; i++ {
			ID := primitive.NewObjectID().Hex()
			Connections.InsertOne(ctx, structs.ChangNumber{
				Id:       ID,
				Quantity: 0,
			})
		}
	}

}

func AddStatistic() {
	client, ctx := mongoconnect.DBConnection()
	Connections := client.Database(env.Data_Name).Collection("statistic")
	count, err := Connections.CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	if count >= 4 {
		fmt.Println("too many documents")
	} else {
		for i := 0; i < 4; i++ {
			ID := primitive.NewObjectID().Hex()
			var AddStatistic structs.AddStatistic
			AddStatistic.Id = ID
			Connections.InsertOne(ctx, AddStatistic)
		}

	}
}
func AddStatisticForProject() {
	client, ctx := mongoconnect.DBConnection()
	Connections := client.Database(env.Data_Name).Collection("statistic_for_project")
	count, err := Connections.CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	if count >= 3 {
		fmt.Println("too many documents")
	} else {
		for i := 0; i < 3; i++ {
			ID := primitive.NewObjectID().Hex()
			var AddStatistic structs.AddStatisticForCenter
			AddStatistic.Id = ID
			Connections.InsertOne(ctx, AddStatistic)
		}

	}
}

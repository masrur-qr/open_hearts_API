package controlers

import (
	"docs/app/Env"
	"docs/app/mongoconnect"
	returnjwt "docs/app/returnJwt"
	"docs/app/structs"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetStatistics(c *gin.Context) {
	var Forlist = []structs.AddStatistic{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("statistic")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("Error: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.AddStatistic
		fmt.Printf("Data from DB: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetPatientStories(c *gin.Context) {
	var Forlist = []structs.Patient_story{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("patientstory")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("Error: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Patient_story
		fmt.Printf("Data from DB: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetPartners(c *gin.Context) {
	var Forlist = []structs.Partner{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("partners")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("Error: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Partner
		fmt.Printf("Data from DB: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetStatisticsForProject(c *gin.Context) {
	var Forlist = []structs.AddStatisticForCenter{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("statistic_for_project")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("Error: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.AddStatisticForCenter
		fmt.Printf("Data from DB: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetTeam(c *gin.Context) {
	var Forlist = []structs.Team{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("team")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("Error: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Team
		fmt.Printf("Data from DB: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetCenterNumbers(c *gin.Context) {
	var Forlist = []structs.ChangNumber{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("center_statistic")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("Error: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.ChangNumber
		fmt.Printf("Data from DB: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetServices(c *gin.Context) {
	var Forlist = []structs.Services{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("services")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("Error: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Services
		fmt.Printf("Data from DB: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetPrograms(c *gin.Context) {
	var Forlist = []structs.Program{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("programs")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("Error: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Program
		fmt.Printf("Data from DB: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetOnePatient(c *gin.Context) {
	var Forlist = []structs.Patient_story{}
	ids := c.Request.URL.Query().Get("id")
	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("patientstory")

	singlerezult := createDB.FindOne(ctx, bson.M{"_id": ids})
	var datafromdb structs.Patient_story
	singlerezult.Decode(&datafromdb)

	if datafromdb.Id != "" {
		Forlist = append(Forlist, datafromdb)
		c.JSON(200, Forlist)
	} else {
		c.JSON(400, "User not found")
	}
}
func ReadFile(c *gin.Context) {
	path := c.Request.URL.Query().Get("Path")

	fmt.Printf("filename: %v\n", path)
	if path == "" {
		c.JSON(404, "empty filed")
	} else {
		c.File("./Statics/" + path)
	}
}

func GetAdmins(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
			c.JSON(401, "error: No Cookie found")
	}
	SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
	if isvalid || (SecretKeyData.Permission != "MainAdmin" ) {
			c.JSON(403, "error: Unauthorized access")
			return
	}

	client, ctx := mongoconnect.DBConnection()
	collection := client.Database(env.Data_Name).Collection("users")

	projection := bson.D{
			{"password", 0}, 
	}

	result, err := collection.Find(ctx, bson.D{}, options.Find().SetProjection(projection))
	if err != nil {
			fmt.Println("error:", err)
			c.JSON(500, "Error while retrieving data")
			return
	}

	var admins []structs.UserStruct
	for result.Next(ctx) {
			var admin structs.UserStruct
			if err := result.Decode(&admin); err != nil {
					fmt.Println("error decoding result:", err)
					continue
			}
			admins = append(admins, admin)
	}

	if err := result.Err(); err != nil {
			fmt.Println("error iterating over result:", err)
	}

	c.JSON(200, admins)
}


func GetAdmin(c *gin.Context) {
	var Forlist = []structs.UserStruct{}
	ids := c.Request.URL.Query().Get("id")
	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("users")

	singlerezult := createDB.FindOne(ctx, bson.M{"_id": ids})
	var datafromdb structs.UserStruct
	singlerezult.Decode(&datafromdb)

	if datafromdb.Id != "" {
		Forlist = append(Forlist, datafromdb)
		c.JSON(200, Forlist)
	} else {
		c.JSON(400, "User not found")
	}
}

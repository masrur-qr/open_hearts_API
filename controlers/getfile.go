package controlers

import (
	"docs/app/Env"
	"docs/app/mongoconnect"
	"docs/app/structs"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetStatistic(c *gin.Context) {
	var Forlist = []structs.AddStatistic{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("statistic")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.AddStatistic
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}


func GetPatientStory(c *gin.Context) {
	var Forlist = []structs.Patient_story{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("patientstory")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Patient_story
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}


func GetPatners(c *gin.Context) {
	var Forlist = []structs.Partner{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("partners")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Partner
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetStatisticforproject(c *gin.Context) {
	var Forlist = []structs.AddStatisticForCenter{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("statistic_for_project")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.AddStatisticForCenter
		fmt.Printf("datafromdb: %v\n", datafromdb)
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
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Team
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func Get_center_number(c *gin.Context) {
	var Forlist = []structs.ChangNumber{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("center_statistic")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.ChangNumber
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}






func GetServices(c *gin.Context) {
	var Forlist = []structs.Services{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("servisec")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Services
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}
func GetProgram(c *gin.Context) {
	var Forlist = []structs.Services{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("programs")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Services
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}
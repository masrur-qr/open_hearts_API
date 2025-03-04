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
	var createDB = connect.Database(env.Data_Name).Collection("Statistic")

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



func GetStatisticforproject(c *gin.Context) {
	var Forlist = []structs.AddStatisticForProject{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("Statistic")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.AddStatisticForProject
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}
func Get_ChangedNumber_for_project(c *gin.Context) {
	var Forlist = []structs.ChangNumber{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("change_number_for_center")

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



func GetPatientStory(c *gin.Context) {
	var Forlist = []structs.Patient_story{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("PatientStory")

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
	var createDB = connect.Database(env.Data_Name).Collection("Partners")

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

func GetProject(c *gin.Context) {
	var Forlist = []structs.Project{}
	
	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("Project")
	
	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}
	
	for singlerezult.Next(ctx) {
		var datafromdb structs.Project
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)
		
		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetDiraction(c *gin.Context) {
	var Forlist = []structs.Diraction{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("Diraction")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Diraction
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetTeam(c *gin.Context) {
	var Forlist = []structs.Team{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("Team")

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
func GetMainDiraction(c *gin.Context) {
	var Forlist = []structs.MainDiraction{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("MainDiraction")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.MainDiraction
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}





func GetReport(c *gin.Context) {
	var Forlist = []structs.AddReport{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("Report")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.AddReport
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}


func ReadFile(c *gin.Context)  {
	path := c.Request.URL.Query().Get("Path")
	
	fmt.Printf("filename: %v\n", path)
	if path == ""{
		c.JSON(404,"empty filed")
	}else {
		c.File("./Statics/"+path)
	}
}
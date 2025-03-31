package controlers

import (
	env "docs/app/Env"
	"docs/app/emptyfieldcheker"
	"docs/app/mongoconnect"
	"docs/app/returnJwt"
	"docs/app/structs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateStatistic(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "error")
		} else {
			var Update_statistic structs.AddStatistic
			c.ShouldBindJSON(&Update_statistic)

			Emptyfield, err := emptyfieldcheker.EmptyField(Update_statistic, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {

				client, ctx := mongoconnect.DBConnection()
				Connections := client.Database(env.Data_Name).Collection("statistic")
				result := Connections.FindOneAndUpdate(
					ctx,
					bson.D{
						{Key: "_id", Value: Update_statistic.Id},
					},
					bson.D{
						{Key: "$set", Value: bson.D{
							{Key: "quantity", Value: Update_statistic.Quantity},
							{Key: "ru", Value: structs.LangForStatistic{
								Description:Update_statistic.Ru.Description,
							}},
							{Key: "en", Value: structs.LangForStatistic{
								Description:Update_statistic.En.Description,
							}},
							
						}},
					},
				)
				var Dbdata structs.AddStatistic
				result.Decode(&Dbdata)
				if Dbdata.Id == "" {
					c.JSON(400, "statistic not founded")
				} else {
					c.JSON(200, "Success")
				}

			}
		}
	}

}
func UpdateProjectStatistic(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "error")
		} else {
			var Update_statistic_Center structs.AddStatisticForCenter
			c.ShouldBindJSON(&Update_statistic_Center)
			Emptyfield, err := emptyfieldcheker.EmptyField(Update_statistic_Center, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				client, ctx := mongoconnect.DBConnection()
				Connections := client.Database(env.Data_Name).Collection("statistic_for_project")
				result := Connections.FindOneAndUpdate(
					ctx,
					bson.D{
						{Key: "_id", Value: Update_statistic_Center.Id},
					},
					bson.D{
						{Key: "$set", Value: bson.D{
							{Key: "quantity", Value: Update_statistic_Center.Quantity},
							{Key: "ru", Value: structs.LangForProjectStatistic{
								Name:        Update_statistic_Center.Ru.Name,
								Description: Update_statistic_Center.Ru.Description,
							}},
							{Key: "en", Value: structs.LangForProjectStatistic{
								Name:        Update_statistic_Center.En.Name,
								Description: Update_statistic_Center.En.Description,
							}},
						}},
					},
				)
				var Dbdata structs.AddStatisticForCenter
				result.Decode(&Dbdata)
				if Dbdata.Id == "" {
					c.JSON(400, "statistic not founded")
				} else {
					c.JSON(200, "Success")
				}
			}
		}
	}
}

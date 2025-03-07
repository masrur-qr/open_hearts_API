package controlers

import (
	env "docs/app/Env"
	"docs/app/emptyfieldcheker"
	"docs/app/mongoconnect"
	returnjwt "docs/app/returnJwt"
	"docs/app/structs"
	"fmt"

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

package controlers

import (
	env "docs/app/Env"
	htmldata "docs/app/HtmlData"
	"docs/app/emptyfieldcheker"
	"docs/app/hashedpasswod"
	"docs/app/mongoconnect"
	"docs/app/sendmail"
	"docs/app/structs"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func random(min, max int) int {
	if min > max {
		return min
	}
	return rand.Intn(max-min+1) + min
}
func SendSecretCode(c *gin.Context) {

	GetEmail := c.Request.URL.Query().Get("email")

	client, ctx := mongoconnect.DBConnection()
	DBConnect := client.Database(env.Data_Name).Collection("users")

	result := DBConnect.FindOne(ctx, bson.M{
		"email": GetEmail,
	})
	var DBdat structs.UserStruct
	result.Decode(&DBdat)
	if DBdat.Email != "" {
		secretCode := uint64(random(100000, 900000))
		DBdat.Code=secretCode
		data := structs.Verify{
			Id:      string(primitive.NewObjectID().Hex()),
			Code:    DBdat.Code,
			Email:   DBdat.Email,
			User_Id: DBdat.Id,
		}
		htmldata.HtmlData(data)
		id2 := primitive.NewObjectID().Hex()
		sendError := sendmail.SendGomail("./controlers/htmls/output.html", "Your secret code", DBdat.Email)
		if sendError != nil {
			fmt.Printf("sendError: %v\n", sendError)
		} else {
			DBConnect2 := client.Database(env.Data_Name).Collection("code")
			_, err := DBConnect2.InsertOne(ctx, bson.M{
				"_id":     id2,
				"code":    DBdat.Code,
				"email":   DBdat.Email,
				"user_Id": DBdat.Id,
			})
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			time.AfterFunc(60*time.Second, func() {
				fmt.Printf("Attempting to delete code with ID: %s\n", id2)
				Delete(id2)
			})
			c.JSON(200, "succes")
		}
	} else {
		c.JSON(404, "error email not found")
	}
}

func Delete(ids string) {
	client, ctx := mongoconnect.DBConnection()
	defer client.Disconnect(ctx)

	connect2 := client.Database(env.Data_Name).Collection("code")
	deleteResult, deleteError := connect2.DeleteOne(ctx, bson.M{
		"_id": ids,
	})

	if deleteError != nil {
		fmt.Printf("Delete error: %v\n", deleteError)
		return
	}

	if deleteResult.DeletedCount == 1 {
		fmt.Println("Code successfully deleted")
	} else {
		fmt.Println("Code not found")
	}
}

func UpdateAdminPassword(c *gin.Context) {
	var Update_Pass structs.UserStruct
	c.ShouldBindJSON(&Update_Pass)
	Emptyfield, err := emptyfieldcheker.EmptyField(Update_Pass, "Email","Id","Photo","Ru","En","Name","Phone","Password","Permission")
	if Emptyfield {
		c.JSON(404, err)
	} else {
		client, ctx := mongoconnect.DBConnection()
		DBConnect := client.Database(env.Data_Name).Collection("code")
		result := DBConnect.FindOne(ctx, bson.M{
			"code": Update_Pass.Code,
		})
		var DBdat structs.UserStruct
		result.Decode(&DBdat)

		if DBdat.Code != 0 {
			digits := "0123456789"
			specials := "~%//!@#$?|"
			all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
				"abcdefghijklmnopqrstuvwxyz" +
				digits + specials
			length := 9
			buf := make([]byte, length)
			buf[0] = digits[rand.Intn(len(digits))]
			buf[1] = specials[rand.Intn(len(specials))]
			for i := 2; i < length; i++ {
				buf[i] = all[rand.Intn(len(all))]
			}
			rand.Shuffle(len(buf), func(i, j int) {
				buf[i], buf[j] = buf[j], buf[i]
			})
			str := string(buf)
			hashpass, _ := hashedpasswod.HashPassword(str)

			DBConnect2 := client.Database(env.Data_Name).Collection("users")
			result := DBConnect2.FindOne(ctx, bson.M{
				"email": DBdat.Email,
			})
			var DBdat2 structs.UserStruct
			result.Decode(&DBdat2)
			if DBdat.Email != "" {
				_, err := DBConnect2.UpdateOne(ctx,
					bson.M{
						"email": DBdat.Email,
					},
					bson.D{
						{Key: "$set", Value: bson.M{
							"password": hashpass,
						},
						},
					},
				)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}
				DBdat2.Password=str
				htmlerror := htmldata.HtmlData2(DBdat2)
				if htmlerror != nil {
					fmt.Printf("htmlerror: %v\n", htmlerror)
				} else {
					sendError := sendmail.SendGomail("./controlers/htmls/passwordoutput.html", "Your new password", DBdat.Email)
					if sendError != nil {
						fmt.Printf("sendError: %v\n", sendError)
					} else {
						c.JSON(200, "success")
					}
				}
			} else {
				c.JSON(404, "error email not found")
			}
		} else {
			c.JSON(404, "can't find the email")
		}

	}

}

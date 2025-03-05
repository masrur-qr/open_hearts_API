package controlers

import (
	env "docs/app/Env"
	"docs/app/emptyfieldcheker"
	"docs/app/hashedpasswod"
	"docs/app/mongoconnect"
	returnjwt "docs/app/returnJwt"
	"docs/app/structs"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *gin.Context) {
	var LoginTemp structs.UserStruct
	c.ShouldBindJSON(&LoginTemp)
	EmptyField, err := emptyfieldcheker.EmptyField(LoginTemp, "Photo", "Name", "Surname", "Email", "Id", "Permission", "Ru", "En", "Audience", "Issuer", "Subject")
	if EmptyField {
		c.JSON(404, err)
	} else {
		client, ctx := mongoconnect.DBConnection()

		DBConnect := client.Database(env.Data_Name).Collection("Users")

		result := DBConnect.FindOne(ctx, bson.M{
			"phone": LoginTemp.Phone,
		})

		var userdata structs.UserStruct
		result.Decode(&userdata)
		isValidPass := hashedpasswod.CompareHashPasswords(userdata.Password, LoginTemp.Password)
		fmt.Println(isValidPass)
		key := returnjwt.GenerateToken(userdata.Phone, userdata.Permission, userdata.Id)
		if userdata.Phone == LoginTemp.Phone {
			if isValidPass {
				http.SetCookie(c.Writer, &http.Cookie{
					Name:     env.Data_Name,
					Value:    key,
					Expires:  time.Now().Add(60 * time.Hour),
					Domain:   "",
					Path:     "/",
					Secure:   false,
					HttpOnly: false,
					SameSite: http.SameSiteLaxMode,
				})
				c.JSON(200, "success")
			} else {
				c.JSON(404, "Not valid pass")
			}
		} else {
			c.JSON(400, "Wrong phone")
		}
	}

}
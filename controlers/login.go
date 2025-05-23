package controlers

import (
	env "docs/app/Env"
	"docs/app/emptyfieldcheker"
	"docs/app/hashedpasswod"
	"docs/app/mongoconnect"
	returnjwt "docs/app/returnJwt"
	"docs/app/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *gin.Context) {
	var LoginTemp structs.UserStruct
	c.ShouldBindJSON(&LoginTemp)
	EmptyField, err := emptyfieldcheker.EmptyField(LoginTemp,"Code", "Photo", "Name", "Surname", "Email", "Id", "Permission", "Ru", "En")
	if EmptyField {
		c.JSON(400, err)
	} else {
		client, ctx := mongoconnect.DBConnection()

		DBConnect := client.Database(env.Data_Name).Collection("users")

		result := DBConnect.FindOne(ctx, bson.M{
			"phone": LoginTemp.Phone,
		})
		var userdata structs.UserStruct
		result.Decode(&userdata)
		isValidPass := hashedpasswod.CompareHashPasswords(userdata.Password, LoginTemp.Password)
		key := returnjwt.GenerateToken(userdata.Phone, userdata.Permission, userdata.Id)
		if userdata.Phone != "" && isValidPass {
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     env.Data_Cockie,
				Value:    key,
				Expires:  time.Now().Add(60 * time.Hour),
				Domain:   "",
				Path:     "/",
				Secure:   false,
				HttpOnly: false,
				SameSite: http.SameSiteLaxMode,
			})
		admindata:=structs.UserStruct{
			Id: userdata.Id,
			Photo: userdata.Photo,
			Ru: structs.LangForUser{Name: userdata.Ru.Name},
			En: structs.LangForUser{Name: userdata.En.Name},
			Email: userdata.Email,
			Phone: userdata.Phone,
			Permission: userdata.Permission,
		}
		c.JSON(200,admindata)
		} else {
			c.JSON(401, "Access is blocked")
		}
	}

}

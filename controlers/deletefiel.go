package controlers

import (
	"docs/app/Env"
	"docs/app/mongoconnect"
	"docs/app/returnJwt"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)



func DeletePatientStory(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error")
		} else {
			ids := c.Request.URL.Query().Get("id")
			Path := c.Request.URL.Query().Get("Path")
			Path2 := c.Request.URL.Query().Get("SecondPath")

			if ids == "" && Path == "" && Path2=="" {
				c.JSON(404, "Empty Field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("patientstory")
				deletrezult, deleteerror := createDB.DeleteOne(ctx, bson.M{
					"_id": ids,
				})
				if deleteerror != nil {
					fmt.Printf("deleteerror: %v\n", deleteerror)
				}
				if deletrezult.DeletedCount == 1 {
					err := os.RemoveAll("./Statics/" + Path)
					if err != nil {
						fmt.Printf("err: %v\n", err)
					}
					err2 := os.RemoveAll("./Statics/" + Path2)
					if err != nil && err2!=nil {
						fmt.Printf("err: %v\n", err)
					}
					c.JSON(200, "success")
					fmt.Printf("deletrezult: %v\n", deletrezult)
				} else {
					c.JSON(404, "error")
				}
			}
		}
	}
}





func DeletePatners(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error")
		} else {
			ids := c.Request.URL.Query().Get("id")
			Path := c.Request.URL.Query().Get("Path")

			client, ctx := mongoconnect.DBConnection()
			var createDB = client.Database(env.Data_Name).Collection("partners")
			deletrezult, deleteerror := createDB.DeleteOne(ctx, bson.M{
				"_id": ids,
			})
			if deleteerror != nil {
				fmt.Printf("deleteerror: %v\n", deleteerror)
			}
			if deletrezult.DeletedCount == 1 {
				err := os.RemoveAll("./Statics/" + Path)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				} else {
					c.JSON(200, "success")
					fmt.Printf("deletrezult: %v\n", deletrezult)
				}
			} else {
				c.JSON(404, "error1")
			}
		}
	}
}




func DeleteTeam(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error")
		} else {

			ids := c.Request.URL.Query().Get("id")
			Path := c.Request.URL.Query().Get("Path")


			if ids == "" && Path == "" {
				c.JSON(404, "error empty field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("team")
				deletrezult, deleteerror := createDB.DeleteOne(ctx, bson.M{
					"_id": ids,
				})
				if deleteerror != nil {
					fmt.Printf("deleteerror: %v\n", deleteerror)
				}
				if deletrezult.DeletedCount == 1 {
					err := os.RemoveAll("./Statics/" + Path)
					if err != nil {
						fmt.Printf("err: %v\n", err)
					} else {
						c.JSON(200, "success")
					}
				} else {
					c.JSON(404, "error person not deleted")
				}
			}
		}
	}
}
func DeleteServisec(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error")
		} else {

			ids := c.Request.URL.Query().Get("id")
			Path := c.Request.URL.Query().Get("Path")


			if ids == "" && Path == "" {
				c.JSON(404, "error empty field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("services")
				deletrezult, deleteerror := createDB.DeleteOne(ctx, bson.M{
					"_id": ids,
				})
				if deleteerror != nil {
					fmt.Printf("deleteerror: %v\n", deleteerror)
				}
				if deletrezult.DeletedCount == 1 {
					err := os.RemoveAll("./Statics/" + Path)
					if err != nil {
						fmt.Printf("err: %v\n", err)
					} else {
						c.JSON(200, "success")
					}
				} else {
					c.JSON(404, "error")
				}
			}
		}
	}
}

func DeleteProgram(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error")
		} else {

			ids := c.Request.URL.Query().Get("id")
			Path := c.Request.URL.Query().Get("Path")


			if ids == "" && Path == "" {
				c.JSON(404, "error empty field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("programs")
				deletrezult, deleteerror := createDB.DeleteOne(ctx, bson.M{
					"_id": ids,
				})
				if deleteerror != nil {
					fmt.Printf("deleteerror: %v\n", deleteerror)
				}
				if deletrezult.DeletedCount == 1 {
					err := os.RemoveAll("./Statics/" + Path)
					if err != nil {
						fmt.Printf("err: %v\n", err)
					} else {
						c.JSON(200, "success")
					}
				} else {
					c.JSON(404, "error")
				}
			}
		}
	}
}
func DeleteAdmin(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		c.JSON(404, "error Not Cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "error")
		} else {

			ids := c.Request.URL.Query().Get("id")
			Path := c.Request.URL.Query().Get("Path")


			if ids == "" && Path == "" {
				c.JSON(404, "error empty field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("users")
				deletrezult, deleteerror := createDB.DeleteOne(ctx, bson.M{
					"_id": ids,
				})
				if deleteerror != nil {
					fmt.Printf("deleteerror: %v\n", deleteerror)
				}
				if deletrezult.DeletedCount == 1 {
					err := os.RemoveAll("./Statics/" + Path)
					if err != nil {
						fmt.Printf("err: %v\n", err)
					} else {
						c.JSON(200, "success")
					}
				} else {
					c.JSON(404, "error")
				}
			}
		}
	}
}











package controlers

import (
	"docs/app/Env"
	"docs/app/mongoconnect"
	"docs/app/returnJwt"
	"docs/app/structs"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeletePatientStory(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "Error: No cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error: You are not an admin")
		} else {
			ids := c.Request.URL.Query().Get("id")

			if ids == "" {
				c.JSON(404, "Empty field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("patientstory")
				deleterezult := createDB.FindOneAndDelete(ctx, bson.M{"_id": ids})
				var Delete structs.Patient_story
				deleterezult.Decode(&Delete)
				if Delete.Id != "" {
					err := os.RemoveAll("./Statics/" + Delete.Photo)
					err2 := os.RemoveAll("./Statics/" + Delete.Smallphoto)
					if err != nil && err2 != nil {
						fmt.Printf("Error: %v\n", err)
					} else {
						c.JSON(200, "Success")
					}
				} else {
					c.JSON(404, "Error: Not found")
				}
			}
		}
	}
}

func DeletePartners(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "Error: No cookie found")

	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error: You are not an admin")
		} else {
			ids := c.Request.URL.Query().Get("id")

			client, ctx := mongoconnect.DBConnection()
			var createDB = client.Database(env.Data_Name).Collection("partners")
			deleterezult := createDB.FindOneAndDelete(ctx, bson.M{"_id": ids})
			var Delete structs.Partner
			deleterezult.Decode(&Delete)
			if Delete.Id != "" {
				err := os.RemoveAll("./Statics/" + Delete.Logo)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				} else {
					c.JSON(200, "Success")
				}
			} else {
				c.JSON(404, "Error: Not found")
			}
		}
	}
}

func DeleteTeam(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {

		c.JSON(404, "Error: No cookie found")

	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error: You are not an admin")
		} else {

			ids := c.Request.URL.Query().Get("id")

			if ids == "" {
				c.JSON(404, "Error: Empty field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("team")
				deleterezult := createDB.FindOneAndDelete(ctx, bson.M{"_id": ids})
				var Delete structs.Team
				deleterezult.Decode(&Delete)
				if Delete.Id != "" {
					err := os.RemoveAll("./Statics/" + Delete.Photo)
					if err != nil {
						fmt.Printf("Error: %v\n", err)
					} else {
						c.JSON(200, "Success")
					}
				} else {
					c.JSON(404, "Error: Not found")
				}
			}
		}
	}
}

func DeleteServices(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {

		c.JSON(404, "Error: No cookie found")

	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error: You are not an admin")
		} else {

			ids := c.Request.URL.Query().Get("id")

			if ids == "" {
				c.JSON(404, "Error: Empty field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("services")
				deleterezult := createDB.FindOneAndDelete(ctx, bson.M{"_id": ids})
				var Delete structs.UserStruct
				deleterezult.Decode(&Delete)
				if Delete.Id != "" {
					err := os.RemoveAll("./Statics/" + Delete.Photo)
					if err != nil {
						fmt.Printf("Error: %v\n", err)
					} else {
						c.JSON(200, "Success")
					}
				} else {
					c.JSON(404, "Error: Not found")
				}
			}
		}
	}
}

func DeleteProgram(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {

		c.JSON(404, "Error: No cookie found")

	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error: You are not an admin")
		} else {

			ids := c.Request.URL.Query().Get("id")

			if ids == "" {
				c.JSON(404, "Error: Empty field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("programs")
				deleterezult := createDB.FindOneAndDelete(ctx, bson.M{"_id": ids})
				var Delete structs.Program
				deleterezult.Decode(&Delete)
				if Delete.Id != "" {
					err := os.RemoveAll("./Statics/" + Delete.Photo)
					if err != nil {
						fmt.Printf("Error: %v\n", err)
					} else {
						c.JSON(200, "Success")
					}
				} else {
					c.JSON(404, "Error: Not found")
				}
			}
		}
	}
}

func DeleteAdmin(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Cockie)
	if cookieerror != nil {
		c.JSON(404, "Error: No cookie found")
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "MainAdmin" && isvalid {
			c.JSON(404, "Error: You are not a Main Admin")
		} else {

			ids := c.Request.URL.Query().Get("id")

			if ids == "" {
				c.JSON(404, "Error: Empty field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("users")
				deleterezult := createDB.FindOneAndDelete(ctx, bson.M{"_id": ids})
				var Delete structs.UserStruct
				deleterezult.Decode(&Delete)
				if Delete.Id != "" {
					err := os.RemoveAll("./Statics/" + Delete.Photo)
					if err != nil {
						fmt.Printf("Error: %v\n", err)
					} else {
						c.JSON(200, "Success")
					}
				} else {
					c.JSON(404, "Error: Not found")
				}
			}
		}
	}
}

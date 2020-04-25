package controllers

import (
	"fmt"
	"net/http"

	"github.com/M-Gadd/photos-project/server/database"
	"github.com/M-Gadd/photos-project/server/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func MongoInit() *mgo.Database {
	db, err := database.Init()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func GetUser(c *gin.Context) {

	db := *MongoInit()
	defer db.Session.Close()

	user := models.User{}
	err := db.C("users").Find(bson.M{"name": "Prathamesh Sonpatki"}).One(&user)

	fmt.Println("Email Id:", user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Get User",
		})
		// c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"user": &user.Name,
	})

}

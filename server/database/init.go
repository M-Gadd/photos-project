package database

import (
	"fmt"
	"log"
	"os"

	"github.com/M-Gadd/photos-project/server/models"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

// type DataStore struct {
// 	session *mgo.Session
// }

var session *mgo.Session

func Init() (*mgo.Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		// fmt.Println("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
		// return nil, err
	}

	session, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		return nil, err
	}
	// defer session.Close()

	session.SetSafe(&mgo.Safe{})

	xxx, err := session.DatabaseNames()
	if err != nil {
		// panic(err)
		// return nil, err
		fmt.Println("Can't get databases")
	}
	fmt.Println("I AM DATABASE NAMES:", xxx)

	// db := session.DB("family-photos")
	db := session.DB("heroku_46r360r3")

	err = db.C("users").Insert(&models.User{"Stefan Klaste", "klaste@posteo.de"},
		&models.User{"Nishant Modak", "modak.nishant@gmail.com"},
		&models.User{"Prathamesh Sonpatki", "csonpatki@gmail.com"},
		&models.User{"murtuza kutub", "murtuzafirst@gmail.com"},
		&models.User{"aniket joshi", "joshianiket22@gmail.com"},
		&models.User{"Michael de Silva", "michael@mwdesilva.com"},
		&models.User{"Alejandro Cespedes Vicente", "cesal_vizar@hotmail.com"})
	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return nil, err
	}

	return db, nil
}

// CloseSession ...
// func (d DataStore) CloseSession() {
// 	fmt.Print("I am closing the freaking session")
// 	d.session.Clone()
// }
// CloseSession ...
// func CloseSession() {
// 	fmt.Print("I am closing the freaking session")
// 	session.Clone()
// }

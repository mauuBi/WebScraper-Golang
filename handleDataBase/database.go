package handledatabase

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Articles struct{
	Title string
	Url string
	Author string
	Topics string
	LinkOfImage string
}

func createDataBase() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("articles.db"), &gorm.Config{})
	if err != nil{
		log.Fatalln(err)
	}
	db.AutoMigrate(&Articles{})
	return db
}

func CreatePosts(NewPostInDB *Articles){
	db := createDataBase()
	res := db.Create(&NewPostInDB)
	if res.Error != nil{
		log.Fatalln(res.Error)
	}
}
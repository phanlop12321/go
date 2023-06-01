package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/phanlop12321/golang/db"
	"github.com/phanlop12321/golang/handler"
)

// type CourseJSON struct {
// 	ID          string `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"desc"`
// }

// var courses = []Course{
// 	{ID: "1", Name: "TDD", Description: "Let's make a bug free software!"},
// 	{ID: "2", Name: "CI&CD", Description: "Continue deliver a customer value"},
// }

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	r.GET("/courses", handler.ListCourses(db))
	r.GET("/courses/:id", handler.GetCourses(db))
	r.POST("/courses", handler.CreateCourses(db))
	r.POST("/classes", handler.CreateClasses(db))
	r.POST("/enroll", handler.EnrollClass(db))
	r.POST("/register", handler.Register(db))

	r.Run(":8080")
}

// func Error(c *gin.Context, status int, err error) {
// 	log.Println(err)
// 	c.JSON(status, gin.H{
// 		"message": err.Error(),
// 	})
// }

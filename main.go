package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phanlop12321/golang/db"
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

	r.GET("/courses", listCourses(db))
	r.GET("/courses/:id", getCourses(db))

	r.Run(":8080")
}

func listCourses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courses, err := db.GetAllCourse()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "server error!",
			})
			return
		}
		c.IndentedJSON(http.StatusOK, courses)
	}
}
func getCourses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "invalid id",
			})
			return
		}
		course, err := db.GetCourse(uint(id))
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"message": "no course of a given ID",
			})
			return
		}
		c.IndentedJSON(http.StatusOK, course)
	}
}

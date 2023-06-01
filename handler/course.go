package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phanlop12321/golang/db"
	"github.com/phanlop12321/golang/model"
)

func ListCourses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courses, err := db.GetAllCourse()
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, courses)
	}
}
func GetCourses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		course, err := db.GetCourse(uint(id))
		if err != nil {
			Error(c, http.StatusNotFound, err)
			return
		}
		c.IndentedJSON(http.StatusOK, course)
	}
}
func CreateCourses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(model.Course)
		if err := c.BindJSON(req); err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		if err := db.CreateCourse(*req); err != nil {
			Error(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, req)
	}
}

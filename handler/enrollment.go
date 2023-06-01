package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phanlop12321/golang/db"
)

type EnrollmentReq struct {
	StudentID uint `json:"student_id"`
	ClassID   uint `json:"class_id"`
}

func EnrollClass(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(EnrollmentReq)
		if err := c.BindJSON(req); err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		class, err := db.GetClass(req.ClassID)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}

		student, err := db.GetStudent(req.StudentID)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		if err := class.AddStudent(*student); err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		if err := db.CreateClassStudent(student.ID, class.ID); err != nil {
			Error(c, http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

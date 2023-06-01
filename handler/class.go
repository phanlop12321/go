package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phanlop12321/golang/db"
)

type ClassReq struct {
	ID        uint      `json:"id"`
	CourseID  uint      `json:"course_id"`
	TrainerID uint      `json:"train_id"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Seats     int       `json:"seats"`
}

func CreateClasses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(ClassReq)
		if err := c.BindJSON(req); err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		course, err := db.GetCourse(req.CourseID)
		if err != nil {
			Error(c, http.StatusNotFound, err)
			return
		}
		class, err := course.CreateClass(req.Start, req.End)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		if err := class.SetSeats(req.Seats); err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		class.Trainer.ID = req.TrainerID

		if err := db.SaveClass(class); err != nil {
			Error(c, http.StatusInternalServerError, err)
			return
		}
		c.Status(http.StatusOK)
	}
}

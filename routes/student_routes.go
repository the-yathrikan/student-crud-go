package routes

import (
	"student-api/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	studentGroup := r.Group("/students")
	{
		studentGroup.POST("/", controllers.CreateStudent)
		studentGroup.GET("/", controllers.GetAllStudents)
		studentGroup.GET("/:id", controllers.GetStudentByID)
		studentGroup.PUT("/:id", controllers.UpdateStudent)
		studentGroup.DELETE("/:id", controllers.DeleteStudent)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"

	"admission-portal-backend/internal/controllers"
	"admission-portal-backend/internal/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/api/students/signup", controllers.Signup)
	router.POST("/api/students/login", controllers.Login)
	router.POST("/api/students/create-admin", controllers.CreateAdmin)

	// Protected routes
	authorized := router.Group("/api")
	authorized.Use(middlewares.AuthMiddleware())
	{
		// Student routes
		authorized.GET("/students/me", controllers.GetProfile)
		authorized.PUT("/students/me", controllers.UpdateProfile)
		authorized.GET("/students/admins", middlewares.AdminOnly(), controllers.ListAdmins)

		// Course routes
		authorized.POST("/courses", middlewares.AdminOnly(), controllers.CreateCourse)
		authorized.GET("/courses", controllers.GetCourses)
		authorized.GET("/courses/:id", controllers.GetCourse)
		authorized.PUT("/courses/:id", middlewares.AdminOnly(), controllers.UpdateCourse)
		authorized.DELETE("/courses/:id", middlewares.AdminOnly(), controllers.DeleteCourse)

		// Admission routes
		authorized.POST("/admissions", controllers.ApplyAdmission)
		authorized.GET("/admissions", controllers.GetAdmissions)
		authorized.GET("/admissions/:id", controllers.GetAdmission)
		authorized.PUT("/admissions/:id", middlewares.AdminOnly(), controllers.UpdateAdmissionStatus)
	}
}

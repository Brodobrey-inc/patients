package webserver

import (
	"github.com/gin-gonic/gin"
	"patients/webserver/controllers"
)

func Initialize() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("patients", controllers.ListPatients)

	patientGroup := r.Group("patient/")
	{
		patientGroup.POST("new", controllers.NewPatient)
		patientGroup.POST("edit", controllers.EditPatient)
		patientGroup.POST("delete", controllers.DeletePatient)
	}

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

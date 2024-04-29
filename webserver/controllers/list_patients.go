package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patients/database"
)

func ListPatients(c *gin.Context) {
	c.JSON(http.StatusOK, database.PatientsData.ListPatients())
}

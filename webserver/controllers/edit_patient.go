package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patients/database"
	"patients/database/structs"
	"patients/logging"
)

func EditPatient(c *gin.Context) {
	var patient structs.Patient
	if err := c.BindJSON(&patient); err != nil {
		logging.LogError(err, "Can`t unmarshal json to struct")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.PatientsData.EditPatient(patient); err != nil {
		logging.LogError(err, "Failed update patient info")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := database.PatientsData.Save(); err != nil {
		logging.LogError(err, "Failed save new data to disk")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

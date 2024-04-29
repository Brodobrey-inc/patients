package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"patients/database"
	"patients/database/structs"
	"patients/logging"
)

func NewPatient(c *gin.Context) {
	var patient structs.Patient
	if err := c.BindJSON(&patient); err != nil {
		logging.LogError(err, "Can`t unmarshal json to struct")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient.GUID = uuid.New()
	database.PatientsData.AddPatient(patient)

	err := database.PatientsData.Save()
	if err != nil {
		logging.LogError(err, "Failed save new data to disk")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"guid": patient.GUID})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"patients/database"
	"patients/logging"
)

type DeleteRequest struct {
	GUID uuid.UUID `json:"guid"`
}

func DeletePatient(c *gin.Context) {
	var input DeleteRequest
	if err := c.BindJSON(&input); err != nil {
		logging.LogError(err, "Can`t unmarshal json to struct")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.PatientsData.DeletePatient(input.GUID); err != nil {
		logging.LogError(err, "Failed delete patient info")
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

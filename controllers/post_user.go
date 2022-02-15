package controllers

import (
	"net/http"

	"github.com/ownfitness/template-go/pkg/gcp"

	"cloud.google.com/go/firestore"

	"github.com/ownfitness/template-go/models"

	"github.com/gin-gonic/gin"
)

func PostUser(client *firestore.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.User

		if err := c.ShouldBindJSON(&data); err != nil {
			errorHandler(c, err.Error(), http.StatusBadRequest)
			return
		}

		valid, reason := data.Validate()

		if !valid {
			errorHandler(c, reason, http.StatusBadRequest)
			return
		}

		id, err := gcp.FirestoreSetUser(c, client, "users", data)
		if err != nil {
			errorHandler(c, err.Error(), http.StatusInternalServerError)
			return
		}

		data.Id = id

		c.JSON(http.StatusAccepted, data)
	}
}

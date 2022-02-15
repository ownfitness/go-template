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

		if err := c.ShouldBind(&data); err != nil {
			errorHandler(c, err.Error(), http.StatusBadRequest)
		}

		valid, reason := data.Validate()

		if !valid {
			errorHandler(c, reason, http.StatusBadRequest)
		}

		id, err := gcp.FirestoreSetUser(c, client, "users", data)
		if err != nil {
			errorHandler(c, err.Error(), http.StatusInternalServerError)
		}

		data.Id = id

		c.JSON(http.StatusAccepted, data)
	}
}

package gcp

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/ownfitness/template-go/models"
)

func FirestoreSetUser(ctx context.Context, client *firestore.Client, collection string, data models.User) (string, error) {
	id := encodeID(data.Email)

	data.Id = id

	if _, err := client.Collection(collection).Doc(id).Set(ctx, data); err != nil {
		return "", fmt.Errorf("error creating or updating document: %s", err.Error())
	}

	return id, nil
}

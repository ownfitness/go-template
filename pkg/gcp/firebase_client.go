package gcp

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"

	firebase "firebase.google.com/go/v4"
)

type Firestore struct {
	Firestore *firestore.Client
}

func FirebaseClient(project string) (*Firestore, error) {
	ctx := context.Background()

	config := &firebase.Config{ProjectID: project}
	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %s", err.Error())
	}

	fsc, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing firestore client: %s", err.Error())
	}

	return &Firestore{
		Firestore: fsc,
	}, nil
}

package providers

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func GetFirebase() *messaging.Client {
	ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		panic("Firebase load error")
	}

	return client
}

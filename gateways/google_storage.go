package gateways

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

func ReadContentFile(bucket string, object string) (string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		return "", err
	}

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)

	if err != nil {
		return "", err
	}
	slurp, err := ioutil.ReadAll(rc)
	rc.Close()

	if err != nil {
		return "", err
	}

	return string(slurp), nil
}

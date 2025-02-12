package service

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

const (
	SCOPE_READ_WRITE = "https://www.googleapis.com/auth/devstorage.read_write"
)

func UploadImages(ctx context.Context, objectName string, postID uint, file io.Reader) (string, error) {
	godotenv.Load()
	var bucket string = os.Getenv("BUCKET")
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("../config/donde-caigo-448902-cba3eb3a1f3b.json"))
	if err != nil {
		return "", fmt.Errorf("error al crear cliente...\nDetalles: %v", err)
	}
	defer client.Close()
	wc := client.Bucket(bucket).Object(fmt.Sprintf("%d_%s", postID, objectName)).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("error al subir archivo...\nDetalles: %v", err)
	}
	wc.Close()
	var imageURL string = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, objectName)
	return imageURL, nil
}

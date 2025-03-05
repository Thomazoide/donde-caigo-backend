package middleware

import (
	"bytes"
	"context"
	"os"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type PicturesMiddleware struct {
	client *storage.Client
	bucket string
}

func NewPicturesMiddleware() *PicturesMiddleware {
	if envErr := godotenv.Load(); envErr != nil {
		return nil
	}
	bucket := os.Getenv("BUCKET")
	if bucket == "" {
		return nil
	}
	credentialsFile := os.Getenv("CREDENTIALS_FILE")
	if credentialsFile == "" {
		return nil
	}
	client, clientErr := storage.NewClient(context.Background(), option.WithCredentialsFile(credentialsFile))
	if clientErr != nil {
		return nil
	}
	return &PicturesMiddleware{
		client: client,
		bucket: bucket,
	}
}

func (pm *PicturesMiddleware) UploadImage(ctx context.Context, fileName string, picture []byte) (string, error) {
	writer := pm.client.Bucket(pm.bucket).Object(fileName).NewWriter(ctx)
	writer.ContentType = "image/jpeg"
	defer writer.Close()
	if _, writeErr := writer.Write(picture); writeErr != nil {
		return "", writeErr
	}
	pictureURL := "https://storage.googleapis.com/" + pm.bucket + "/" + fileName
	return pictureURL, nil
}

func (pm *PicturesMiddleware) GetImage(ctx context.Context, fileName string) ([]byte, error) {
	reader, readErr := pm.client.Bucket(pm.bucket).Object(fileName).NewReader(ctx)
	if readErr != nil {
		return nil, readErr
	}
	defer reader.Close()
	var buffer bytes.Buffer
	if _, copyErr := buffer.ReadFrom(reader); copyErr != nil {
		return nil, copyErr
	}
	return buffer.Bytes(), nil
}

func (pm *PicturesMiddleware) DeleteImage(ctx context.Context, fileName string) error {
	if err := pm.client.Bucket(pm.bucket).Object(fileName).Delete(ctx); err != nil {
		return err
	}
	return nil
}

func (pm *PicturesMiddleware) Base64ToImage(base64 string) ([]byte, error) {
	return []byte(base64), nil
}

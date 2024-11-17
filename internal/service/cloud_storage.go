package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type CloudStorageService struct {
	Client *storage.Client
}

func InitCloudStorageClient() *CloudStorageService {
	client, err := storage.NewClient(context.Background())

	if err != nil {
		panic(fmt.Sprintf("Failed to init cloud storage client: %v", err.Error()))
	}

	return &CloudStorageService{
		Client: client,
	}
}

func (cs *CloudStorageService) UploadToBucket(ctx context.Context, bucketName string, file multipart.File, filename string) error {
	bucketHandler := cs.Client.Bucket(bucketName)
	objectHandler := bucketHandler.Object(filename)
	writer := objectHandler.NewWriter(ctx)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(fileBytes)
	err = writer.Close()
	if err != nil {
		return err
	}

	return nil
}

package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/net/context"
	"io"
	"log/slog"
	"os"
	"project-auction/internal/config"
)

type MinioStorage interface {
	Bucket()
	Upload(fileName string) error
	Object(fileName string) ([]byte, error)
}

type minioStorage struct {
	minioClient *minio.Client
	bucketName  string
	log         *slog.Logger
}

func NewMinioStorage(log *slog.Logger, cfg config.Config) MinioStorage {
	minioClient, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.MinioAccessKey, cfg.MinioSecretAccessKey, ""),
	})
	if err != nil {
		panic("failed to connect storage: " + err.Error())
	}

	return &minioStorage{
		minioClient: minioClient,
		log:         log,
		bucketName:  cfg.MinioBucketName,
	}
}

func (m *minioStorage) Bucket() {
	location := "us-east-1"

	if err := m.minioClient.MakeBucket(context.TODO(), m.bucketName, minio.MakeBucketOptions{Region: location}); err != nil {
		exists, errBucketExists := m.minioClient.BucketExists(context.TODO(), m.bucketName)
		if errBucketExists == nil && exists {
			m.log.Warn("bucket already exists", slog.String("bucket", m.bucketName))
		} else {
			panic("failed to make bucket: " + err.Error())
		}
	} else {
		m.log.Info("successfully created bucket", slog.String("bucket", m.bucketName))
	}
}

func (m *minioStorage) Upload(fileName string) error {
	contentType := "application/json"

	info, err := m.minioClient.FPutObject(context.TODO(), m.bucketName, fileName, fileName, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		m.log.Error("failed to put file", slog.String("fileName", fileName), slog.String("error", err.Error()))
		return err
	}

	m.log.Info("successfully uploaded", slog.String("fileName", fileName), slog.Int64("size", info.Size))

	if err := os.Remove(fileName); err != nil {
		m.log.Error("failed to remove file", slog.String("fileName", fileName), slog.String("error", err.Error()))
	}

	return nil
}

func (m *minioStorage) Object(fileName string) ([]byte, error) {
	obj, err := m.minioClient.GetObject(context.TODO(), m.bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		m.log.Error("failed to get object", slog.String("error", err.Error()))
		return nil, err
	}

	objBytes, err := io.ReadAll(obj)
	if err != nil {
		m.log.Error("failed to read obj", slog.String("error", err.Error()))
		return nil, err
	}

	return objBytes, nil
}

package awss3

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	bucketName = "clubya-archivos"
	region     = "us-east-1"
)

func UploadToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(file); err != nil {
		return "", err
	}

	key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), fileHeader.Filename)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
	})

	if err != nil {
		return "", err
	}

	// URL pública del archivo subido
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
	return url, nil
}

func UploadFromBytes(fileData []byte, fileName string, contentType string) (string, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
    if err != nil {
        return "", fmt.Errorf("error al cargar configuración AWS: %w", err)
    }

    client := s3.NewFromConfig(cfg)

    // Generar clave única para el archivo
    key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), fileName)

    // Si no se proporciona contentType, usar uno por defecto
    if contentType == "" {
        contentType = "application/octet-stream"
    }

    _, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket:      aws.String(bucketName),
        Key:         aws.String(key),
        Body:        bytes.NewReader(fileData),
        ContentType: aws.String(contentType),
    })

    if err != nil {
        return "", fmt.Errorf("error al subir archivo a S3: %w", err)
    }

    // URL pública del archivo subido
    url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
    return url, nil
}


func GetPresignedURL(fileName string, expiration time.Duration) (string, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
    if err != nil {
        return "", fmt.Errorf("error al cargar configuración AWS: %w", err)
    }

    client := s3.NewFromConfig(cfg)
    presigner := s3.NewPresignClient(client)

    // Crear la clave completa del archivo
    key := fmt.Sprintf("uploads/%s", fileName)

    presignedReq, err := presigner.PresignGetObject(context.TODO(), &s3.GetObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(key),
    }, s3.WithPresignExpires(expiration))

    if err != nil {
        return "", fmt.Errorf("error al generar URL pre-firmada: %w", err)
    }

    return presignedReq.URL, nil
}

// Función helper con expiración por defecto
func GetPresignedURLDefault(fileName string) (string, error) {
    return GetPresignedURL(fileName, 10*time.Minute)
}
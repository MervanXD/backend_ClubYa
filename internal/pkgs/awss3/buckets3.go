package awss3

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	bucketName = "clubya-archivos"
	region     = "us-east-1"
)

func getAWSConfig() (aws.Config, error) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	sessionToken := os.Getenv("AWS_SESSION_TOKEN")
	awsRegion := os.Getenv("AWS_REGION")

	if awsRegion == "" {
		awsRegion = region // usar valor por defecto
	}

	if accessKey == "" || secretKey == "" {
		return aws.Config{}, fmt.Errorf("AWS_ACCESS_KEY_ID y AWS_SECRET_ACCESS_KEY son requeridos")
	}

	// Para cuentas de laboratorio con session token
	if sessionToken != "" {
		logs.Logger.Println("🔑 Configurando AWS con session token (cuenta de laboratorio)")
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(awsRegion),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(accessKey, secretKey, sessionToken),
			),
		)
		return cfg, err
	}

	// Para credenciales permanentes (sin session token)
	logs.Logger.Println("🔑 Configurando AWS sin session token (credenciales permanentes)")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)
	return cfg, err
}

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
    logs.Logger.Printf("🔄 Iniciando UploadFromBytes: %s (%d bytes)", fileName, len(fileData))
    
    // Verificar que tenemos datos
    if len(fileData) == 0 {
        return "", fmt.Errorf("archivo vacío")
    }
    if fileName == "" {
        return "", fmt.Errorf("nombre de archivo vacío")
    }
    
    // Cargar configuración AWS
    cfg, err := getAWSConfig()
    if err != nil {
        logs.Logger.Printf("❌ Error configuración AWS: %v", err)
        return "", fmt.Errorf("error al cargar configuración AWS: %w", err)
    }

    client := s3.NewFromConfig(cfg)
    key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), fileName)

    if contentType == "" {
        contentType = DetectContentType(fileData, fileName)
    }

    logs.Logger.Printf("🔄 Subiendo a S3: bucket=%s, key=%s, contentType=%s", bucketName, key, contentType)

    _, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket:      aws.String(bucketName),
        Key:         aws.String(key),
        Body:        bytes.NewReader(fileData),
        ContentType: aws.String(contentType),
    })

    if err != nil {
        logs.Logger.Printf("❌ Error PutObject S3: %v", err)
        return "", fmt.Errorf("error al subir archivo a S3: %w", err)
    }

    url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
    logs.Logger.Printf("✅ Upload exitoso: %s", url)
    return url, nil
}

func GetPresignedURL(fileName string, expiration time.Duration) (string, error) {
    cfg, err := getAWSConfig()
    if err != nil {
        return "", fmt.Errorf("error al cargar configuración AWS: %w", err)
    }

    client := s3.NewFromConfig(cfg)
    presigner := s3.NewPresignClient(client)

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

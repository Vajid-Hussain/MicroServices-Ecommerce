package helper_product_svc

import (
	"bytes"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	config_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/config"
)

func CreateSession(cfg *config_product_svc.Config) *session.Session {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(cfg.Region),
			Credentials: credentials.NewStaticCredentials(
				cfg.AccessKeyID,
				cfg.AccessKeySecret,
				"",
			),
		},
	))
	return sess
}

func CreateS3Session(sess *session.Session) *s3.S3 {
	s3Session := s3.New(sess)
	return s3Session
}

func UploadImageToS3(file []byte, sess *session.Session) (string, error) {

	imageReader := bytes.NewReader(file)

	fileName := uuid.New().String()

	uploader := s3manager.NewUploader(sess)

	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("mobile-mart"),                // Bucket name
		Key:    aws.String("product images/" + fileName), // S3 object key
		Body:   imageReader,                              // Reader containing the image data
		ACL:    aws.String("public-read"),                // ACL for the uploaded object
	})
	if err != nil {
		return "", err
	}
	return upload.Location, nil
}

func UploadFilesToS3(file bytes.Buffer, sess *session.Session) (string, error) {

	fmt.Println("##", bytes.Buffer(file))
	fileName := uuid.New().String()

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("mobile-mart"),
		Key:    aws.String("files/" + fileName),
		Body:   bytes.NewReader(file.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return upload.Location, nil
}

func UploadExcelToS3(file *excelize.File, sess *session.Session) (string, error) {

	var excelBuffer bytes.Buffer

	err := file.Write(&excelBuffer)
	if err != nil {
		return "", err
	}

	fileName := uuid.New().String()

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("mobile-mart"),
		Key:    aws.String("files/" + fileName),
		Body:   bytes.NewReader(excelBuffer.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return upload.Location, nil
}

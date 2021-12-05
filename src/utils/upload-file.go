package utils

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(file io.Reader, contentType string) *s3manager.UploadOutput {
	session := ConnectAWS()
	uploader := s3manager.NewUploader(session)
	bucketName := os.Getenv("AWS_S3_BUCKET_NAME")
	fileExt := GetContentTypeExt(contentType)

	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		ACL:         aws.String("public-read"),
		Key:         aws.String("APP_UPLOAD#" + CreateUlid() + fileExt),
		Body:        file,
		ContentType: &contentType,
	})

	if err != nil {
		panic(err)
	}

	return upload
}

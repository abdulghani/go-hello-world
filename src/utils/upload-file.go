package utils

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gabriel-vasile/mimetype"
)

const UPLOAD_PREFIX string = "APP_UPLOAD#"

func UploadFile(file io.Reader) *s3manager.UploadOutput {
	session := ConnectAWS()
	uploader := s3manager.NewUploader(session)
	bucketName := os.Getenv("AWS_S3_BUCKET_NAME")
	mtype, err := mimetype.DetectReader(file)

	if err != nil {
		panic(err)
	}

	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		ACL:         aws.String("public-read"),
		Key:         aws.String(UPLOAD_PREFIX + CreateUlid() + mtype.Extension()),
		Body:        file,
		ContentType: aws.String(mtype.String()),
	})

	if err != nil {
		panic(err)
	}

	return upload
}

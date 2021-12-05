package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess_singleton *session.Session

func ConnectAWS() *session.Session {
	if sess_singleton != nil {
		return sess_singleton
	}

	awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region:      &awsRegion,
		Credentials: credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, ""),
	})

	if err != nil {
		panic(err)
	}

	sess_singleton = sess

	return sess
}

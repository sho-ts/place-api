package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sho-ts/place-api/util"
	"mime/multipart"
	"os"
)

func UploadToS3Bucket(file multipart.File, name string) (string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_S3_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_S3_ACCESS_KEY"), os.Getenv("AWS_S3_SECRET_KEY"), ""),
	}))

	uploader := s3manager.NewUploader(sess)

	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET_NAME")),
		Body:   aws.ReadSeekCloser(file),
		Key:    aws.String(util.GetULID() + name),
		ACL:    aws.String("public-read"),
	})

	return output.Location, err
}

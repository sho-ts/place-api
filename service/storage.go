package service

import (
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/entity"
	"github.com/sho-ts/place-api/util"
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
	})

	return output.Location, err
}

func CreateStorage(postId string, authId string, path string) (entity.Storage, error) {
	storage := entity.Storage{
		Id:     util.GetULID(),
		UserId: authId,
		PostId: postId,
		Url:    path,
	}

	result := database.DB.Create(&storage)

	return storage, result.Error
}

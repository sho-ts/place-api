package repository

import (
	"mime/multipart"
)

type IStorageRepository interface {
	UploadToS3Bucket(file multipart.File, name string) (string, error)
}

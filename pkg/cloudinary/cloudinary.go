package cloudinary

import (
	"context"
	"errors"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var (
	errFailedInitializeCloudinary = errors.New("FAILED_INITIALIZE_CLOUDINARY")
)

type CloudinaryItf interface {
	UploadFile(ctx context.Context, file multipart.File) (string, error)
}

type Cloudinary struct {
	cloudinary *cloudinary.Cloudinary
}

func NewCloudinary() (CloudinaryItf, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, errFailedInitializeCloudinary
	}
	return &Cloudinary{
		cloudinary: cld,
	}, nil
}

func (c *Cloudinary) UploadFile(ctx context.Context, file multipart.File) (string, error) {
	unique := true
	folder := os.Getenv("CLOUDINARY_FOLDER")
	res, err := c.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:         folder,
		UniqueFilename: &unique,
	})

	if err != nil {
		return "", err
	}

	return res.SecureURL, nil
}

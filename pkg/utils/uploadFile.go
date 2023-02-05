package utils

import (
	"os"

	"BiteDans.com/tiktok-backend/pkg/constants"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(filename string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(constants.REGION)},
	)
	uploader := s3manager.NewUploader(sess)
	if err != nil {
		return "", err
	}

	file, _ := os.Open("./files/" + filename)
	defer file.Close()

	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(constants.BUCKET_NAME),
		Key: aws.String(filename),
		Body: file,
	})
	if err != nil {
		return "", err
	}
	
	return output.Location, nil
}
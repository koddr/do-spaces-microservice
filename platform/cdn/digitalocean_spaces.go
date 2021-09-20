package cdn

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/koddr/do-spaces-microservice/pkg/utils"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// DOSpacesConnection func for connect to DO Spaces CDN.
func DOSpacesConnection() (*minio.Client, error) {
	// Define DO Spaces variables for connection.
	endpoint := os.Getenv("DO_SPACES_ENDPOINT")
	accessKeyID := os.Getenv("DO_SPACES_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("DO_SPACES_SECRET_ACCESS_KEY")
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	// Set public policy to user_uploads folder.
	if errSetBucketPolicy := minioClient.SetBucketPolicy(
		context.Background(),
		os.Getenv("DO_SPACES_BUCKET_NAME"),
		fmt.Sprintf(
			`{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Action": ["s3:GetObject"],
						"Effect": "Allow",
						"Principal": {
							"AWS": ["*"]
						},
						"Resource": ["arn:aws:s3:::%v/%v/*"],
						"Sid": ""
					}
				]
			}`,
			os.Getenv("DO_SPACES_BUCKET_NAME"),
			os.Getenv("DO_SPACES_UPLOADS_FOLDER_NAME"),
		),
	); errSetBucketPolicy != nil {
		return nil, errSetBucketPolicy
	}

	return minioClient, err
}

//
func UploadObjectToCDN(minioClient *minio.Client, pathToFile, fileType, userID string) (minio.UploadInfo, error) {
	// Open the file from system path.
	file, errOpen := os.Open(filepath.Clean(pathToFile))
	if errOpen != nil {
		// Return empty info and error message.
		return minio.UploadInfo{}, errOpen
	}
	defer file.Close() // auto close file

	// Generate a new file name.
	newFileName, errGenerateFileName := utils.GenerateFileName(userID)
	if errGenerateFileName != nil {
		// Return empty info and error message.
		return minio.UploadInfo{}, errGenerateFileName
	}

	// Get file info.
	fileInfo, errGetFileInfo := utils.GetLocalFileInfo(filepath.Clean(pathToFile), fileType)
	if errGetFileInfo != nil {
		// Return empty info and error message.
		return minio.UploadInfo{}, errGetFileInfo
	}

	// Start uploading file to upload folder on CDN.
	// Folder: user ID, File name: SHA256 hash with origin extension.
	// Return info of the successfully uploaded file.
	return minioClient.PutObject(
		context.Background(),
		os.Getenv("DO_SPACES_BUCKET_NAME"),
		fmt.Sprintf(
			"%v/%v/%v.%v",
			os.Getenv("DO_SPACES_UPLOADS_FOLDER_NAME"), // upload folder name
			userID,             // user's folder name
			newFileName,        // file name
			fileInfo.Extension, // file extension
		),
		file,
		fileInfo.Size,
		minio.PutObjectOptions{
			ContentType: fileInfo.ContentType,
		},
	)
}

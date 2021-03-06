package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/koddr/do-spaces-microservice/app/models"

	"github.com/h2non/filetype"
)

// GetLocalFileInfo func for get local file's info (Content-Type, extension, size, etc).
func GetLocalFileInfo(pathToFile, fileType string) (*models.LocalFileInfo, error) {
	// Get size of local file.
	fileSize, err := GetFileSize(pathToFile)
	if err != nil {
		// Throw error with message.
		return nil, err
	}

	// Checking environment variable for max file size.
	maxFileSize, err := strconv.ParseInt(os.Getenv("MAX_UPLOAD_FILE_SIZE"), 10, 64)
	if err != nil {
		// Throw error with message.
		return nil, err
	}

	// Checking, if max file size is not reached.
	if fileSize > maxFileSize {
		// Throw error with message.
		return nil, fmt.Errorf("file is too large for upload")
	}

	// Read local file and buffering.
	buf, err := ioutil.ReadFile(filepath.Clean(pathToFile))
	if err != nil {
		// Throw error with message.
		return nil, err
	}

	// Set error messages.
	errMessageForFileType := fmt.Sprintf("only %s files are supported", fileType)
	errMessageForDefault := fmt.Sprintf("wrong or unsupported file type to upload (%s)", fileType)

	// Checking content type of the local file.
	switch fileType {
	case "archive":
		if !filetype.IsArchive(buf) {
			// Throw error with message.
			return nil, fmt.Errorf(errMessageForFileType)
		}
	case "audio":
		if !filetype.IsAudio(buf) {
			// Throw error with message.
			return nil, fmt.Errorf(errMessageForFileType)
		}
	case "document":
		if !filetype.IsDocument(buf) {
			// Throw error with message.
			return nil, fmt.Errorf(errMessageForFileType)
		}
	case "image":
		if !filetype.IsImage(buf) {
			// Throw error with message.
			return nil, fmt.Errorf(errMessageForFileType)
		}
	case "video":
		if !filetype.IsVideo(buf) {
			// Throw error with message.
			return nil, fmt.Errorf(errMessageForFileType)
		}
	default:
		// Throw error with message.
		return nil, fmt.Errorf(errMessageForDefault)
	}

	// Matching file type.
	kind, err := filetype.Match(buf)
	if err != nil {
		// Throw error with message.
		return nil, err
	}

	// Checking file type for unknown type.
	if kind == filetype.Unknown {
		// Throw error with message.
		return nil, fmt.Errorf("unknown file type")
	}

	// Return file info.
	return &models.LocalFileInfo{
		ContentType: kind.MIME.Value,
		Extension:   kind.Extension,
		Size:        fileSize,
	}, nil
}

// GenerateFileName func for generate name for uploaded file.
func GenerateFileName(userID string) (string, error) {
	// Create a new SHA256 hash.
	sha256 := sha256.New()

	// Create a new string with user ID and time string.
	name := userID + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := sha256.Write([]byte(name))
	if err != nil {
		// Return error, if generation failed.
		return "", err
	}

	// Return a new file name.
	return hex.EncodeToString(sha256.Sum(nil)), nil
}

// GetFileSize func for getting the file size.
func GetFileSize(pathToFile string) (int64, error) {
	// Get file from system path.
	file, err := os.Open(filepath.Clean(pathToFile))
	if err != nil {
		return 0, err
	}

	// Get file statistic.
	fileStat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	// Check, if file size is zero.
	if fileStat.Size() == 0 {
		// Return error message.
		return 0, fmt.Errorf("file have no size")
	}

	// Return file size in bytes.
	return fileStat.Size(), nil
}

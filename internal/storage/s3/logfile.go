package s3

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/agniBit/cryptonian/internal/logger"
)

const (
	s3Bucket = "cryptonian"
	s3Prefix = "logs/"
)

var uploadedFiles = map[string]bool{}

func SyncLogsToS3(ctx context.Context, logFileDir string) error {
	// Upload rotated log files
	files, err := os.ReadDir(logFileDir)
	if err != nil {
		return fmt.Errorf("unable to read logs directory: %v", err)
	}

	for _, file := range files {
		if _, ok := uploadedFiles[file.Name()]; !ok && file.Name() != "app.log" && strings.Contains(file.Name(), ".log.gz") {
			err := UploadToS3(ctx, s3Bucket, s3Prefix+file.Name(), logFileDir+file.Name())
			if err != nil {
				logger.Error(ctx, "Error uploading rotated log file to S3", err, map[string]interface{}{"file": file.Name()})
			}
		}
	}

	return nil
}

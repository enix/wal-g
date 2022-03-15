package storagetools

import (
	"os"

	"github.com/enix/wal-g/pkg/storages/storage"
	"github.com/wal-g/tracelog"
)

func HandleCatObject(objectPath string, folder storage.Folder, decrypt, decompress bool) {
	dstFile := os.Stdout
	err := downloadObject(objectPath, folder, dstFile, decrypt, decompress)
	tracelog.ErrorLogger.FatalfOnError("Failed to download the file: %v", err)
}

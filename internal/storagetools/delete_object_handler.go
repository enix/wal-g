package storagetools

import (
	"github.com/enix/wal-g/pkg/storages/storage"
	"github.com/wal-g/tracelog"
)

func HandleDeleteObject(objectPath string, folder storage.Folder) {
	// some storages may not produce an error on deleting the non-existing object
	exists, err := folder.Exists(objectPath)
	tracelog.ErrorLogger.FatalfOnError("Failed to check object existence: %v", err)
	if !exists {
		tracelog.ErrorLogger.Fatalf("Object %s does not exist", objectPath)
	}
	err = folder.DeleteObjects([]string{objectPath})
	tracelog.ErrorLogger.FatalfOnError("Failed to delete the specified object: %v", err)
}

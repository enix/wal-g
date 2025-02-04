package internal

import (
	"github.com/enix/wal-g/utility"
)

func HandleBackupMark(uploader *Uploader, backupName string, toPermanent bool, metaInteractor GenericMetaInteractor) {
	folder := uploader.UploadingFolder
	baseBackupFolder := uploader.UploadingFolder.GetSubFolder(utility.BaseBackupPath)
	uploader.UploadingFolder = baseBackupFolder

	markHandler := NewBackupMarkHandler(metaInteractor, folder)
	markHandler.MarkBackup(backupName, toPermanent)
}

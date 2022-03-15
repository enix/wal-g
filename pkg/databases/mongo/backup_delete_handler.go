package mongo

import (
	"github.com/enix/wal-g/pkg/databases/mongo/archive"
	"github.com/enix/wal-g/pkg/databases/mongo/models"
	"github.com/wal-g/tracelog"
)

// HandleBackupDelete deletes backup.
func HandleBackupDelete(backupName string, downloader archive.Downloader, purger archive.Purger, dryRun bool) error {
	backup, err := downloader.BackupMeta(backupName)
	if err != nil {
		return err
	}

	if dryRun {
		tracelog.InfoLogger.Printf("Skipping backup deletion due to dry-run: %+v", backup)
		return nil
	}

	if err := purger.DeleteBackups([]models.Backup{backup}); err != nil {
		return err
	}
	tracelog.InfoLogger.Printf("Backup was deleted: %+v", backup)
	return nil
}

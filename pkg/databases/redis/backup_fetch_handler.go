package redis

import (
	"context"
	"os/exec"

	"github.com/enix/wal-g/internal"
	"github.com/enix/wal-g/pkg/storages/storage"
	"github.com/enix/wal-g/utility"
)

func HandleBackupFetch(ctx context.Context, folder storage.Folder, backupName string, restoreCmd *exec.Cmd) error {
	backup, err := internal.GetBackupByName(backupName, utility.BaseBackupPath, folder)
	if err != nil {
		return err
	}
	return internal.StreamBackupToCommandStdin(restoreCmd, backup)
}

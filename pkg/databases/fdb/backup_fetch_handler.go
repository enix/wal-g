package fdb

import (
	"context"
	"os/exec"

	"github.com/enix/wal-g/internal"
	"github.com/enix/wal-g/pkg/storages/storage"
)

func HandleBackupFetch(ctx context.Context,
	folder storage.Folder,
	targetBackupSelector internal.BackupSelector,
	restoreCmd *exec.Cmd) {
	internal.HandleBackupFetch(folder, targetBackupSelector, internal.GetBackupToCommandFetcher(restoreCmd))
}

package mongo

import (
	"fmt"
	"os/exec"

	"github.com/enix/wal-g/internal"
	"github.com/enix/wal-g/pkg/databases/mongo/archive"
	"github.com/enix/wal-g/utility"
)

// HandleBackupPush starts backup procedure.
func HandleBackupPush(uploader archive.Uploader,
	metaConstructor internal.MetaConstructor,
	backupCmd *exec.Cmd) error {
	if err := metaConstructor.Init(); err != nil {
		return fmt.Errorf("can not initiate meta provider: %+v", err)
	}

	stdout, err := utility.StartCommandWithStdoutPipe(backupCmd)
	if err != nil {
		return fmt.Errorf("can not start backup command: %+v", err)
	}

	return uploader.UploadBackup(stdout, backupCmd, metaConstructor)
}

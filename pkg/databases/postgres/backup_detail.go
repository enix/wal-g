package postgres

import (
	"github.com/enix/wal-g/internal"
)

// BackupDetails is used to append ExtendedMetadataDto details to BackupTime struct
type BackupDetail struct {
	internal.BackupTime
	ExtendedMetadataDto
}

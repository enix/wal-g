package greenplum_test

import (
	"testing"

	"github.com/enix/wal-g/pkg/databases/greenplum"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRecoveryConf(t *testing.T) {
	walgPath := "/usr/bin/wal-g"
	cfgPath := "/etc/wal-g/wal-g.yaml"
	recoveryTargetName := "some_backup"
	recCfgMaker := greenplum.NewRecoveryConfigMaker(walgPath, cfgPath, recoveryTargetName)
	contentId := -1

	expectedCfg := `restore_command = '/usr/bin/wal-g seg wal-fetch "%f" "%p" --content-id=-1 --config /etc/wal-g/wal-g.yaml'
recovery_target_name = 'some_backup'`
	actualCfg := recCfgMaker.Make(contentId)
	assert.Equal(t, expectedCfg, actualCfg, "Actual recovery.conf does not match the expected one")
}

package walparser_test

import (
	"testing"

	"github.com/enix/wal-g/internal/walparser"
	"github.com/enix/wal-g/testtools"
	"github.com/stretchr/testify/assert"
)

func TestExtractBlockLocations(t *testing.T) {
	record, _ := testtools.GetXLogRecordData()
	expectedLocations := []walparser.BlockLocation{record.Blocks[0].Header.BlockLocation}
	actualLocations := walparser.ExtractBlockLocations([]walparser.XLogRecord{record})
	assert.Equal(t, expectedLocations, actualLocations)
}

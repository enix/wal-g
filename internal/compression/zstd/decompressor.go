package zstd

import (
	"io"

	"github.com/DataDog/zstd"
	"github.com/enix/wal-g/internal/compression/computils"
)

type Decompressor struct{}

func (decompressor Decompressor) Decompress(src io.Reader) (io.ReadCloser, error) {
	return zstd.NewReader(computils.NewUntilEOFReader(src)), nil
}

func (decompressor Decompressor) FileExtension() string {
	return FileExtension
}

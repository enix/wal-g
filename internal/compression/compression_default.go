//go:build !windows
// +build !windows

package compression

import (
	"github.com/enix/wal-g/internal/compression/gzip"
	"github.com/enix/wal-g/internal/compression/lz4"
	"github.com/enix/wal-g/internal/compression/lzma"
	"github.com/enix/wal-g/internal/compression/zstd"
)

var CompressingAlgorithms = []string{lz4.AlgorithmName, lzma.AlgorithmName}

var Compressors = map[string]Compressor{
	lz4.AlgorithmName:  lz4.Compressor{},
	lzma.AlgorithmName: lzma.Compressor{},
}

var Decompressors = []Decompressor{
	lz4.Decompressor{},
	lzma.Decompressor{},
	zstd.Decompressor{},
	gzip.Decompressor{},
}

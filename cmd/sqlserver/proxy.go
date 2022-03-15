package sqlserver

import (
	"github.com/enix/wal-g/internal"
	"github.com/enix/wal-g/pkg/databases/sqlserver"
	"github.com/spf13/cobra"
	"github.com/wal-g/tracelog"
)

const proxyShortDescription = "Run local azure blob emulator"

// proxyCmd represents the streamFetch command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: proxyShortDescription,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		folder, err := internal.ConfigureFolder()
		tracelog.ErrorLogger.FatalOnError(err)
		sqlserver.RunProxy(folder)
	},
}

func init() {
	cmd.AddCommand(proxyCmd)
}

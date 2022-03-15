package mongo

import (
	"os"

	"github.com/enix/wal-g/pkg/databases/mongo"
	"github.com/enix/wal-g/pkg/databases/mongo/archive"
	"github.com/spf13/cobra"
	"github.com/wal-g/tracelog"
)

const backupListShortDescription = "Prints available backups"

var verbose bool

// backupListCmd represents the backupList command
var backupListCmd = &cobra.Command{
	Use:   "backup-list",
	Short: backupListShortDescription, // TODO : improve description
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		downloader, err := archive.NewStorageDownloader(archive.NewDefaultStorageSettings())
		tracelog.ErrorLogger.FatalOnError(err)
		listing := archive.NewDefaultTabbedBackupListing()
		err = mongo.HandleBackupsList(downloader, listing, os.Stdout, verbose)
		tracelog.ErrorLogger.FatalOnError(err)
	},
}

func init() {
	cmd.AddCommand(backupListCmd)
	backupListCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose mode")
}

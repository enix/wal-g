package st

import (
	"github.com/enix/wal-g/internal"
	"github.com/enix/wal-g/internal/storagetools"
	"github.com/spf13/cobra"
	"github.com/wal-g/tracelog"
)

const folderListShortDescription = "Prints objects in the provided storage folder"
const recursiveFlag = "recursive"
const recursiveShortHand = "r"

// folderListCmd represents the folderList command
var folderListCmd = &cobra.Command{
	Use:   "ls [relative folder path]",
	Short: folderListShortDescription,
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		folder, err := internal.ConfigureFolder()
		tracelog.ErrorLogger.FatalOnError(err)

		if len(args) > 0 {
			folder = folder.GetSubFolder(args[0])
		}

		storagetools.HandleFolderList(folder, recursive)
	},
}

var recursive bool

func init() {
	folderListCmd.Flags().BoolVarP(&recursive, recursiveFlag, recursiveShortHand, false, "List folder recursively")
	StorageToolsCmd.AddCommand(folderListCmd)
}

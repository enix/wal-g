package sqlserver

import (
	"github.com/enix/wal-g/pkg/databases/sqlserver"
	"github.com/spf13/cobra"
)

const databaseListShortDescription = "List datbases in the backup"

var databaseListCmd = &cobra.Command{
	Use:   "database-list",
	Short: databaseListShortDescription,
	Run: func(cmd *cobra.Command, args []string) {
		sqlserver.HandleDatabaseList(args[0])
	},
}

func init() {
	cmd.AddCommand(databaseListCmd)
}

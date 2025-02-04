package common

import (
	"github.com/enix/wal-g/cmd/common/st"
	"github.com/enix/wal-g/internal"
	"github.com/spf13/cobra"
)

const usageTemplate = `Usage:{{if .Runnable}}
{{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
{{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
{{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
{{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}` +
	// additional custom message : cli flags introduced by 'internal.AddConfigFlags()' are hidden by default
	`

To get the complete list of all global flags, run: 'wal-g flags'` +
	`{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

func Init(cmd *cobra.Command, dbName string) {
	internal.ConfigureSettings(dbName)
	cobra.OnInitialize(internal.InitConfig, internal.Configure)

	cmd.SetUsageTemplate(usageTemplate)
	cmd.InitDefaultVersionFlag()
	internal.AddConfigFlags(cmd)

	cmd.PersistentFlags().StringVar(&internal.CfgFile, "config", "", "config file (default is $HOME/.walg.json)")

	// Init help subcommand
	cmd.InitDefaultHelpCmd()
	helpCmd, _, _ := cmd.Find([]string{"help"})
	// fix to disable the required settings check for the help subcommand
	helpCmd.PersistentPreRun = func(*cobra.Command, []string) {}

	// Add flags subcommand
	cmd.AddCommand(FlagsCmd)

	// Add storage tools
	cmd.AddCommand(st.StorageToolsCmd)
}

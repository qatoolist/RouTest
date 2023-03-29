package internal

import (
	"fmt"
	"os"

	routest "github.com/qatoolist/RouTest"
	"github.com/spf13/cobra"
)

// CreateVersionCmd creates the version subcommand.
func CreateVersionCmd() cobra.Command {
	versionCmd := cobra.Command{
		Use:     "version",
		Short:   "Show current version",
		Run:     versionCmdRunFunc,
		Version: routest.Version,
	}

	return versionCmd
}

func versionCmdRunFunc(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "RouTest version is:", routest.Version)
}

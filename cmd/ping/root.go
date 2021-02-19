package ping

import (
	"github.com/ChanKunggc/ping-go/internal"
	"github.com/ChanKunggc/ping-go/pkg/logger"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping tool",
	Long:  "ping tool which used by  icmp/tcp",
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	if args == nil || len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	} else if len(args) > 1 {
		logger.Error("too much args ")
		os.Exit(1)
	}
	internal.Ping(tcpUse, args[0], "")
}

var tcpUse bool

func init() {
	flagSets := RootCmd.PersistentFlags()
	flagSets.BoolVarP(&tcpUse, "tcp", "t", false, "use tcp")
	flagSets.BoolVarP(&logger.Verbose, "verbose", "v", false, "more print ")
}

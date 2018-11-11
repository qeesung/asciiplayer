package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

// commonOptions define the global options that all sub command
// will inherit.
type commonOptions struct {
	Debug bool
}

// Cli is wrapper for root command and add common options
type Cli struct {
	commonOptions
	rootCmd *cobra.Command
}

// NewCli create a new cli object
func NewCli() *Cli {
	return &Cli{
		rootCmd: &cobra.Command{
			Use:   "asciiplayer",
			Short: "asciiplayer is a command line tool to play gif and video in ASCII mode",
			Long: SummaryTitle + `

asciiplayer is a library that can convert gif and video to ASCII image
and provide the cli for easy use.
`,
			DisableAutoGenTag: true,
		},
	}
}

// SetFlags for cli define the global flags
func (cli *Cli) SetFlags() *Cli {
	flags := cli.rootCmd.PersistentFlags()
	flags.BoolVarP(&cli.Debug, "debug", "D", false, "Switch log level to DEBUG mode")
	return cli
}

// InitLog init the log , and config the log
func (cli *Cli) InitLog() {
	if cli.Debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Infof("start client at debug level")
	}

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	}
	logrus.SetFormatter(formatter)
}

// AddCommand add a child command to parent command
func (cli *Cli) AddCommand(parent, child Command) {
	child.Init()

	parentCmd := parent.Cmd()
	childCmd := child.Cmd()

	childCmd.PreRun = func(cmd *cobra.Command, args []string) {
		cli.InitLog()
	}

	parentCmd.AddCommand(childCmd)
}

// Run start to run the cli command
func (cli *Cli) Run() error {
	return cli.rootCmd.Execute()
}

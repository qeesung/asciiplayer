package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

type commonOptions struct {
	Debug bool
}

type Cli struct {
	commonOptions
	rootCmd *cobra.Command
}

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

func (cli *Cli) SetFlags() *Cli {
	flags := cli.rootCmd.PersistentFlags()
	flags.BoolVarP(&cli.Debug, "debug", "D", false, "Switch log level to DEBUG mode")
	return cli
}

func (cli *Cli) InitLog() {
	if cli.Debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Infof("start client at debug level")
	}

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	}
	logrus.SetFormatter(formatter)
}

func (cli *Cli) AddCommand(parent, child Command) {
	child.Init()

	parentCmd := parent.Cmd()
	childCmd := child.Cmd()

	childCmd.PreRun = func(cmd *cobra.Command, args []string) {
		cli.InitLog()
	}

	parentCmd.AddCommand(childCmd)
}

func (cli *Cli) Run() error {
	return cli.rootCmd.Execute()
}

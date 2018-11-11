package cmd

// CommandBuilder responsible for building root command
// and register all sub commands to the root command
type CommandBuilder struct {
}

// Build create root command and register all sub command
// and return the cli object that contain the root command
func (builder *CommandBuilder) Build() *Cli {
	cli := NewCli()

	// set global flags for rootCmd in cli.
	cli.SetFlags()

	base := &baseCommand{cmd: cli.rootCmd}
	base.Cmd().SilenceErrors = true

	cli.AddCommand(base, &PlayCommand{})
	cli.AddCommand(base, &VersionCommand{})
	cli.AddCommand(base, &EncodeCommand{})
	cli.AddCommand(base, &ServerCommand{})
	return cli
}

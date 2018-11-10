package cmd

import (
	"errors"
	"fmt"
	"github.com/qeesung/asciiplayer/pkg/remote"
	"github.com/qeesung/asciiplayer/pkg/util"
	"github.com/qeesung/image2ascii/convert"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"net/http"
)

type ServerCommand struct {
	baseCommand
	convert.Options
	Delay float64
	Host  string
	Port  string
}

func (serverCommand *ServerCommand) Init() {
	serverCommand.cmd = &cobra.Command{
		Use:   "server",
		Short: "Server command setup a http share server",
		Args:  cobra.ExactArgs(1),
		Long: SummaryTitle + `

Setup a http server, and share your ascii image with others. Setup a http server, then access through curl command.

Setup server

$ asciiplayer server demo.gif
# Server available on : http://0.0.0.0:8080
Access from remote

$ curl http://hostname:8080
# play ascii image here
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return serverCommand.server(args)
		},
		Example: serverExample(),
	}
	serverCommand.addFlags()
}

func (serverCommand *ServerCommand) server(args []string) error {
	filename := args[0]
	flushHandler, supported := remote.NewFlushHandler(filename, &serverCommand.Options)
	if !supported {
		return errors.New("not supported file type")
	}

	err := flushHandler.Init()
	if err != nil {
		return err
	}

	http.HandleFunc("/", flushHandler.HandlerFunc())

	err = serverCommand.printServerAddress()
	if err != nil {
		return err
	}

	addr := serverCommand.Host + ":" + serverCommand.Port
	logrus.Debug("Going to listen and serve on " + addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}

	return nil
}

func (serverCommand *ServerCommand) addFlags() {
	flagSet := serverCommand.cmd.Flags()

	flagSet.Float64VarP(&serverCommand.Ratio, "ratio", "r", 1.0, "Scale ratio")
	flagSet.IntVarP(&serverCommand.FixedWidth, "width", "w", -1, "Scale to fixed width")
	flagSet.IntVarP(&serverCommand.FixedHeight, "height", "g", -1, "Scale to fixed height")
	flagSet.BoolVarP(&serverCommand.StretchedScreen, "stretched", "t", false, "Stretch the image to fit screen")
	flagSet.BoolVarP(&serverCommand.Colored, "colored", "c", true, "Play with color")
	flagSet.BoolVarP(&serverCommand.Reversed, "reversed", "i", false, "Play with the ascii reversed")
	flagSet.BoolVarP(&serverCommand.FitScreen, "fit", "s", true, "Play fit the screen")
	flagSet.Float64VarP(&serverCommand.Delay, "delay", "d", 0.15, "Play delay duration between two frames")
	flagSet.StringVarP(&serverCommand.Host, "host", "H", "0.0.0.0", "Server host address")
	flagSet.StringVarP(&serverCommand.Port, "port", "p", "8080", "Server host port")
}

func (serverCommand *ServerCommand) printServerAddress() (err error) {
	serverAddrStyle := chalk.Green.NewStyle().
		WithBackground(chalk.Black).
		WithTextStyle(chalk.Bold).Style

	fmt.Println(Title)
	fmt.Println("Server available on :")
	if serverCommand.Host == "0.0.0.0" || serverCommand.Host == "" {
		serverIPList, err := util.GetIPList()
		if err != nil {
			return err
		}
		for _, ipAddr := range serverIPList {
			fmt.Println(serverAddrStyle(fmt.Sprintf("http://%s:%s", ipAddr, serverCommand.Port)))
		}
	} else {
		fmt.Println(serverAddrStyle(fmt.Sprintf("http://%s:%s", serverCommand.Host, serverCommand.Port)))
	}
	fmt.Println(chalk.Blue.Color(fmt.Sprintf("Access from remote with command : `curl http://hostname:%s`",
		serverCommand.Port)))
	fmt.Println("Hit CTRL-C to stop the server")
	return
}

func serverExample() string {
	return `Setup a http server with default port and host
$ asciiplayer server demo.gif

Setup a http server with the custom port
$ asciiplayer server demo.gif --port 8888
`
}

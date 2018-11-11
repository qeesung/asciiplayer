package cmd

import (
	"errors"
	"fmt"
	"github.com/qeesung/asciiplayer/pkg/decoder"
	"github.com/qeesung/asciiplayer/pkg/encoder"
	"github.com/qeesung/asciiplayer/pkg/progress"
	"github.com/qeesung/asciiplayer/pkg/render"
	"github.com/qeesung/asciiplayer/pkg/util"
	"github.com/qeesung/image2ascii/convert"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

type EncodeCommand struct {
	baseCommand
	convert.Options
	Delay           float64
	FontSize        int
	BackGroundColor string
	ForeGroundColor string
	OutputFilename  string
}

func (encodeCommand *EncodeCommand) Init() {
	encodeCommand.cmd = &cobra.Command{
		Use:   "encode",
		Short: "Encode gif or video to ascii gif or video",
		Args:  cobra.ExactArgs(1),
		Long: SummaryTitle + `

Encode command can convert gif or video to a ascii gif or video.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return encodeCommand.encode(args)
		},
		Example: encodeExample(),
	}
	encodeCommand.addFlags()
}

func (encodeCommand *EncodeCommand) encode(args []string) error {
	waitingBar := progress.WaitingBar{}
	waitingBar.Start()

	convertOptions, drawOptions, err := encodeCommand.parseFlags()
	if err != nil {
		return err
	}

	inputFilename := args[0]
	outputFilename := encodeCommand.OutputFilename

	logrus.Debugf("Start encoding %s to %s...", inputFilename, outputFilename)
	inputDecoder, supported := decoder.NewDecoder(inputFilename)
	if !supported {
		return errors.New("not supported input file type")
	}

	outputEncoder, supported := encoder.NewEncoder(outputFilename)
	if !supported {
		return errors.New("not supported output file type")
	}

	logrus.Debugf("Decoding the input file %s...", inputFilename)
	decodeNotifier := waitingBar.AddBar("Decoding ", progress.MaxSteps)
	frames, err := inputDecoder.DecodeFromFile(inputFilename, decodeNotifier)
	if err != nil {
		return err
	}

	drawer := render.NewImageDrawer()

	logrus.Debugf("Rendering the frames to ASCII frames...")
	convertNotifier := waitingBar.AddBar("Rendering", len(frames))
	asciiImages, err := drawer.BatchConvertThenDraw(frames, convertOptions, drawOptions, convertNotifier)
	if err != nil {
		return err
	}

	encodeNotifier := waitingBar.AddBar("Encoding ", len(asciiImages))
	logrus.Debugf("Encoding the frames to output file %s...", outputFilename)
	outputEncoder.EncodeToFile(outputFilename, asciiImages, encodeNotifier)
	waitingBar.Wait()
	fmt.Printf("File saved to %s\n", chalk.Green.Color(outputFilename))
	return nil
}

func (encodeCommand *EncodeCommand) addFlags() {
	flagSet := encodeCommand.cmd.Flags()

	flagSet.Float64VarP(&encodeCommand.Ratio, "ratio", "r", 1.0, "Scale ratio")
	flagSet.IntVarP(&encodeCommand.FixedWidth, "width", "w", -1, "Scale to fixed width")
	flagSet.IntVarP(&encodeCommand.FixedHeight, "height", "g", -1, "Scale to fixed height")
	flagSet.BoolVarP(&encodeCommand.Colored, "colored", "c", true, "Encode with color")
	flagSet.BoolVarP(&encodeCommand.Reversed, "reversed", "i", false, "Encode with the ascii reversed")
	flagSet.Float64VarP(&encodeCommand.Delay, "delay", "d", 0.15, "Encode delay duration between two frames")
	flagSet.BoolVarP(&encodeCommand.FitScreen, "fit", "s", true, "Encode fit the screen")
	flagSet.BoolVarP(&encodeCommand.StretchedScreen, "stretched", "t", false, "Encode stretch the image to fit screen")
	flagSet.StringVarP(&encodeCommand.OutputFilename, "out", "o", "", "Encode output filename")
	encodeCommand.cmd.MarkFlagRequired("out")
	flagSet.IntVarP(&encodeCommand.FontSize, "font_size", "z", 15, "Encode ASCII font size(pt)")
	flagSet.StringVarP(&encodeCommand.BackGroundColor, "bg", "j", "#000000", "Encode ASCII background color")
	flagSet.StringVarP(&encodeCommand.ForeGroundColor, "fg", "k", "#FFFFFF", "Encode ASCII foreground color")
}

func (encodeCommand *EncodeCommand) parseFlags() (convertOptions convert.Options,
	drawOptions render.DrawOptions, err error) {
	convertOptions = encodeCommand.Options

	drawOptions = render.DefaultDrawOptions
	drawOptions.FontSize = float64(encodeCommand.FontSize)
	drawOptions.Colored = encodeCommand.Colored
	if bgColor, err := util.ConvertHexToRGB(encodeCommand.BackGroundColor); err == nil {
		drawOptions.BackGroundColor = bgColor
	} else {
		return convert.Options{}, render.DrawOptions{}, err
	}

	if fgColor, err := util.ConvertHexToRGB(encodeCommand.BackGroundColor); err == nil {
		drawOptions.ForeGroundColor = fgColor
	} else {
		return convert.Options{}, render.DrawOptions{}, err
	}
	return
}

func encodeExample() string {
	return `Encode gif image to ascii gif image
$ asciiplayer encode demo.gif -o output.gif

Encode gif image to ascii gif image with custom font size
$ asciiplayer encode demo.gif -o output.gif --font_size=5

Zoom to the original 1/10, then encode gif image to ascii gif image
$ asciiplayer encode demo.gif -o output.gif -r 0.1
`
}

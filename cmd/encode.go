package cmd

import (
	"errors"
	"github.com/qeesung/asciiplayer/pkg/asciiimage"
	"github.com/qeesung/asciiplayer/pkg/decoder"
	"github.com/qeesung/asciiplayer/pkg/encoder"
	"github.com/qeesung/asciiplayer/pkg/util"
	"github.com/qeesung/image2ascii/convert"
	"github.com/spf13/cobra"
	"image"
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

Play command only work in terminal, decoding the gif or video
info multi frames and convert the frames to ASCII character matrix,
finally, output the matrix to stdout at a certain frequency.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return encodeCommand.encode(args)
		},
		Example: encodeExample(),
	}
	encodeCommand.addFlags()
}

func (encodeCommand *EncodeCommand) encode(args []string) error {
	convertOptions, drawOptions, err := encodeCommand.parseFlags()
	if err != nil {
		return err
	}

	inputFilename := args[0]

	inputDecoder, supported := decoder.NewDecoder(inputFilename)
	if !supported {
		return errors.New("not supported input file type")
	}

	frames, err := inputDecoder.DecodeFromFile(inputFilename)
	if err != nil {
		return err
	}

	imageConverter := convert.NewImageConverter()
	drawer := asciiimage.NewImageDrawer()

	asciiImages := make([]image.Image, 0, len(frames))
	for _, frame := range frames {
		charPixelMatrix := imageConverter.Image2CharPixelMatrix(frame, &convertOptions)
		asciiImage, err := drawer.DrawCharPixelMatrix2Image(charPixelMatrix, drawOptions)
		if err != nil {
			return err
		}
		asciiImages = append(asciiImages, asciiImage)
	}

	outputFilename := encodeCommand.OutputFilename
	outputEncoder, supported := encoder.NewEncoder(outputFilename)
	if !supported {
		return errors.New("not supported output file type")
	}
	outputEncoder.EncodeToFile(outputFilename, asciiImages)
	return nil
}

func (encodeCommand *EncodeCommand) addFlags() {
	flagSet := encodeCommand.cmd.Flags()

	flagSet.Float64VarP(&encodeCommand.Ratio, "ratio", "r", 1.0, "Scale ratio")
	flagSet.IntVarP(&encodeCommand.FixedWidth, "width", "w", -1, "Scale to fixed width")
	flagSet.IntVarP(&encodeCommand.FixedHeight, "height", "g", -1, "Scale to fixed height")
	flagSet.BoolVarP(&encodeCommand.Colored, "colored", "c", true, "Play with color")
	flagSet.BoolVarP(&encodeCommand.Reversed, "reversed", "i", false, "Play with the ascii reversed")
	flagSet.Float64VarP(&encodeCommand.Delay, "delay", "d", 0.15, "Play delay duration between two frames")
	flagSet.StringVarP(&encodeCommand.OutputFilename, "out", "o", "", "Encode output filename")
	encodeCommand.cmd.MarkFlagRequired("out")
	flagSet.IntVarP(&encodeCommand.FontSize, "font_size", "z", 20, "Encode ASCII font size(pt)")
	flagSet.StringVarP(&encodeCommand.BackGroundColor, "bg", "j", "#000000", "Encode ASCII background color")
	flagSet.StringVarP(&encodeCommand.ForeGroundColor, "fg", "k", "#FFFFFF", "Encode ASCII foreground color")
}

func (encodeCommand *EncodeCommand) parseFlags() (convertOptions convert.Options,
	drawOptions asciiimage.DrawOptions, err error) {
	convertOptions = encodeCommand.Options
	convertOptions.FitScreen = false
	convertOptions.StretchedScreen = false

	drawOptions = asciiimage.DefaultDrawOptions
	drawOptions.FontSize = float64(encodeCommand.FontSize)
	drawOptions.Colored = encodeCommand.Colored
	if bgColor, err := util.ConvertHexToRGB(encodeCommand.BackGroundColor); err == nil {
		drawOptions.BackGroundColor = bgColor
	} else {
		return convert.Options{}, asciiimage.DrawOptions{}, err
	}

	if fgColor, err := util.ConvertHexToRGB(encodeCommand.BackGroundColor); err == nil {
		drawOptions.ForeGroundColor = fgColor
	} else {
		return convert.Options{}, asciiimage.DrawOptions{}, err
	}
	return
}

func encodeExample() string {
	return `$ asciiplay encode input.gif -o output.gif`
}

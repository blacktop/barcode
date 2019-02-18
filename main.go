package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"sort"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"github.com/urfave/cli"
)

var (
	// Version stores the plugin's version
	Version string
	// BuildTime stores the plugin's build time
	BuildTime string
)

func main() {
	var (
		height int
		width  int
		size   int
	)
	app := cli.NewApp()
	app.Name = "barcode"
	app.Author = "blacktop"
	app.Email = "https://github.com/blacktop"
	app.Version = Version + ", BuildTime: " + BuildTime
	app.Compiled, _ = time.Parse("20060102", BuildTime)
	app.Usage = "Create Barcodes or QR codes"
	app.Commands = []cli.Command{
		{
			Name:  "bar",
			Usage: "create barcode",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "height",
					Usage:       "Height of barcode",
					Value:       200,
					Destination: &height,
				},
				cli.IntFlag{
					Name:        "width",
					Usage:       "Width of barcode",
					Value:       600,
					Destination: &width,
				},
			},
			Action: func(c *cli.Context) error {
				if c.Args().Present() {
					// Create the barcode
					bar, _ := code128.EncodeWithoutChecksum(c.Args().First())
					// Scale the barcode to 200x200 pixels
					image, _ := barcode.Scale(bar, width, height)
					// create the output file
					file, _ := os.Create("barcode.png")
					defer file.Close()

					// encode the barcode as png
					png.Encode(file, image)
				} else {
					log.Fatal(fmt.Errorf("Please supply string to encode into barcode"))
				}
				return nil
			},
		},
		{
			Name:  "qr",
			Usage: "create qr code",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "size",
					Usage:       "Size of qrcode",
					Value:       200,
					Destination: &size,
				},
			},
			Action: func(c *cli.Context) error {
				if c.Args().Present() {
					// Create the barcode
					qrCode, _ := qr.Encode(c.Args().First(), qr.M, qr.Auto)
					// Scale the barcode to 200x200 pixels
					qrCode, _ = barcode.Scale(qrCode, size, size)
					// create the output file
					file, _ := os.Create("qrcode.png")
					defer file.Close()

					// encode the barcode as png
					png.Encode(file, qrCode)
				} else {
					log.Fatal(fmt.Errorf("Please supply string to encode into QR code"))
				}
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

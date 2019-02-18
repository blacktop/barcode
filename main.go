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

	app := cli.NewApp()
	app.Name = "barcode"
	app.Author = "blacktop"
	app.Email = "https://github.com/blacktop"
	app.Version = Version + ", BuildTime: " + BuildTime
	app.Compiled, _ = time.Parse("20060102", BuildTime)
	app.Usage = "Create Barcodes or QR codes"
	// app.Flags = []cli.Flag{
	// 	cli.StringFlag{
	// 		Name:  "lang, l",
	// 		Value: "english",
	// 		Usage: "Language for the greeting",
	// 	},
	// 	cli.StringFlag{
	// 		Name:  "config, c",
	// 		Usage: "Load configuration from `FILE`",
	// 	},
	// }
	app.Commands = []cli.Command{
		{
			Name:    "bar",
			Aliases: []string{"b"},
			Usage:   "create barcode",
			Action: func(c *cli.Context) error {
				if c.Args().Present() {
					// Create the barcode
					bar, _ := code128.EncodeWithoutChecksum(c.Args().First())
					// Scale the barcode to 200x200 pixels
					image, _ := barcode.Scale(bar, 600, 200)
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
			Name:    "qr",
			Aliases: []string{"q"},
			Usage:   "create qr code",
			Action: func(c *cli.Context) error {
				if c.Args().Present() {
					// Create the barcode
					qrCode, _ := qr.Encode(c.Args().First(), qr.M, qr.Auto)
					// Scale the barcode to 200x200 pixels
					qrCode, _ = barcode.Scale(qrCode, 200, 200)
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

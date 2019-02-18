<p align="center">
  <a href="https://github.com/blacktop/barcode"><img alt="Malice Logo" src="https://github.com/blacktop/barcode/raw/master/docs/qrcode.png" height="140" /></a>
  <a href="https://github.com/blacktop/barcode"><h3 align="center">barcode</h3></a>
  <p align="center">Create barcodes (QR/Code128)</p>
</p>

[![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org) [![Github All Releases](https://img.shields.io/github/downloads/blacktop/barcode/total.svg)](https://github.com/blacktop/barcode/releases/latest) [![GitHub release](https://img.shields.io/github/release/blacktop/barcode.svg)](https://github.com/blacktop/barcode/releases)

---

## Getting Started

### Install

Download `barcode` from [releases](https://github.com/blacktop/barcode/releases)

### Usage

```
NAME:
   barcode - Create Barcodes or QR codes

USAGE:
   barcode [global options] command [command options] [arguments...]

VERSION:
   v0.1.0, BuildTime: 2019-02-18T18:59:22Z

AUTHOR:
   blacktop <https://github.com/blacktop>

COMMANDS:
     bar      create barcode
     qr       create qr code
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

#### Create QR Code

```
NAME:
   barcode qr - create qr code

USAGE:
   barcode qr [command options] [arguments...]

OPTIONS:
   --size value  Size of qrcode (default: 200)
   --output value, -o value  output png file (default: "qrcode.png")
```

```bash
$ barcode qr --size 200 --output secret.png "secret message"
```

#### Create Barcode

```
NAME:
   barcode bar - create barcode

USAGE:
   barcode bar [command options] [arguments...]

OPTIONS:
   --height value            Height of barcode (default: 200)
   --width value             Width of barcode (default: 600)
   --output value, -o value  output png file (default: "barcode.png")
```

```bash
$ barcode bar --height 200 --width 600 --output secret.png "secret message"
```

## License

MIT Copyright (c) 2019 blacktop

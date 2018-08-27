package main

import (
	"fmt"
	"github.com/andy-zhangtao/systemd-service-make/gpl"
	"github.com/andy-zhangtao/systemd-service-make/module"
	"github.com/andy-zhangtao/systemd-service-make/tools"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
	"strings"
)

var name string
var desc string
var after string
var requires string
var image string
var raw string

func main() {
	app := cli.NewApp()
	app.Name = "Service Maker"
	app.Usage = "Generate Systemd Service"
	app.Version = "0.1.0"
	app.Author = "ztao8607@gmail.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name, n",
			Usage:       "Service Name. Require",
			Destination: &name,
		},
		cli.StringFlag{
			Name:        "desc, d",
			Usage:       "Service Desc",
			Destination: &desc,
		},
		cli.StringFlag{
			Name:        "After, af",
			Usage:       "The After Service List. Multiple service use ',' for split",
			Destination: &after,
		},
		cli.StringFlag{
			Name:        "Requires, r",
			Usage:       "The Require Services. Multiple service use ',' for split",
			Destination: &requires,
		},
		cli.StringSliceFlag{
			Name:  "Args,ar",
			Usage: "The Docker run args",
		},
		cli.StringFlag{
			Name:        "Image, i",
			Usage:       "The Image Name. Require",
			Destination: &image,
		},
		cli.StringFlag{
			Name:        "Raw, raw",
			Usage:       "The raw docker run command",
			Destination: &raw,
		},
	}

	app.Action = func(c *cli.Context) error {

		if !check() {
			fmt.Println("Name and Image are mandatory choose. PLease see smarker usage")
			cli.ShowAppHelpAndExit(c, -1)
		}
		service := module.SystemdServiceModule{
			Name:         name,
			Desc:         desc,
			AfterService: strings.Split(after, ","),
			Requires:     requires,
			Image:        image,
		}

		if raw != "" {
			service.Args = tools.ParseRawArgs(raw)
		}
		if len(c.StringSlice("ar")) > 0 {
			service.Args = c.StringSlice("ar")
		}
		return gpl.GenerateSystemdService(service)
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func check() bool {
	if name == "" {
		return false
	}

	if image == "" {
		return false
	}

	return true
}

package main

import (
	"github.com/bearbin/go-mcaccutils"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:      "name",
			ShortName: "n",
			Usage:     "get a username from a uuid",
			Action: func(c *cli.Context) {
				name, err := mcaccutils.GetName(c.Args().First())
				if err != nil {
					println("ERROR! " + err.Error())
				}
				println(name)
			},
		},
		{
			Name:      "uuid",
			ShortName: "u",
			Usage:     "get a uuid from a username",
			Action: func(c *cli.Context) {
				uuid, _, err := mcaccutils.GetUUID(c.Args().First())
				if err != nil {
					println("ERROR! " + err.Error())
				}
				println(uuid)
			},
		},
	}
	app.Run(os.Args)
}

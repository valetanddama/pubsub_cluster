package main

import (
	"github.com/satori/go.uuid"
	"github.com/valetanddama/cluster/pkg/config"
	"github.com/valetanddama/cluster/pkg/services"
	"gopkg.in/urfave/cli.v2"
	"os"
	"time"
)

func main() {
	app := cli.App{
		Name:  uuid.Must(uuid.NewV4(), nil).String(),
		Usage: "Publisher or subscriber",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "getErrors",
				Usage: "Print errors",
				Value: false,
			},
		},
		Action: func(c *cli.Context) error {
			defer config.Conn.Redis().Close()

			if c.Bool("getErrors") {
				services.GetListOfWrongMessages()
				return nil
			}

			config.Conn.ServerRole = config.RoleHandlerMessages
			config.Conn.ServerID = c.App.Name

			getMessage := time.NewTicker(500 * time.Millisecond).C
			sendMessage := time.NewTicker(500 * time.Millisecond).C
			checkServerRole := time.NewTicker(500 * time.Millisecond).C

			for {
				select {
				case <-getMessage:
					services.GetMessage()
				case <-sendMessage:
					services.SendMessage()
				case <-checkServerRole:
					services.CheckServerRole()
				}
			}

			return nil
		},
	}

	app.Run(os.Args)
}

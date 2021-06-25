package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sarrufat/natsk8s/webpub/service"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	version string
	build   string
)

func mainApp() *cli.App {
	return &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "BindAddress",
				EnvVars: []string{"BIND_ADDRESS"},
				Value:   ":9090",
			},
			&cli.StringFlag{
				Name:    "URL",
				EnvVars: []string{"NATS_URL"},
				Value:   "nats://my-nats",
			},
		},
		Action: mainAction,
	}
}

func mainAction(context *cli.Context) error {
	r := gin.Default()
	logger := log.New(os.Stdout, "webpub ", log.LstdFlags)
	mp := service.NewMessageProducer(logger, context.String("URL"))
	r.PUT("/nats", mp.DoAction)
	// startServer(server)
	r.Run(":9090")
	return nil
}

func main() {
	app := mainApp()
	app.Run(os.Args)

}

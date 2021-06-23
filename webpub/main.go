package main

import (
	"context"
	"github.com/sarrufat/natsk8s/webpub/service"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	logger := log.New(os.Stdout, "webpub ", log.LstdFlags)
	mp := service.NewMessageProducer(logger, context.String("URL"))
	smux := http.NewServeMux()
	smux.Handle("/", mp)
	server := http.Server{
		Addr:         context.String("BindAddress"),
		Handler:      smux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	startServer(server)
	return nil
}

func startServer(server http.Server) {
	go func() {
		server.ErrorLog.Println("Starting server on ", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			server.ErrorLog.Fatal("Error starting server ", err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)
	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}

func main() {
	app := mainApp()
	app.Run(os.Args)

}

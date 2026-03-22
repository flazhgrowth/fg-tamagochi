package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/flazhgrowth/fg-tamagochi/appconfig"
)

func (app *Server) Run(cfg appconfig.HTTPServerConfig) error {
	srv := http.Server{
		ReadTimeout:  cfg.Timeout.ReadTimeout.Val("second"),
		WriteTimeout: cfg.Timeout.WriteTimeout.Val("second"),
		IdleTimeout:  cfg.Timeout.IdleTimeout.Val("second"),
		Handler:      app.serverRouter.Routes(),
	}
	listener, err := getListener(cfg.Server)
	if err != nil {
		return err
	}

	go func() {
		fmt.Println("listening with pid " + fmt.Sprint(os.Getpid()))
		fmt.Printf("webserver runs on: %s\n", cfg.Server)
		srv.Serve(listener)
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("got signal", <-sigs)
	fmt.Println("shutting down..")
	srv.Shutdown(context.Background())
	return nil
}

func getListener(port string) (net.Listener, error) {
	if port == "" {
		port = "8988" // default ports
	}
	var listener net.Listener

	EINHORN_FDS := os.Getenv("EINHORN_FDS")
	if EINHORN_FDS != "" {
		fds, err := strconv.Atoi(EINHORN_FDS)
		if err != nil {
			return nil, err
		}
		log.Println("using socket master, listening on", EINHORN_FDS)
		f := os.NewFile(uintptr(fds), "listener")
		listener, err = net.FileListener(f)
		if err != nil {
			log.Fatalln("error create listener", err)
			return nil, err
		}
		return listener, nil
	}
	return net.Listen("tcp4", fmt.Sprintf(":%s", port))
}

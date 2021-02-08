package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	cm "tugas_6/Framework/git/order/common"
	"tugas_6/Framework/git/order/middleware"
	"tugas_6/Framework/git/order/services"
	"tugas_6/Framework/git/order/transport"

	log "github.com/Sirupsen/logrus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func initHandlers() {

	var svc services.PaymentServices

	svc = services.PaymentService{}
	svc = middleware.BasicMiddleware()(svc)

	root := cm.Config.RootURL

	http.Handle(fmt.Sprintf("%s/trip", root), httptransport.NewServer(
		transport.TripEndpoint(svc), transport.DecodeTripRequest, transport.EncodeResponse,
	))

}

var logger *log.Entry

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

}

func main() {

	configFile := flag.String("conf", "conf-dev.yml", "main configuration file")
	flag.Parse()
	initLogger()
	log.WithField("file", *configFile).Info("Loading configuration file")
	cm.LoadConfigFromFile(configFile)
	initHandlers()

	var err error
	if cm.Config.RootURL != "" || cm.Config.ListenPort != "" {
		err = http.ListenAndServe(cm.Config.ListenPort, nil)
	}

	if err != nil {
		log.WithField("error", err).Error("Unable to start the server")
		os.Exit(1)
	}

}

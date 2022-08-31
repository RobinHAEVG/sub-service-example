package main

import (
	"context"
	"fmt"
	"os/signal"
	"subServiceSystem/internal/certmgr"
	"subServiceSystem/internal/config"
	"subServiceSystem/internal/global"
	"subServiceSystem/internal/serviceA"
	"subServiceSystem/internal/serviceB"
	"subServiceSystem/internal/serviceC"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	app := &mainApp{
		services: make([]global.SubService, 0),
	}

	app.registerService(&serviceA.ServiceA{})
	app.registerService(&serviceB.ServiceB{})
	app.registerService(&serviceC.ServiceC{})

	app.Run()
	fmt.Println("finished!")
}

type mainApp struct {
	services []global.SubService
	ctx      context.Context
	cf       context.CancelFunc
}

func (app *mainApp) registerService(svc global.SubService) {
	app.services = append(app.services, svc)
}

func (app *mainApp) Run() {

	config := new(config.Configuration)
	certMgr := new(certmgr.CertManager)
	logger := new(logrus.Entry)

	app.ctx, app.cf = signal.NotifyContext(context.Background(), syscall.SIGHUP,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	wg := new(sync.WaitGroup)

	wg.Add(len(app.services))
	for _, s := range app.services {
		go func(ss global.SubService) {
			ss.Run(app.ctx, wg, logger, config, certMgr)
		}(s)
	}

	wg.Wait()
}

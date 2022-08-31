package global

import (
	"context"
	"sync"

	"subServiceSystem/internal/certmgr"
	"subServiceSystem/internal/config"

	"github.com/sirupsen/logrus"
)

type SubService interface {
	Run(context.Context, *sync.WaitGroup, *logrus.Entry, *config.Configuration, *certmgr.CertManager)
}

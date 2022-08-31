package serviceA

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"subServiceSystem/internal/certmgr"
	"subServiceSystem/internal/config"

	"github.com/sirupsen/logrus"
)

type ServiceA struct{}

func (sa *ServiceA) Run(ctx context.Context, wg *sync.WaitGroup, logger *logrus.Entry, cfg *config.Configuration, certMgr *certmgr.CertManager) {
	defer wg.Done()
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Service A!")
	})

	srv := http.Server{
		Addr:    "localhost:8001",
		Handler: r,
	}

	go func() {
		<-ctx.Done()

		if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, context.Canceled) {
			fmt.Printf("failed to shut down Service A server: %s\n", err.Error())
		}
	}()

	fmt.Println("starting service A on :8001")
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server error:", err.Error())
	}
}

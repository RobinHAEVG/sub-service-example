package serviceC

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

type ServiceC struct {
}

func (sc *ServiceC) Run(ctx context.Context, wg *sync.WaitGroup, logger *logrus.Entry, cfg *config.Configuration, certMgr *certmgr.CertManager) {
	defer wg.Done()
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Service C!")
	})

	srv := http.Server{
		Addr:    "localhost:8003",
		Handler: r,
	}

	go func() {
		<-ctx.Done()

		if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, context.Canceled) {
			fmt.Printf("failed to shut down Service C server: %s\n", err.Error())
		}
	}()

	fmt.Println("starting service C on :8003")
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server error:", err.Error())
	}
}

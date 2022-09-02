package serviceB

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

type ServiceB struct {
}

func (sb *ServiceB) Run(ctx context.Context, wg *sync.WaitGroup, logger *logrus.Entry, cfg *config.Configuration, certMgr *certmgr.CertManager) {
	defer wg.Done()
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Service B!")
	})

	srv := http.Server{
		Addr:    "localhost:8002",
		Handler: r,
	}

	go func() {
		<-ctx.Done()

		// give the server 10 seconds to shut down
		ctx2, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()
		
		if err := srv.Shutdown(ctx2); err != nil && !errors.Is(err, context.Canceled) {
			fmt.Printf("failed to shut down Service B server: %s\n", err.Error())
		}
	}()

	fmt.Println("starting service B on :8002")
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server error:", err.Error())
	}
}

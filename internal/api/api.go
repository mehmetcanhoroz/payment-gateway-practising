package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/config"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/pkg/imposter_bank"
	"github.com/go-chi/chi/v5/middleware"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"

	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/repository"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/services"
)

type Api struct {
	router           *chi.Mux
	paymentsRepo     *repository.PaymentsRepository
	paymentsFastRepo *repository.PaymentsFastRepository
	paymentsService  *services.PaymentsService
}

func New() *Api {
	// TODO: Move resources to main level not api level
	a := &Api{}
	a.paymentsRepo = repository.NewPaymentsRepository()
	a.paymentsFastRepo = repository.NewPaymentsFastRepository()

	bankHTTPClient := &http.Client{
		Timeout: time.Second * 60,
	}
	imposterBankConfig := imposter_bank.ConnectorConfig{
		HTTPClient: bankHTTPClient,
		BankURL:    config.Config.Bank.PosURL,
	}
	bankConnector := imposter_bank.NewImposterBankConnector(imposterBankConfig)

	a.paymentsService = services.NewPaymentsService(a.paymentsRepo, a.paymentsFastRepo, bankConnector)

	a.setupRouter()

	return a
}

func (a *Api) Run(ctx context.Context, addr string) error {
	httpServer := &http.Server{
		Addr:        addr,
		Handler:     a.router,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-ctx.Done()
		fmt.Printf("shutting down HTTP server\n")
		return httpServer.Shutdown(ctx)
	})

	g.Go(func() error {
		fmt.Printf("starting HTTP server on %s\n", addr)
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	return g.Wait()
}

func (a *Api) setupRouter() {
	a.router = chi.NewRouter()
	a.router.Use(middleware.Logger)

	a.router.Get("/ping", a.PingHandler())
	a.router.Get("/swagger/*", a.SwaggerHandler())

	paymentsHandler := a.GetPaymentHandlers()
	a.router.Get("/api/payments/{id}", paymentsHandler.GetHandler())
	a.router.Post("/api/payments", paymentsHandler.PostHandler())
}

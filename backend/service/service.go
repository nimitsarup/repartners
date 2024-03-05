package service

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nimitsarup/rep/api"
	"github.com/nimitsarup/rep/config"
	"github.com/nimitsarup/rep/db"
	"github.com/nimitsarup/rep/handlers"
	"github.com/pkg/errors"
)

type Service struct {
	Server      HTTPServer
	Router      *mux.Router
	ServiceList ServiceContainer
	DB          db.PacksInMemoryDB
}

func Run(ctx context.Context,
	services ServiceContainer,
	svcErrors chan error,
	cfg *config.Config, r *mux.Router) (*Service, error) {

	log.Println("starting service")

	db := services.GetDB()
	api := &api.API{Handlers: &handlers.Handlers{DB: db}}

	// use middleware for logging, cors
	// also can be used for auth etc
	r.Use(LoggingMiddleware)
	r.Use(SetCors)

	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "*")
	})

	// setup routes
	// 1. PUT for update pack sizes
	// 2. GET for getting pack configuration for items
	r.Path("/packs").HandlerFunc(api.UpdatePacks).Methods(http.MethodPut)
	r.Path("/packs").HandlerFunc(api.GetPacksForItems).Methods(http.MethodGet)

	s := services.GetHTTPServer()

	// Run the http server async
	go func() {
		log.Printf("listening at port [%s]", cfg.PortNumber)
		if err := s.ListenAndServe(); err != nil {
			svcErrors <- errors.Wrap(err, "failure in http listen and serve")
		}
	}()

	return &Service{Server: s, Router: r, ServiceList: services, DB: db}, nil
}

func (svc *Service) Close(ctx context.Context, timeout time.Duration) error {
	log.Println("commencing shutdown")
	ctx, cancel := context.WithTimeout(ctx, timeout)
	var err error
	go func() {
		defer cancel()
		err = svc.ServiceList.Shutdown(ctx)
	}()

	// wait for shutdown
	<-ctx.Done()

	// timeout expired
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("shutdown timed out")
		return ctx.Err()
	}
	// other error
	if err != nil {
		log.Println(ctx, "failed to shutdown gracefully ", err)
		return err
	}

	log.Println("shutdown successful")
	return nil
}

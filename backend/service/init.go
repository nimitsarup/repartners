package service

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nimitsarup/rep/config"
	"github.com/nimitsarup/rep/db"
)

type ExternalServices struct {
	cfg        *config.Config
	db         db.PacksInMemoryDB
	httpServer HTTPServer
	router     *mux.Router
}

func NewServices(cfg *config.Config,
	router *mux.Router) (*ExternalServices, error) {
	e := &ExternalServices{
		cfg:    cfg,
		router: router,
	}

	return e, e.setup(cfg)
}

func (e *ExternalServices) GetHTTPServer() HTTPServer {
	return e.httpServer
}

func (e *ExternalServices) GetDB() db.PacksInMemoryDB {
	return e.db
}

func (e *ExternalServices) Shutdown(ctx context.Context) error {
	// any shutdown logic
	return e.httpServer.Shutdown(ctx)
}

func (e *ExternalServices) setup(cfg *config.Config) error {
	e.createDB()
	e.createHttpServer()
	return nil
}

func (e *ExternalServices) createHttpServer() {
	s := &http.Server{
		Handler: e.router,
		Addr:    e.cfg.PortNumber,
	}
	e.httpServer = s
}

func (e *ExternalServices) createDB() {
	e.db = db.New()
}

package service

import (
	"context"

	"github.com/nimitsarup/rep/db"
)

//go:generate moq -out mock/ServiceContainer.go -pkg mock . ServiceContainer

type HTTPServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type ServiceContainer interface {
	GetHTTPServer() HTTPServer
	GetDB() db.PacksInMemoryDB
	Shutdown(ctx context.Context) error
}

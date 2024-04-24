package main

import (
	"context"
	"database/sql"

	"github.com/ServiceWeaver/weaver"
)

// Reverser component.
type Reverser interface {
	Reverse(context.Context, string) (string, error)
	Active(context.Context, string) (string, error)
	Init(ctx context.Context) error
}

// Implementation of the Reverser component.
type reverser struct {
	weaver.Implements[Reverser]
	weaver.WithConfig[config]
	db *sql.DB
}

type config struct {
	Driver string // Name of the DB driver.
	Source string // DB data source.
}

func (r *reverser) Init(ctx context.Context) error {
	logger := r.Logger(ctx)
	logger.Info(r.Config().Driver)
	logger.Info(r.Config().Source)

	db, err := sql.Open(r.Config().Driver, r.Config().Source)
	r.db = db

	return err
}

func (r *reverser) Reverse(ctx context.Context, s string) (string, error) {
	logger := r.Logger(ctx)
	logger.Info("lalalala")
	if s == "wyf" {
		panic("Something went wrong!")

	}
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes), nil
}

func (r *reverser) Active(ctx context.Context, s string) (string, error) {

	return "", nil

}

package main

import (
	"net/http"
	"regexp"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ThoughtHandler struct {
	db *pgxpool.Pool
}

var (
	ThoughtRE = regexp.MustCompile(`^\/thoughts\/?$`)
)

func newThoughtHandler(db *pgxpool.Pool) *ThoughtHandler {
	return &ThoughtHandler{db: db}
}

func (h *ThoughtHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	}
}

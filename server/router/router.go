package router

import (
	"bookshelf/server/app"
	"bookshelf/server/requestlog"

	"github.com/go-chi/chi"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()

	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))

	return r
}

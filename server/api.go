package main

import (
	"errors"
	"net/http"

	"github.com/FyraLabs/subatomic/server/ent"
	"github.com/FyraLabs/subatomic/server/keyedmutex"
	"github.com/FyraLabs/subatomic/server/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/jwtauth/v5"
)

type apiRouter struct {
	*chi.Mux
	database         *ent.Client
	enviroment       *types.Enviroment
	jwtAuthenticator *jwtauth.JWTAuth
	repoMutex        *keyedmutex.KeyedMutex
}

func (router *apiRouter) setup() {
	router.Mux = chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/heartbeat"))
	router.Use(middleware.Recoverer)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		if err := render.Render(w, r, types.ErrNotFound(errors.New("route not found"))); err != err {
			panic(err)
		}
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		if err := render.Render(w, r, types.MethodNotAllowed(errors.New("method not allowed for route"))); err != err {
			panic(err)
		}
	})

	// Authenticated
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(router.jwtAuthenticator))
		r.Use(Authenticator)

		repos := reposRouter{
			database:   router.database,
			enviroment: router.enviroment,
			repoMutex:  router.repoMutex,
		}
		repos.setup()
		r.Mount("/repos", repos)

		keys := keysRouter{
			database:   router.database,
			enviroment: router.enviroment,
		}
		keys.setup()
		r.Mount("/keys", keys)
	})

	// Public
	router.Group(func(r chi.Router) {
		router.Handle("/docs", http.RedirectHandler("/docs/index.html", http.StatusFound))
		router.Get("/docs/*", httpSwagger.Handler(
			httpSwagger.URL("/docs/doc.json"), //The url pointing to API definition
		))
	})
}

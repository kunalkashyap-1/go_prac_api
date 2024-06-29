package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/kunalkashyap-1/go_prac_api/internal/middleware"
)

func Handler(r *chi.Mux){
	r.Use(chimiddle.StripSlashes)

	r.Route("/accounts", func(router chi.Router){
		r.Use(middleware.Autherization)
		router.Get("/coins", GetCoinBalance)
	})
}
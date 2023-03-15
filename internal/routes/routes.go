package routes

import (
	"github.com/go-chi/chi"
	"net/http"
)

func Routes() (r *chi.Mux) {
	r = chi.NewRouter()

	http.Handle("/", r)

	return
}

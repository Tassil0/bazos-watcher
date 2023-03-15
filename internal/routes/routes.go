package routes

import (
	"encoding/json"
	"github.com/Tassil0/bazos-watcher/internal/database"
	"github.com/go-chi/chi"
	"net/http"
)

func Routes() (r *chi.Mux) {
	r = chi.NewRouter()

	//r.Use(render.SetContentType(render.ContentTypeJSON))

	http.Handle("/", r)

	r.Get("/queries", queries)

	return
}

func queries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := database.GetQueries()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	res, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

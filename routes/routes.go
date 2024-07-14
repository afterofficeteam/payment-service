package routes

import (
	"base-project/config"
	"base-project/util/helper"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Routes struct {
	Router *mux.Router
}

func EnabledCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PUT, PATCH")
		header.Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (r *Routes) Run(port string) {
	r.SetupRouter()

	log.Printf("[HTTP SRV] clients on localhost port :%s", port)
	srv := &http.Server{
		Handler:      r.Router,
		Addr:         "localhost:" + port,
		WriteTimeout: config.WriteTimeout() * time.Second,
		ReadTimeout:  config.ReadTimeout() * time.Second,
	}

	log.Panic(srv.ListenAndServe())
}

func (r *Routes) SetupRouter() {
	r.Router = mux.NewRouter()
	r.Router.Use(EnabledCors, helper.LoggerMiddleware())

	r.SetupBaseURL()
}

func (r *Routes) SetupBaseURL() {
	baseURL := viper.GetString("BASE_URL_PATH")
	if baseURL != "" && baseURL != "/" {
		r.Router.PathPrefix(baseURL).HandlerFunc(helper.URLRewriter(r.Router, baseURL))
	}
}

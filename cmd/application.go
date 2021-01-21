package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"wateringsystem/pkg/wateringsystem"

	"github.com/julienschmidt/httprouter"
)

type application struct {
	ws *wateringsystem.WateringSystem
}

func wrapHandler(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		ctx := r.Context()

		ctx = context.WithValue(ctx, "params", p)

		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/health", wrapHandler(healthCheckHandler))
	router.GET("/status", wrapHandler(app.ws.GetStatus))
	router.GET("/history", wrapHandler(app.ws.GetHistory))
	router.POST("/water", wrapHandler(app.ws.TurnOnWateringSystem))
	router.POST("/timed", wrapHandler(app.ws.TimedWateringSystem))

	//Just to show how to retrieve the params from httprouter.Params
	router.GET("/user/:id/post/:postId", wrapHandler(func(rw http.ResponseWriter, r *http.Request) {
		values := r.Context().Value("params").(httprouter.Params)
		log.Println(values.ByName("id"))
	}))

	return router
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"healthy": true}`)
}

func (app *application) RunServer() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	log.Println("Listening at :8080")
	srv.ListenAndServe()
}

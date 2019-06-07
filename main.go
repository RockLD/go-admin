package main

import (
	r "go-admin/router"
	"net/http"
	"time"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("go-admin"))
func main() {
	router := r.Engine()

	s := &http.Server{
		Addr:":8090",
		Handler:router,
		ReadTimeout:10 * time.Second,
		WriteTimeout:10 * time.Second,
	}

	s.ListenAndServe()
}

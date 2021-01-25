package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type httpServer struct {
	server       *http.Server
}

func NewServer(host, port string, keycloak *keycloak) *httpServer {

	// create a root router
	router := mux.NewRouter()

	// add a subrouter based on matcher func
	// note, routers are processed one by one in order, so that if one of the routing matches other won't be processed
	noAuthRouter := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return r.Header.Get("Authorization") == ""
	}).Subrouter()

	// add one more subrouter for the authenticated service methods
	authRouter := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return true
	}).Subrouter()

	// instantiate a new controller which is supposed to serve our routes
	controller := newController(keycloak)

	// map url routes to controller's methods
	noAuthRouter.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		controller.login(writer, request)
	}).Methods("POST")

	authRouter.HandleFunc("/docs", func(writer http.ResponseWriter, request *http.Request) {
		controller.getDocs(writer, request)
	}).Methods("GET")

	// apply middleware
	mdw := newMiddleware(keycloak)
	authRouter.Use(mdw.verifyToken)

	// create a server object
	s := &httpServer{
		server: &http.Server{
			Addr:         fmt.Sprintf("%s:%s", host, port),
			Handler:      router,
			WriteTimeout: time.Hour,
			ReadTimeout:  time.Hour,
		},
	}

	return s
}

func (s *httpServer) listen() error {
	return s.server.ListenAndServe()
}

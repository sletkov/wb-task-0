package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sletkov/backend/wb-task-0/internal/app/store"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
	cache  cache.Cache
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
		cache:  *cache.New(cache.NoExpiration, cache.NoExpiration),
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/orders/{id}", s.handleOrdersGetById()).Methods("GET")
}

func (s *server) handleOrdersGetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderId := mux.Vars(r)["id"]

		order, ok := s.cache.Get(orderId)

		if !ok {
			s.error(w, r, http.StatusNotFound, errors.New(fmt.Sprintf("Can't find order with id: %s", orderId)))
		}

		s.respond(w, r, http.StatusFound, order)
	}

}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
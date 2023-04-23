package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	DefaultPort = "8080"
)

type MethodToAdd struct {
	Verb    string
	Path    string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Methods       []MethodToAdd
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddMethod(verb string, path string, handler http.HandlerFunc) {
	// s.Router.MethodFunc(verb, path, handler)
	s.Methods = append(s.Methods, MethodToAdd{
		Verb:    verb,
		Path:    path,
		Handler: handler,
	})
}

func (s *WebServer) Start() {
	r := s.Router
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	if s.WebServerPort == "" {
		s.WebServerPort = DefaultPort
	}

	for _, method := range s.Methods {
		fmt.Println("Adding method", method.Verb, method.Path)
		r.MethodFunc(method.Verb, method.Path, method.Handler)
	}

	fmt.Println("Starting web server on port", s.WebServerPort)
	err := http.ListenAndServe(":"+s.WebServerPort, r)
	if err != nil {
		panic(err)
	}
}

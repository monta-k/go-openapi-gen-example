package user

import (
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	*http.Server
}

func NewServer() (*Server, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT is not set")
	}

	deps := &Dependency{}
	deps.Inject()

	mux := routing(deps)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	return &Server{&server}, nil
}

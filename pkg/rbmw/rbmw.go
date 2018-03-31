package rbmw

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Backend interface {
}

type HandlerConfig struct {
	Backend Backend
}

func NewHandler(config *HandlerConfig) (*Handler, error) {
	h := &Handler{
		config: config,
	}
	return h, nil
}

type Handler struct {
	config *HandlerConfig
}

func (h *Handler) AddHandlers(mux *http.ServeMux, prefix string) {
	var routes = []struct {
		Path      string
		Handler   http.Handler
		Protected bool
	}{
		{"/login", http.HandlerFunc(h.loginHandler), false},
	}

	prefix = strings.TrimRight(prefix, "/")
	for _, route := range routes {
		resource := prefix + route.Path
		mux.Handle(resource, route.Handler)
	}
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	payload := &LoginResponse{
		Token: "TBD",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

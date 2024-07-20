package http

import "net/http"

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)

}

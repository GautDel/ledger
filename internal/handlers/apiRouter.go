package handlers

import (
	"net/http"
)

func ApiRouter(mux *http.ServeMux) {

    mux.HandleFunc("/api/", homeHandler)
}

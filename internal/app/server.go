package app

import "net/http"

func SetMux() *http.ServeMux {
	mux := http.NewServeMux()

	//mux.HandleFunc("")

	return mux
}

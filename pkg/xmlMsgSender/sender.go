package xmlsender

import (
	"encoding/xml"
	"net/http"
)

type Response struct {
	Message string `xml:"message"`
}

func Sender(w http.ResponseWriter, status int, message string) {
	resp := Response{Message: message}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	xml.NewEncoder(w).Encode(resp)
}

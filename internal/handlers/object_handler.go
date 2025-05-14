package handlers

import (
	"GonIO/internal/domain"
	"net/http"
)

type ObjectHandler struct {
	serv domain.ObjectService
}

func NewObjectHandler(serv domain.ObjectService) *ObjectHandler {
	return &ObjectHandler{serv: serv}
}

func (h *ObjectHandler) GetObjectList(w http.ResponseWriter, r *http.Request) {

}

func (h *ObjectHandler) RetrieveObject(w http.ResponseWriter, r *http.Request) {

}

func (h *ObjectHandler) UpdateObject(w http.ResponseWriter, r *http.Request) {

}

func (h *ObjectHandler) DeleteObject(w http.ResponseWriter, r *http.Request) {

}

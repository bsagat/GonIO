package handlers

import (
	"GonIO/internal/domain"
	xmlsender "GonIO/pkg/xmlMsgSender"
	"net/http"
)

type BucketHandler struct {
	serv domain.BucketService
}

func NewBucketHandler(serv domain.BucketService) *BucketHandler {
	return &BucketHandler{serv: serv}
}

func (h *BucketHandler) BucketListsHandler(w http.ResponseWriter, r *http.Request) {
	h.serv.BucketList(w)
}

func (h *BucketHandler) CreateBucketHandler(w http.ResponseWriter, r *http.Request) {
	bucketName := r.FormValue("BucketName")
	if bucketName == "" {
		xmlsender.Sender(w, http.StatusBadRequest, domain.ErrEmptyBucketName.Error())
		return
	}
}

func (h *BucketHandler) DeleteBucketHandler(w http.ResponseWriter, r *http.Request) {

}

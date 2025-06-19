package handlers

import (
	"GonIO/internal/domain"
	"log/slog"
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
	bucketName := r.PathValue("BucketName")
	if bucketName == "" {
		slog.Error("Failed to get path value: ", "error", domain.ErrEmptyBucketName)
		http.Error(w, domain.ErrEmptyBucketName.Error(), http.StatusBadRequest)
		return
	}

	h.serv.CreateBucket(w, bucketName)
}

func (h *BucketHandler) DeleteBucketHandler(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	if bucketName == "" {
		slog.Error("Failed to get path value: ", "error", domain.ErrEmptyBucketName)
		http.Error(w, domain.ErrEmptyBucketName.Error(), http.StatusBadRequest)
		return
	}

	h.serv.DeleteBucket(w, bucketName)
}

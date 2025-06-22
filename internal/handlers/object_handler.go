package handlers

import (
	"GonIO/internal/domain"
	xmlsender "GonIO/pkg/xmlMsgSender"
	"fmt"
	"log/slog"
	"net/http"
)

type ObjectHandler struct {
	serv domain.ObjectService
}

func NewObjectHandler(serv domain.ObjectService) *ObjectHandler {
	return &ObjectHandler{serv: serv}
}

func (h *ObjectHandler) GetObjectList(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	if bucketName == "" {
		slog.Error("Failed to get path value: ", "error", domain.ErrEmptyBucketName)
		http.Error(w, domain.ErrEmptyBucketName.Error(), http.StatusBadRequest)
		return
	}

	objectList, code, err := h.serv.ObjectList(bucketName)
	if err != nil {
		slog.Error("Failed to get object list: ", "error", err)
		http.Error(w, err.Error(), code)
		return
	}

	if err = xmlsender.SendObjectList(w, objectList); err != nil {
		slog.Error("Failed to send object list: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ObjectHandler) RetrieveObject(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	if bucketName == "" {
		slog.Error("Failed to get path value: ", "error", domain.ErrEmptyBucketName)
		http.Error(w, domain.ErrEmptyBucketName.Error(), http.StatusBadRequest)
		return
	}

	objectName := r.PathValue("ObjectKey")
	if objectName == "" {
		slog.Error("Failed to get path value: ", "error", domain.ErrEmptyObjectName)
		http.Error(w, domain.ErrEmptyObjectName.Error(), http.StatusBadRequest)
		return
	}

	code, err := h.serv.RetrieveObject(w, bucketName, objectName)
	if err != nil {
		slog.Error("Failed to retrieve object: ", "bucket", bucketName, "object", objectName, "error", err)
		http.Error(w, err.Error(), code)
		return
	}
}

func (h *ObjectHandler) UpdateObject(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	if bucketName == "" {
		slog.Error("Failed to get path value: ", "error", domain.ErrEmptyBucketName)
		http.Error(w, domain.ErrEmptyBucketName.Error(), http.StatusBadRequest)
		return
	}

	objectName := r.PathValue("ObjectKey")
	if objectName == "" {
		slog.Error("Failed to get path value: ", "error", domain.ErrEmptyObjectName)
		http.Error(w, domain.ErrEmptyObjectName.Error(), http.StatusBadRequest)
		return
	}

	code, err := h.serv.UploadObject(r, bucketName, objectName)
	if err != nil {
		slog.Error("Failed to upload object: ", "bucket", bucketName, "object", objectName, "error", err)
		http.Error(w, err.Error(), code)
		return
	}

	if err = xmlsender.SendMessage(w, code, fmt.Sprintf("object with name %s created succesfully", objectName)); err != nil {
		slog.Error("Failed to send xml message: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ObjectHandler) DeleteObject(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	if bucketName == "" {
		slog.Error("Failed to get path value: ", "error", domain.ErrEmptyBucketName)
		http.Error(w, domain.ErrEmptyBucketName.Error(), http.StatusBadRequest)
		return
	}

	objectName := r.PathValue("ObjectKey")
	if objectName == "" {
		slog.Error("Failed to get path value: ", "error", domain.ErrEmptyObjectName)
		http.Error(w, domain.ErrEmptyObjectName.Error(), http.StatusBadRequest)
		return
	}

	code, err := h.serv.DeleteObject(bucketName, objectName)
	if err != nil {
		slog.Error("Failed to delete object: ", "bucket", bucketName, "object", objectName, "error", err)
		http.Error(w, err.Error(), code)
		return
	}

	if err = xmlsender.SendMessage(w, code, fmt.Sprintf("object with name %s deleted succesfully", objectName)); err != nil {
		slog.Error("Failed to send xml message: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

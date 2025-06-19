package service

import (
	"GonIO/internal/dal"
	"GonIO/internal/domain"
	xmlsender "GonIO/pkg/xmlMsgSender"
	"encoding/xml"
	"fmt"
	"log/slog"
	"net/http"
)

type ObjectServiceImp struct {
	dal domain.ObjectDal
}

var _ domain.ObjectService = (*ObjectServiceImp)(nil)

func NewObjectService(dal domain.ObjectDal) *ObjectServiceImp {
	return &ObjectServiceImp{dal: dal}
}

// Object list retrieve logics
func (serv ObjectServiceImp) ObjectList(w http.ResponseWriter, bucketname string) {
	unique, err := dal.NewBucketXMLRepo().IsUniqueBucket(bucketname)
	if err != nil {
		slog.Error("Failed to check if bucket name is unique: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if unique {
		slog.Info("Bucket is not exist")
		http.Error(w, domain.ErrBucketIsNotExist.Error(), http.StatusNotFound)
		return
	}

	objectList, err := serv.dal.List_Object(bucketname)
	if err != nil {
		slog.Error("Failed to retrieve object list: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	if err := xml.NewEncoder(w).Encode(objectList); err != nil {
		slog.Error("Failed to encode response: ", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// Object retrieve logic
func (serv ObjectServiceImp) RetrieveObject(w http.ResponseWriter, bucketname, objectname string) {
	unique, err := dal.NewBucketXMLRepo().IsUniqueBucket(bucketname)
	if err != nil {
		slog.Error("Failed to check if bucket name exists: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if unique {
		slog.Info("Bucket is not exist")
		http.Error(w, domain.ErrBucketIsNotExist.Error(), http.StatusNotFound)
		return
	}

	if err = Validate(objectname); err != nil {
		slog.Error("Object name validation error: ", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exist, err := serv.dal.IsObjectExist(domain.BucketsPath+"/"+bucketname+"/objects.csv", objectname)
	if err != nil {
		slog.Error("Failed to check if object name exists: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exist {
		slog.Info("Object is not exist")
		http.Error(w, domain.ErrObjectIsNotExist.Error(), http.StatusNotFound)
		return
	}

	if err = serv.dal.RetrieveObject(bucketname, objectname, w); err != nil {
		slog.Error("Failed to retrieve object: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Object upload logic
func (serv ObjectServiceImp) UploadObject(w http.ResponseWriter, r *http.Request, bucketname, objectname string) {
	unique, err := dal.NewBucketXMLRepo().IsUniqueBucket(bucketname)
	if err != nil {
		slog.Error("Failed to check if bucket name is unique: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if unique {
		slog.Info("Bucket is not exist")
		http.Error(w, domain.ErrBucketIsNotExist.Error(), http.StatusNotFound)
		return
	}

	if err = Validate(objectname); err != nil {
		slog.Error("Object name validation error: ", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = serv.dal.UploadObject(bucketname, objectname, r); err != nil {
		slog.Error("Failed to upload object: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = xmlsender.SendMessage(w, http.StatusCreated, fmt.Sprintf("object with name %s created succesfully", objectname)); err != nil {
		slog.Error("Failed to send xml message: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Object delete logic
func (serv ObjectServiceImp) DeleteObject(w http.ResponseWriter, r *http.Request, bucketname, objectname string) {
	unique, err := dal.NewBucketXMLRepo().IsUniqueBucket(bucketname)
	if err != nil {
		slog.Error("Failed to check if bucket name is unique: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if unique {
		slog.Info("Bucket is not exist")
		http.Error(w, domain.ErrBucketIsNotExist.Error(), http.StatusNotFound)
		return
	}

	if err = Validate(objectname); err != nil {
		slog.Error("Object name validation error: ", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exist, err := serv.dal.IsObjectExist(domain.BucketsPath+"/"+bucketname+"/objects.csv", objectname)
	if err != nil {
		slog.Error("Failed to check if object name exists: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exist {
		slog.Info("Object is not exist")
		http.Error(w, domain.ErrObjectIsNotExist.Error(), http.StatusNotFound)
		return
	}

	if err := serv.dal.DeleteObject(bucketname, objectname); err != nil {
		slog.Error("Failed to delete object : ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = xmlsender.SendMessage(w, http.StatusOK, fmt.Sprintf("object with name %s deleted succesfully", objectname)); err != nil {
		slog.Error("Failed to send xml message: ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

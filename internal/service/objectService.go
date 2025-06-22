package service

import (
	"GonIO/internal/dal"
	"GonIO/internal/domain"
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
func (serv ObjectServiceImp) ObjectList(bucketname string) ([]domain.Object, int, error) {
	unique, err := dal.NewBucketXMLRepo().IsUniqueBucket(bucketname)
	if err != nil {
		slog.Error("Failed to check if bucket name is unique: ", "error", err)
		return nil, http.StatusInternalServerError, err
	}
	if unique {
		slog.Info("Bucket is not exist")
		return nil, http.StatusNotFound, domain.ErrBucketIsNotExist
	}

	objectList, err := serv.dal.List_Object(bucketname)
	if err != nil {
		slog.Error("Failed to retrieve object list: ", "error", err)
		return nil, http.StatusInternalServerError, err
	}
	return objectList, http.StatusOK, nil
}

// Object retrieve logic
func (serv ObjectServiceImp) RetrieveObject(w http.ResponseWriter, bucketname, objectname string) (int, error) {
	unique, err := dal.NewBucketXMLRepo().IsUniqueBucket(bucketname)
	if err != nil {
		slog.Error("Failed to check if bucket name exists: ", "error", err)
		return http.StatusInternalServerError, err
	}
	if unique {
		slog.Info("Bucket is not exist")
		return http.StatusNotFound, domain.ErrBucketIsNotExist
	}

	if err = Validate(objectname); err != nil {
		slog.Error("Object name validation error: ", "error", err)
		return http.StatusBadRequest, err
	}

	exist, err := serv.dal.IsObjectExist(domain.BucketsPath+"/"+bucketname+"/objects.csv", objectname)
	if err != nil {
		slog.Error("Failed to check if object name exists: ", "error", err)
		return http.StatusInternalServerError, err
	}

	if !exist {
		slog.Info("Object is not exist")
		return http.StatusNotFound, domain.ErrObjectIsNotExist
	}

	if err = serv.dal.RetrieveObject(bucketname, objectname, w); err != nil {
		slog.Error("Failed to retrieve object: ", "error", err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// Object upload logic
func (serv ObjectServiceImp) UploadObject(r *http.Request, bucketname, objectname string) (int, error) {
	unique, err := dal.NewBucketXMLRepo().IsUniqueBucket(bucketname)
	if err != nil {
		slog.Error("Failed to check if bucket name is unique: ", "error", err)
		return http.StatusInternalServerError, err
	}
	if unique {
		slog.Info("Bucket is not exist")
		return http.StatusNotFound, domain.ErrBucketIsNotExist
	}

	if err = Validate(objectname); err != nil {
		slog.Error("Object name validation error: ", "error", err)
		return http.StatusBadRequest, err
	}

	if err = serv.dal.UploadObject(bucketname, objectname, r); err != nil {
		slog.Error("Failed to upload object: ", "error", err)
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

// Object delete logic
func (serv ObjectServiceImp) DeleteObject(bucketname, objectname string) (int, error) {
	unique, err := dal.NewBucketXMLRepo().IsUniqueBucket(bucketname)
	if err != nil {
		slog.Error("Failed to check if bucket name is unique: ", "error", err)
		return http.StatusInternalServerError, err
	}
	if unique {
		slog.Info("Bucket is not exist")
		return http.StatusNotFound, domain.ErrBucketIsNotExist
	}

	if err = Validate(objectname); err != nil {
		slog.Error("Object name validation error: ", "error", err)
		return http.StatusBadRequest, err
	}

	exist, err := serv.dal.IsObjectExist(domain.BucketsPath+"/"+bucketname+"/objects.csv", objectname)
	if err != nil {
		slog.Error("Failed to check if object name exists: ", "error", err)
		return http.StatusInternalServerError, err
	}

	if !exist {
		slog.Info("Object is not exist")
		return http.StatusNotFound, domain.ErrObjectIsNotExist
	}

	if err := serv.dal.DeleteObject(bucketname, objectname); err != nil {
		slog.Error("Failed to delete object : ", "error", err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (serv *ObjectServiceImp) UploadObjectjar() (int, error) {

	return http.StatusCreated, nil
}

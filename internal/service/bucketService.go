package service

import (
	"GonIO/internal/domain"
	"log/slog"
	"net/http"
)

type BucketServiceImp struct {
	dal domain.BucketDal
}

var _ domain.BucketService = (*BucketServiceImp)(nil)

func NewBucketService(dal domain.BucketDal) *BucketServiceImp {
	return &BucketServiceImp{dal: dal}
}

// Bucket List retrieve logic
func (serv BucketServiceImp) BucketList() ([]domain.Bucket, error) {
	list, err := serv.dal.GetBucketList()
	if err != nil {
		slog.Error("Failed to get bucket list: ", "error", err)
		return nil, err
	}
	return list, nil
}

// Bucket create logic (validation)
func (serv BucketServiceImp) CreateBucket(bucketName string) (int, error) {
	unique, err := serv.dal.IsUniqueBucket(bucketName)
	if err != nil {
		slog.Error("Failed to check if bucket name is unique: ", "error", err)
		return http.StatusInternalServerError, err
	}
	if !unique {
		slog.Info("Bucket name is not unique")
		return http.StatusConflict, domain.ErrNotUniqueName
	}

	if err = Validate(bucketName); err != nil {
		slog.Error("Bucket name validation error: ", "error", err)
		return http.StatusBadRequest, err
	}

	if err = serv.dal.CreateBucket(bucketName); err != nil {
		slog.Error("Failed to create bucket: ", "error", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

// Bucket Delete logic
func (serv BucketServiceImp) DeleteBucket(bucketName string) (int, error) {
	unique, err := serv.dal.IsUniqueBucket(bucketName)
	if err != nil {
		slog.Error("Failed to check if bucket name is unique: ", "error", err)
		return http.StatusInternalServerError, err
	}
	if unique {
		slog.Info("Bucket is not exist")
		return http.StatusNotFound, domain.ErrBucketIsNotExist
	}

	if err = serv.dal.DeleteBucket(bucketName); err != nil {
		slog.Error("Failed to delete bucket: ", "error", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

package domain

import (
	"net/http"
)

type BucketDal interface {
	GetBucketList() ([]Bucket, error)
	IsUniqueBucket(bucketName string) (bool, error)
	CreateBucket(bucketname string) error
	DeleteBucket(bucketName string) error
}

type ObjectDal interface {
	IsObjectExist(path, name string) (bool, error)
	List_Object(bucketname string) ([]Object, error)
	UploadObject(bucketname, objectname string, r *http.Request) error
	RetrieveObject(bucketname, objectname string, w http.ResponseWriter) error
	DeleteObject(bucketname, objectname string) error
}

type BucketService interface {
	BucketList() ([]Bucket, error)
	CreateBucket(bucketName string) (int, error)
	DeleteBucket(bucketName string) (int, error)
}

type ObjectService interface {
	ObjectList(bucketname string) ([]Object, int, error)
	RetrieveObject(w http.ResponseWriter, bucketname, objectname string) (int, error)
	UploadObject(r *http.Request, bucketname, objectname string) (int, error)
	DeleteObject(bucketname, objectname string) (int, error)
}

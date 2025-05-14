package domain

import "net/http"

type BucketDal interface {
	GetBucketList() ([]Bucket, error)
	IsUniqueBucket(bucketName string) (bool, error)
	CreateBucket(bucketname string) error
	DeleteBucket(bucketName string) error
}

type ObjectDal interface {
	RetrieveObject() error
}

type BucketService interface {
	BucketList(w http.ResponseWriter)
	CreateBucket(w http.ResponseWriter, bucketName string)
	DeleteBucket(w http.ResponseWriter, bucketName string)
}

type ObjectService interface {
	ObjectList(w http.ResponseWriter)
}

package domain

import "net/http"

type BucketDal interface {
	GetBucketList() ([]Bucket, error)
}

type ObjectDal interface {
	RetrieveObject() error
}

type BucketService interface {
	BucketList(w http.ResponseWriter)
}

type ObjectService interface {
	ObjectList(w http.ResponseWriter)
}

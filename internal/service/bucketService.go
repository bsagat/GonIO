package service

import (
	"GonIO/internal/domain"
	"fmt"
	"net/http"
)

type BucketService struct {
	dal domain.BucketDal
}

func NewBucketService(dal domain.BucketDal) *BucketService {
	return &BucketService{dal: dal}
}

func (serv BucketService) BucketList(w http.ResponseWriter) {
	fmt.Println(serv.dal.GetBucketList())
}

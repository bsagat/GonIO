package service

import (
	"GonIO/internal/domain"
	"net/http"
)

type ObjectService struct {
	dal domain.ObjectDal
}

func NewObjectService(dal domain.ObjectDal) *ObjectService {
	return &ObjectService{dal: dal}
}

func (serv ObjectService) ObjectList(w http.ResponseWriter) {

}

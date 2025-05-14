package service

import (
	"GonIO/internal/domain"
	"net/http"
)

type ObjectServiceImp struct {
	dal domain.ObjectDal
}

var _ domain.ObjectService = (*ObjectServiceImp)(nil)

func NewObjectService(dal domain.ObjectDal) *ObjectServiceImp {
	return &ObjectServiceImp{dal: dal}
}

func (serv ObjectServiceImp) ObjectList(w http.ResponseWriter) {

}

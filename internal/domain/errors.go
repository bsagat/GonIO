package domain

import "errors"

var (
	ErrInvalidPortStr  = errors.New("port number is invalid, must be in range between 1100 and 65535")
	ErrEmptyDomain     = errors.New("domain config is empty")
	ErrEmptyBucketName = errors.New("bucket name is empty")
)

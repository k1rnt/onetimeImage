package model

import (
	"golang.org/x/xerrors"
)

type Health struct {
	Status  string
	Message string
}

func NewHealth(status, message string) (*Health, error) {
	if status == "" {
		return nil, xerrors.New("required: status")
	}
	if message == "" {
		return nil, xerrors.New("required: message")
	}

	health := &Health{
		Status:  status,
		Message: message,
	}

	return health, nil
}

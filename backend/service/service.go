package service

import (
	"backend/model"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	validate *validator.Validate
}

func New() Service {
	v := validator.New()
	return Service{
		validate: v,
	}
}

func (s Service) UpdateConfig(config model.NetworkConf) error {
	err := s.validate.Struct(config)
	if err != nil {
		return err
	}
	fmt.Println(config)
	return nil
}

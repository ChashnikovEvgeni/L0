package service

import (
	"L0/app_data"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
)

// Валидация данных на соответствие типам данных и содержание
func (s *Service) Validation(Data *[]byte) bool {
	var validate *validator.Validate
	validate = validator.New()
	var data app_data.OrderJson
	var err error
	err = json.Unmarshal(*Data, &data)
	if err != nil {
		log.Println("Ошибка", err)
		return false
	}
	fmt.Println(*data.Orderuid)
	_, found := s.Cache.Get(*data.Orderuid)
	if found {
		log.Println("Данные с указанным order_uid уже находятся в системе")
		return false
	}
	err = validate.Struct(data)
	if err != nil {
		s.validbug_report(&err)
		return false
	}
	return true
}

// Вывод в log подробного отчёта о возникшей ошибки при осуществлении валидации
func (s *Service) validbug_report(err *error) {
	if _, ok := (*err).(*validator.InvalidValidationError); ok {
		log.Println("Ошибка", err)
	}
	for _, err := range (*err).(validator.ValidationErrors) {
		log.Println(err.Namespace())
		log.Println(err.Field())
		log.Println(err.StructNamespace())
		log.Println(err.StructField())
		log.Println(err.Tag())
		log.Println(err.ActualTag())
		log.Println(err.Kind())
		log.Println(err.Type())
		log.Println(err.Value())
		log.Println(err.Param())
		log.Println()
	}

}

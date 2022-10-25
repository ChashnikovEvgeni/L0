package service

import (
	"L0/models"
	"github.com/patrickmn/go-cache"
	"log"
)

// Создание кэша при запуске сервиса
// Данные считываются из БД и добавляются в кэш
func (s *Service) Make_start_cache() {
	rows := s.Get_all_orders()
	for rows.Next() {
		var order_uid string
		var mjson []byte
		rows.Scan(&order_uid, &mjson)
		s.Add_in_cache(&order_uid, &mjson)
	}
	log.Println("Данные  добавлены в кэш")

}

// Добавление данных одного заказа в кэш
func (s *Service) Add_in_cache(order_uid *string, mjson *[]byte) {
	order := models.Order{*order_uid, *mjson}
	s.Cache.Set(*order_uid, order, cache.NoExpiration)

}

package service

import (
	"L0/app_data"
	"database/sql"
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
)

// Получение списка всех заказов из БД
func (s *Service) Get_all_orders() *sql.Rows {
	rows, err := s.DB.Query("SELECT * FROM orders")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Данные успешно получены из бд.")
	}
	return rows
}

// Сохранение полученных данных в БД и кжш
func (s *Service) Save_data(order_uid *string, json_data *[]byte) {
	order := app_data.Order{*order_uid, *json_data}
	result, err := s.DB.Exec("insert into orders (Order_uid, json) values ($1, $2)",
		order_uid, json_data)
	if err != nil {
		panic(err)
	} else {
		log.Println("Данные успешно внесенны в БД")
	}
	s.Cache.Set(*order_uid, order, cache.NoExpiration)
	log.Println("Данные успешно внесенны в кэш")
	fmt.Println(result.RowsAffected())

}

// Подключение к базе данных
func (s *Service) DB_connect() *sql.DB {
	db, err := sql.Open("postgres", s.СonnStr)
	if err != nil {
		panic(err)
	}
	s.DB = db
	return db
}

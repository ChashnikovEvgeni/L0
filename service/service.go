package service

import (
	"database/sql"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type Service struct {
	Cache *cache.Cache
	DB    *sql.DB

	// Данные для подключения к БД
	ConnStr string

	// Данные для подключение к nats streaming service
	ClusterID string
	ClientID  string
}

// Инициализация сервиса, в случае отсутствия данных, полям присваиваются значения по умолчанию
func (s *Service) Set_config(connStr string, clusterID string, clientID string) bool {
	s.Cache = cache.New(5*time.Minute, 10*time.Minute)
	if connStr == "" {
		s.ConnStr = "user=postgres password=123 dbname=L0 sslmode=disable"
	} else {
		s.ConnStr = connStr
	}
	fmt.Println(s.ConnStr)

	if clusterID == "" {
		s.ClusterID = "test-cluster"
	} else {
		s.ClusterID = clusterID
	}

	if clientID == "" {
		s.ClientID = "clientidid"
	} else {
		s.ClientID = clientID
	}

	return true
}

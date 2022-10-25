package service

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"time"
)

// подключение к nats streamig server, подписка на тему и передача данных в функцию
// обработки полученных сообщений
func (s Service) Connect_listen() {
	sc, err := stan.Connect(s.ClusterID, s.ClientID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Начато прослушивание")
	aw, _ := time.ParseDuration("10s")
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		m.Ack()
		s.Receiving(m)
	}, stan.SetManualAckMode(), stan.AckWait(aw), stan.DurableName("clientid"))

	if err != nil {
		log.Fatal(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	s.close_connect(&signalChan, &sub, &sc)

}

// обработка полученных сообщений и добавление их в кэш
func (s *Service) Receiving(m *stan.Msg) {

	var data map[string]interface{}
	log.Println("Полученное сообщение:\n", string(m.Data))
	result := s.Validation(&m.Data)
	if result == true {
		json.Unmarshal(m.Data, &data)
		order_uid := data["order_uid"].(string)
		delete(data, "order_uid")
		json_data, _ := json.Marshal(data)
		s.Save_data(&order_uid, &json_data)
	} else {
		log.Println("В результате получения данных произошла ошибка")
	}
}

// отписка от темы и закрытие соединени с nats streamig server
func (s *Service) close_connect(signalChan *chan os.Signal, sub *stan.Subscription, sc *stan.Conn) {
	for range *signalChan {
		fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
		(*sub).Unsubscribe()
		(*sc).Close()
	}
}

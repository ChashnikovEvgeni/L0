package server

import (
	"L0/models"
	"L0/service"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

type Httpserver struct {
	Port    string
	Service service.Service
}

func (h *Httpserver) Set_config(port *string, service *service.Service) {
	if *port == "" {
		h.Port = ":8082"
	} else {
		h.Port = *port
	}
	h.Service = *service
}

// Реализация маршрутизации для http сервера
func (h *Httpserver) Server_work() {
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/showone/", h.Showone)
	log.Fatal(http.ListenAndServe(h.Port, nil))
}

// Обработчик для вывода uid всех заказов доступных в базе
func (h *Httpserver) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./server/templates/show_all.html",
		"./server/templates/base.html",
	}

	items := h.Service.Cache.Items()

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Httpserver Error", 500)
		return
	}

	err = ts.Execute(w, items)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Httpserver Error", 500)
	}
}

// Обработчик вывода подрбной информации о конкретном заказе
func (h *Httpserver) Showone(w http.ResponseWriter, r *http.Request) {

	var formData struct {
		Order_uid string
		Data      string
	}

	result, found := h.Service.Cache.Get(r.URL.Path[len("/showone/"):])
	order := result.(models.Order)
	if found {
		formData.Order_uid = order.Order_uid
		formData.Data = string(order.Mjson)

	}

	if r.URL.Path[:len("/showone/")] != "/showone/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./server/templates/show_one.html",
		"./server/templates/base.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Httpserver Error", 500)
		return
	}

	err = ts.Execute(w, formData)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Httpserver Error", 500)
	}

}

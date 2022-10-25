package main

import (
	"L0/models"
	"L0/service"
	"database/sql"
	"testing"
)

func TestValidation(t *testing.T) {
	service := service.Service{}
	service.Set_config("", "", "")

	t.Run("correct_data", func(t *testing.T) {
		data := []byte("{\n\t\t\"order_uid\": \"gfhjghy23\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 13489573945,\n\t\t  \"goods_total\": 366,\n\t\t  \"custom_fee\": 28\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
		result := true

		realresult := service.Validation(&data)
		if realresult != result {
			t.Errorf("Ожидался %t, получен: %t", result, realresult)

		}
	})

	t.Run("int_instead_string", func(t *testing.T) {
		//вместо int в amount написан string
		data := []byte("{\n\t\t\"order_uid\": \"gfhjghy23\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": \"fghjfgh\",\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 13489573945,\n\t\t  \"goods_total\": 366,\n\t\t  \"custom_fee\": 28\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
		result := false

		realresult := service.Validation(&data)
		if realresult != result {
			t.Errorf("Ожидался %t, получен: %t", result, realresult)

		}
	})

	t.Run("int_instead_string", func(t *testing.T) {
		//вместо string в track_number написан int
		data := []byte("{\n\t\t\"order_uid\": \"gfhjghy23\",\n\t\t\"track_number\": 10,\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 13489573945,\n\t\t  \"goods_total\": 366,\n\t\t  \"custom_fee\": 28\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
		result := false

		realresult := service.Validation(&data)
		if realresult != result {
			t.Errorf("Ожидался %t, получен: %t", result, realresult)

		}
	})

	t.Run("order_len_more_than_max", func(t *testing.T) {
		//вместо string в track_number написан int
		data := []byte("{\n\t\t\"order_uid\": \"dfgjdfghjfghjfghjfghjfghjfghj233\",\n\t\t\"track_number\": 10,\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 13489573945,\n\t\t  \"goods_total\": 366,\n\t\t  \"custom_fee\": 28\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
		result := false

		realresult := service.Validation(&data)
		if realresult != result {
			t.Errorf("Ожидался %t, получен: %t", result, realresult)

		}
	})

	t.Run("order_len_less_than_max", func(t *testing.T) {
		//вместо string в track_number написан int
		data := []byte("{\n\t\t\"order_uid\": \"\",\n\t\t\"track_number\": 10,\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 13489573945,\n\t\t  \"goods_total\": 366,\n\t\t  \"custom_fee\": 28\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
		result := false

		realresult := service.Validation(&data)
		if realresult != result {
			t.Errorf("Ожидался %t, получен: %t", result, realresult)

		}
	})

}

func TestDB(t *testing.T) {

	service := service.Service{}
	service.Set_config("", "", "")
	t.Run("db_connect", func(t *testing.T) {
		//проверка соединения с БД
		var db *sql.DB
		db = service.DB_connect()
		err := db.Ping()
		if err != nil {
			t.Errorf("Подключение к базе данных некорректно")
		}
	})

}

func TestCash(t *testing.T) {
	service := service.Service{}
	service.Set_config("", "", "")

	t.Run("add_in_cash", func(t *testing.T) {
		// проверка правильности сохранения данных в кэш
		order_uid := "fghjfghjgh25"
		mjson := []byte("{\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 13489573945,\n\t\t  \"goods_total\": 366,\n\t\t  \"custom_fee\": 28\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
		standart_order := models.Order{Order_uid: order_uid, Mjson: mjson}
		service.Add_in_cache(&order_uid, &mjson)
		result, found := service.Cache.Get(order_uid)
		resultorder := result.(models.Order)
		if found == false {
			t.Errorf("Заказ не найден в кэше")
		}
		if resultorder.Order_uid != standart_order.Order_uid && string(resultorder.Mjson) != string(standart_order.Mjson) {
			t.Errorf("Данные полученные из кэша отличаются от эталонных")
		}
	})
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zykunov/courseGoFirst/WB0/storage"
	"github.com/zykunov/courseGoFirst/WB0/third_party/nats/source"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	fmt.Println("Listen to your heart")

	nc.Subscribe("intros", func(m *nats.Msg) {

		var order source.SummaryModel

		//Проверка на валидность получаемых данных
		err := json.Unmarshal([]byte(m.Data), &order)
		if err != nil {
			fmt.Println("Не верные данные:", err)
		}
		// fmt.Println(order)
		insert2DB(order, m.Data)

	})

	time.Sleep(1 * time.Hour)
}

func insert2DB(data source.SummaryModel, stringData []byte) {

	DBconnection := storage.GetDB()

	insertStmt := fmt.Sprintf(`insert into orders ("order_uid", "data") values('%s', '%s')`, data.OrderUid, stringData)
	_, e := DBconnection.Exec(insertStmt)
	storage.CheckError(e)
}

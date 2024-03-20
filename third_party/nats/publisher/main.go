package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/zykunov/courseGoFirst/WB0/third_party/nats/source"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	msg := ReadFile()

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	c.Publish("intros", msg) //json

	log.Printf("sent message ")
	time.Sleep(1 * time.Second)

}

func ReadFile() (orderData source.SummaryModel) {

	jsonFile, err := os.Open("./third_party/nats/source/model3.json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened model.json")
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var orderModel source.SummaryModel

	err = json.Unmarshal(data, &orderModel)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	return orderModel

}

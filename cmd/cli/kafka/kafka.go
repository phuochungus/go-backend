package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaUrl   = "localhost:9092"
	kafkaTopic = "user_topic_vip"
)

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func getKafkaReader(kafkaURL, topic string, groupId string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(
		kafka.ReaderConfig{
			Brokers:        brokers,
			GroupID:        groupId,
			Topic:          topic,
			MinBytes:       10e3,
			MaxBytes:       10e6,
			CommitInterval: time.Second,
			// StartOffset:    kafka.LastOffset,
			StartOffset: kafka.FirstOffset,
		})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg string, msgType string) *StockInfo {
	return &StockInfo{Message: msg, Type: msgType}
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))

	body := make(map[string]interface{})

	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)

	if err != nil {
		fmt.Println("failed to write messages:", err)
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"info":   s,
			"msg":    "failed to send message",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"info":   s,
		"msg":    "send message success",
	})

}

func RegisterConsumerATC(id int) {
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id)

	reader := getKafkaReader(kafkaUrl, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer(%d) Hong Phien ATC::\n", id)

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) Hong Phien ATC Error: %v\n", id, err)
			continue
		}

		fmt.Printf("Consumer(%d), hong topic: %v, partition: %v, offset: %v, time: %v, message: %v \n", id, message.Topic, message.Partition, message.Offset, message.Time, string(message.Value))
	}

}

func main() {

	kafkaProducer = getKafkaWriter(kafkaUrl, kafkaTopic)
	defer kafkaProducer.Close()

	r := gin.Default()

	r.POST("/action", actionStock)

	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)
	go RegisterConsumerATC(4)
	go RegisterConsumerATC(5)

	r.Run(":9999")

}

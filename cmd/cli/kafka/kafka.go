package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaUrl   = "localhost:19092"
	kafkaTopic = "user_topic_vip"
)

// For producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// For consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua bán chứng khoán
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s
	jsonBody, _ := json.Marshal(body)

	// create message
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	// write message by producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)

	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
	})
}

// Consumer hóng mua ATC
func RegisterConsumerATC(id int) {
	kafkaGroupId := "consumer-group-"
	reader := getKafkaReader(kafkaUrl, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("consumer (%d) Hóng Phiên ATC::\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("consumer (%d) error: %v", id, err)
		}

		fmt.Printf("consumer (%d), hóng topic: %v, partition: %v, offset: %v, time: %d %s = %s \n",
			id,
			m.Topic,
			m.Partition,
			m.Offset,
			m.Time.Unix(),
			string(m.Key),
			string(m.Value),
		)
	}
}

func main() {
	r := gin.Default()
	kProducer := getKafkaWriter(kafkaUrl, kafkaTopic)
	defer kProducer.Close()

	r.POST("action/stock", actionStock)

	// đăng ký 2 user để mua Stock trong ATC (1) (2)
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	r.Run(":8999")
}

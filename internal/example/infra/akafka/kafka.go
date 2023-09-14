package akafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}

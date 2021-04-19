package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic = "buff-test-events"
	addr  = "localnor.com:9094"
)

func main() {
	// testWriteMsg()
	testConsumeMsg()
}

func testWriteMsg() {
	// to produce messages
	conn, err := kafka.DialLeader(context.Background(), "tcp", addr, topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	_ = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func testConsumeMsg() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{addr},
		GroupID: "buff-group-id-1",
		Topic:   topic,
		// MinBytes:               10e3, // 10KB
		MinBytes:        1,    // 10KB
		MaxBytes:        10e6, // 10MB
		StartOffset:     -1,
		MaxWait:         time.Second,
		ReadLagInterval: time.Second,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("read msg error: ", err)
			break
		}
		now := time.Now()
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		fmt.Printf("msgTime:%v,curr:%v\n", m.Time, now)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

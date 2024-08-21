package data

import "time"

type AsyncProcessor struct {
}

type AsyncInterface interface {
	Publish(topic string, message []byte) error
	Metadata(server string) error
	Republish(topic string, time time.Time) error
}

//Implement Async Processor Methods

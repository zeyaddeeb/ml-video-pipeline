package main

import (
	"github.com/ml-video-pipeline/src"
)

var (
	conf  src.Config
	mode  string
	topic string
)

func main() {
	producer, _ := src.NewProducer(conf, "go_test", 0)
	_ = producer.Push([]byte("{\"name\":\"test\"}"), nil)

}

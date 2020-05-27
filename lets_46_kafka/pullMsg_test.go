package main

import (
	"testing"
)

func TestKafkaConsumer_Consumption(t *testing.T) {
	tests := []struct {
		name     string
		consumer KafkaConsumer
	}{
		// TODO: Add test cases.
		{name: "success", consumer: KafkaConsumer{Address: []string{"104.215.49.91:9092"}, Topic: "test_log"}},
		{name: "failed", consumer: KafkaConsumer{Address: []string{"000.215.49.91:9092"}, Topic: "test_log"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.consumer.Consumption(t)
		})
	}
}

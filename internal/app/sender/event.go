package sender

import (
	"github.com/Slobodskoy/rcn-service-api/internal/model"
)

type EventSender interface {
	Send(Service *model.ServiceEvent) error
}

type KafkaEventSender struct {
}

func NewKafkaEventSender() EventSender {
	return &KafkaEventSender{}
}

func (sender *KafkaEventSender) Send(Service *model.ServiceEvent) error {

	return nil
}

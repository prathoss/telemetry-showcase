package kafka

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/rides/config"
	"github.com/prathoss/telemetry_showcase/rides/messages"
	"go.opentelemetry.io/otel"
)

func NewProducer(cfg config.Config) (*Producer, error) {
	kCfg := sarama.NewConfig()
	kCfg.Version = sarama.V3_6_0_0
	kCfg.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewAsyncProducer(cfg.KafkaBrokers, kCfg)
	if err != nil {
		return nil, err
	}
	producer = otelsarama.WrapAsyncProducer(kCfg, producer)
	return &Producer{producer: producer}, nil
}

type Producer struct {
	producer sarama.AsyncProducer
}

func (p *Producer) SendRideEnd(ctx context.Context, rideID uuid.UUID) error {
	val, err := json.Marshal(messages.EndRide{RideID: rideID})
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: "rides.end",
		Value: sarama.ByteEncoder(val),
	}
	otel.GetTextMapPropagator().Inject(ctx, otelsarama.NewProducerMessageCarrier(msg))
	p.producer.Input() <- msg
	return nil
}

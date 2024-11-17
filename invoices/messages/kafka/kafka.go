package kafka

import (
	"context"
	"log/slog"
	"time"

	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"github.com/prathoss/telemetry_showcase/invoices/config"
	"github.com/prathoss/telemetry_showcase/shared"
)

func NewConsumer(cfg config.Config) (*Consumer, error) {
	kCfg := sarama.NewConfig()
	kCfg.Version = sarama.V3_6_0_0
	kCfg.Consumer.Offsets.AutoCommit.Enable = true
	kCfg.Consumer.Offsets.AutoCommit.Interval = time.Second

	consumerGroup, err := sarama.NewConsumerGroup(cfg.KafkaBrokers, "invoices", kCfg)
	if err != nil {
		return nil, err
	}

	reHandler, err := NewRideEndHandler(cfg)
	if err != nil {
		return nil, err
	}
	wrappedHandler := otelsarama.WrapConsumerGroupHandler(reHandler)

	return &Consumer{
		consumerGroup:  consumerGroup,
		rideEndHandler: wrappedHandler,
	}, nil
}

type Consumer struct {
	consumerGroup  sarama.ConsumerGroup
	rideEndHandler sarama.ConsumerGroupHandler
}

func (c *Consumer) ConsumeRideEnd(ctx context.Context) {
	slog := slog.With("component", "kafka")
	for {
		select {
		case <-ctx.Done():
			break
		default:
			err := c.consumerGroup.Consume(ctx, []string{"rides.end"}, c.rideEndHandler)
			if err != nil {
				slog.ErrorContext(ctx, "could not consume message", shared.Err(err))
			}
		}
	}
}

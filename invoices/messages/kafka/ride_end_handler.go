package kafka

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"github.com/prathoss/telemetry_showcase/invoices/config"
	"github.com/prathoss/telemetry_showcase/invoices/messages"
	"github.com/prathoss/telemetry_showcase/proto/rides"
	"github.com/prathoss/telemetry_showcase/shared"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ sarama.ConsumerGroupHandler = (*RideEndHandler)(nil)

func NewRideEndHandler(cfg config.Config) (*RideEndHandler, error) {
	ridesConn, err := grpc.NewClient(
		cfg.RidesAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}
	ridesClient := rides.NewRidesClient(ridesConn)
	return &RideEndHandler{
		ridesClient: ridesClient,
		tracer:      otel.GetTracerProvider().Tracer("invoice.end_ride"),
	}, nil
}

type RideEndHandler struct {
	ridesClient rides.RidesClient
	tracer      trace.Tracer
}

func (r *RideEndHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (r *RideEndHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (r *RideEndHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		// extract context with span information from message
		linkableContext := otel.GetTextMapPropagator().Extract(context.Background(), otelsarama.NewConsumerMessageCarrier(message))

		// create context with timeout for message handling
		ctx, cFunc := context.WithTimeout(context.Background(), 10*time.Second)
		ctx, span := r.tracer.Start(ctx, "consume.ride_end")
		// link span from message
		span.AddLink(trace.LinkFromContext(linkableContext))
		er := messages.EndRide{}
		if err := json.Unmarshal(message.Value, &er); err != nil {
			slog.ErrorContext(ctx, "could not unmarshal message", shared.Err(err))
			span.RecordError(err)
			span.End()
			continue
		}
		slog.InfoContext(ctx, "received ride end", "ride_id", er.RideID)
		span.SetAttributes(attribute.String("ride_id", er.RideID.String()))

		ride, err := r.ridesClient.GetRide(ctx, &rides.GetRideRequest{RideId: er.RideID.String()})
		if err != nil {
			slog.ErrorContext(ctx, "could not get ride", shared.Err(err))
			span.RecordError(err)
			span.End()
			continue
		}

		ctx, invoiceSpan := r.tracer.Start(ctx, "generate invoice")
		slog.InfoContext(ctx, "generating invoice", "ride", ride)
		// hardcore processing task
		time.Sleep(2 * time.Second)
		slog.InfoContext(ctx, "invoice generated", "ride", ride)
		invoiceSpan.End()

		slog.InfoContext(ctx, "ride end processed", "ride", ride)
		session.MarkMessage(message, "")
		span.End()
		cFunc()
	}
	return nil
}

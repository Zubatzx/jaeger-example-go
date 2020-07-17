package showtime

import (
	"context"

	jaegerLog "showtime/pkg/log"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// Data ...
type Data interface {
	GetShowtime(ctx context.Context, movieID string) (string, error)
}

// Service ...
type Service struct {
	data   Data
	tracer opentracing.Tracer
	logger jaegerLog.Factory
}

// New ...
func New(data Data, tracer opentracing.Tracer, logger jaegerLog.Factory) Service {
	return Service{
		data:   data,
		tracer: tracer,
		logger: logger,
	}
}

// GetShowtime ...
func (s Service) GetShowtime(ctx context.Context, movieID string) (string, error) {
	var (
		showtime string
		err      error
	)

	// Check if have span on context
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := s.tracer.StartSpan("GetShowtime", opentracing.ChildOf(span.Context()))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	s.logger.For(ctx).Info("Get showtime", zap.String("movieID", movieID))
	showtime, err = s.data.GetShowtime(ctx, movieID)
	if err != nil {
		s.logger.For(ctx).Error("Failed to get showtime", zap.Error(err))
	}
	s.logger.For(ctx).Info("Showtime fetched", zap.String("showtime", showtime))

	return showtime, err
}

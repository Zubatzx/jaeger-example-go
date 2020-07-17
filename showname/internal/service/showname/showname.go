package showname

import (
	"context"

	jaegerLog "showname/pkg/log"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// Data ...
type Data interface {
	GetShowname(ctx context.Context, movieID string) (string, error)
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

// GetShowname ...
func (s Service) GetShowname(ctx context.Context, movieID string) (string, error) {
	var (
		showname string
		err      error
	)

	// Check if have span on context
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := s.tracer.StartSpan("GetShowname", opentracing.ChildOf(span.Context()))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	s.logger.For(ctx).Info("Get showname", zap.String("movieID", movieID))
	showname, err = s.data.GetShowname(ctx, movieID)
	if err != nil {
		s.logger.For(ctx).Error("Failed to get showname", zap.Error(err))
	}
	s.logger.For(ctx).Info("Showname fetched", zap.String("showname", showname))

	return showname, err
}

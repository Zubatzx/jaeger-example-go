package book

import (
	"context"

	jaegerLog "book/pkg/log"

	bookEntity "book/internal/entity/book"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// Data ...
type Data interface {
	GetShowname(ctx context.Context, movieID string) (string, error)
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

// GetShowDetail ...
func (s Service) GetShowDetail(ctx context.Context, movieID string) (bookEntity.BookDetail, error) {
	var (
		bookDetail bookEntity.BookDetail
		showname   string
		showtime   string
		err        error
	)

	// Check if have span on context
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := s.tracer.StartSpan("GetShowDetail", opentracing.ChildOf(span.Context()))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	s.logger.For(ctx).Info("Get showname", zap.String("movieID", movieID))
	showname, err = s.data.GetShowname(ctx, movieID)
	if err != nil {
		s.logger.For(ctx).Error("Failed to get showname", zap.Error(err))
		return bookDetail, err
	}
	s.logger.For(ctx).Info("Showname fetched", zap.String("showname", showname))

	s.logger.For(ctx).Info("Get showtime", zap.String("movieID", movieID))
	showtime, err = s.data.GetShowtime(ctx, movieID)
	if err != nil {
		s.logger.For(ctx).Error("Failed to get showtime", zap.Error(err))
		return bookDetail, err
	}
	s.logger.For(ctx).Info("Showtime fetched", zap.String("showtime", showtime))

	bookDetail.Showname = showname
	bookDetail.Showtime = showtime

	return bookDetail, err
}

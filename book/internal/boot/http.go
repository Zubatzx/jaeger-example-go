package boot

import (
	"book/pkg/httpclient"
	"book/pkg/tracing"
	"log"

	"net/http"

	"book/internal/config"

	jaegerLog "book/pkg/log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	bookData "book/internal/data/book"
	bookServer "book/internal/delivery/http"
	bookHandler "book/internal/delivery/http/book"
	bookService "book/internal/service/book"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		s     bookServer.Server    // HTTP Server Object
		bd    bookData.Data        // Domain data layer
		bs    bookService.Service  // Domain service layer
		bh    *bookHandler.Handler // Domain handler
		cfg   *config.Config       // Configuration object
		httpc *httpclient.Client

		logger *zap.Logger
	)

	// Initialize Config
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg = config.Get()

	// Open MySQL DB Connection
	// db, err := sqlx.Open("mysql", cfg.Database.Master)
	// if err != nil {
	// 	log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	// }

	// Set logger used for jaeger
	logger, _ = zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	zapLogger := logger.With(zap.String("service", "book"))
	zlogger := jaegerLog.NewFactory(zapLogger)

	// Set tracer for service
	tracer, closer := tracing.Init("book", zlogger)
	defer closer.Close()

	// Open HTTPClient
	httpc = httpclient.NewClient(tracer)

	bd = bookData.New(httpc)
	bs = bookService.New(bd, tracer, zlogger)
	bh = bookHandler.New(bs, tracer, zlogger)

	s = bookServer.Server{
		Book: bh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}

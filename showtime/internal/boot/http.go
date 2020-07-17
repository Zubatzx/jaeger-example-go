package boot

import (
	"log"
	"showtime/pkg/tracing"

	"net/http"

	"showtime/internal/config"

	jaegerLog "showtime/pkg/log"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	showtimeData "showtime/internal/data/showtime"
	showtimeServer "showtime/internal/delivery/http"
	showtimeHandler "showtime/internal/delivery/http/showtime"
	showtimeService "showtime/internal/service/showtime"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		s   showtimeServer.Server    // HTTP Server Object
		sd  showtimeData.Data        // Domain data layer
		ss  showtimeService.Service  // Domain service layer
		sh  *showtimeHandler.Handler // Domain handler
		cfg *config.Config           // Configuration object
		// httpc *httpclient.Client

		logger *zap.Logger
	)

	// Initialize Config
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg = config.Get()

	// Open MySQL DB Connection
	db, err := sqlx.Open("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	// Set logger used for jaeger
	logger, _ = zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	zapLogger := logger.With(zap.String("service", "showtime"))
	zlogger := jaegerLog.NewFactory(zapLogger)

	// Set tracer for service
	tracer, closer := tracing.Init("showtime", zlogger)
	defer closer.Close()

	// Open HTTPClient
	// httpc = httpclient.NewClient(tracer)

	sd = showtimeData.New(db, zlogger)
	ss = showtimeService.New(sd, tracer, zlogger)
	sh = showtimeHandler.New(ss, tracer, zlogger)

	s = showtimeServer.Server{
		Showtime: sh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}

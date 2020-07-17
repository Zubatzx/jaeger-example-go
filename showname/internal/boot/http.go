package boot

import (
	"log"
	"showname/pkg/tracing"

	"net/http"

	"showname/internal/config"

	jaegerLog "showname/pkg/log"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	shownameData "showname/internal/data/showname"
	shownameServer "showname/internal/delivery/http"
	shownameHandler "showname/internal/delivery/http/showname"
	shownameService "showname/internal/service/showname"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		s   shownameServer.Server    // HTTP Server Object
		sd  shownameData.Data        // Domain data layer
		ss  shownameService.Service  // Domain service layer
		sh  *shownameHandler.Handler // Domain handler
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
	zapLogger := logger.With(zap.String("service", "showname"))
	zlogger := jaegerLog.NewFactory(zapLogger)

	// Set tracer for service
	tracer, closer := tracing.Init("showname", zlogger)
	defer closer.Close()

	// Open HTTPClient
	// httpc = httpclient.NewClient(tracer)

	sd = shownameData.New(db, zlogger)
	ss = shownameService.New(sd, tracer, zlogger)
	sh = shownameHandler.New(ss, tracer, zlogger)

	s = shownameServer.Server{
		Showname: sh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}

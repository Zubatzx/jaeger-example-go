package showtime

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	jaegerLog "showtime/pkg/log"
	"showtime/pkg/tracing"
)

type (
	// Data ...
	Data struct {
		mysql     *sqlx.DB
		stmtMySQL map[string]*sqlx.Stmt

		tracer opentracing.Tracer
		logger jaegerLog.Factory
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

const (
// MySQL Statement
// updateStockSuccess  = "UpdateStockSuccess"
// qUpdateStockSuccess = `INSERT INTO reporting_crawling.download_stock_ecommerce (OutletID, EcommerceID, EcommerceProductID, ProductID, Stock, Price, Stock_OldDown, DownloadDate) VALUES (?, 0, '000', ?, ?, 0, 0, NOW())ON DUPLICATE KEY UPDATE Stock = VALUES(Stock), Price = VALUES(Price), Stock_OldDown = VALUES(Stock_OldDown), DownloadDate = NOW()`
)

var (
	readStmtMySQL = []statement{
		// {updateStockSuccess, qUpdateStockSuccess},
	}
)

// New ...
func New(mysql *sqlx.DB, logger jaegerLog.Factory) Data {
	d := Data{
		mysql:  mysql,
		logger: logger,
	}

	d.tracer, _ = tracing.Init("mysql", logger)
	d.initStmt()

	return d
}

func (d *Data) initStmt() {
	var (
		err        error
		stmtsMySQL = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmtMySQL {
		stmtsMySQL[v.key], err = d.mysql.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}

	d.stmtMySQL = stmtsMySQL
}

// GetShowtime ...
func (d Data) GetShowtime(ctx context.Context, movieID string) (string, error) {
	var (
		showtime string
		err      error
	)

	// Check if have span on context
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
		span.SetTag("mysql.server", "123.72.156.4")
		span.SetTag("mysql.database", "movie")
		span.SetTag("mysql.table", "showtime")
		span.SetTag("mysql.query", "SELECT * FROM movie.showtime WHERE movie_id="+movieID)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	// assumed data fetched from database
	showtime = "11.00 PM"
	d.logger.For(ctx).Info("SQL Query Success", zap.String("showtime", showtime))

	// case if err
	// if err != nil {
	// 	d.logger.For(ctx).Error("SQL Query Failed", zap.Error(err))
	// 	return showtime, err
	// }

	return showtime, err
}

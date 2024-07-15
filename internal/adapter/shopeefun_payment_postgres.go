package adapter

import (
	// "log"

	"payment-service/internal/infrastructure"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func WithShopeefunPaymentPostgres() Option {
	return func(a *Adapter) {
		dbUser := infrastructure.Envs.ShopeefunPaymentPostgres.Username
		dbPassword := infrastructure.Envs.ShopeefunPaymentPostgres.Password
		dbName := infrastructure.Envs.ShopeefunPaymentPostgres.Database
		dbHost := infrastructure.Envs.ShopeefunPaymentPostgres.Host
		dbSSLMode := infrastructure.Envs.ShopeefunPaymentPostgres.SslMode
		dbPort := infrastructure.Envs.ShopeefunPaymentPostgres.Port

		dbMaxPoolSize := infrastructure.Envs.DB.MaxOpenCons
		dbMaxIdleConns := infrastructure.Envs.DB.MaxIdleCons
		dbConnMaxLifetime := infrastructure.Envs.DB.ConnMaxLifetime

		connectionString := "user=" + dbUser + " password=" + dbPassword + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=" + dbSSLMode + " TimeZone=UTC"
		db, err := sqlx.Connect("postgres", connectionString)
		if err != nil {
			log.Fatal().Err(err).Msg("Error connecting to Postgres")
		}

		db.SetMaxOpenConns(dbMaxPoolSize)
		db.SetMaxIdleConns(dbMaxIdleConns)
		db.SetConnMaxLifetime(time.Duration(dbConnMaxLifetime) * time.Second)

		// check connection
		err = db.Ping()
		if err != nil {
			log.Fatal().Err(err).Msg("Error connecting to Shopeefun Payment Postgres")
		}

		a.ShopeefunPaymentPostgres = db
		log.Info().Msg("Shopeefun Payment Postgres connected")
	}
}

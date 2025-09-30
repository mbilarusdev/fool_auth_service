package infrasturcture

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	. "github.com/mbilarusdev/fool_auth_service/internal/logger"
	"github.com/mbilarusdev/fool_auth_service/internal/utils"
)

func ConnectPostgres() *utils.DBPool {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	maxConns, err := strconv.Atoi(os.Getenv("AUTH_SERVICE_DB_CONNS"))

	if err != nil {
		PanicErrWithMsg(err, "Failed to parse maxConns")
	}

	connString := fmt.Sprintf(
		"host=%v port=%v dbname=%v user=%v password=%v pool_max_conns=%v pool_max_conn_lifetime=5m sslmode=disable",
		host,
		port,
		db,
		user,
		password,
		maxConns,
	)

	conf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		PanicErrWithMsg(err, "Failed to parse postgres config!")
	}

	outerPool := new(utils.DBPool)
	outerPool.Limiter = make(chan struct{}, maxConns)

	ctx := context.Background()

	innerPool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		PanicErrWithMsg(err, "Failed to connect to postgres")
	}

	outerPool.InnerPool = innerPool

	LogInfo("Postgres connected!")

	c, err := outerPool.Acquire(ctx)
	if err != nil {
		PanicErrWithMsg(err, "Failed to acquire conn when SET TIME ZONE UTC")
	}

	defer outerPool.Release(ctx, c)

	_, err = c.Exec(ctx, "SET TIME ZONE 'UTC'")
	if err != nil {
		PanicErrWithMsg(err, "Failed to SET TIME ZONE UTC")
	}
	LogInfo("TIME ZONE UTS Set upped in Postgres!")

	return outerPool
}

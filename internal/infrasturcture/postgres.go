package infrasturcture

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	. "github.com/mbilarusdev/fool_auth_service/internal/logger"
	"github.com/mbilarusdev/fool_auth_service/internal/utils"
)

func ConnectPostgres() *utils.DBPool {
	db := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	maxConns, err := strconv.Atoi(os.Getenv("POSTGRES_CONNS_AUTH_SERVICE"))
	if err != nil {
		PanicErrWithMsg(err, "Failed to parse maxConns")
	}

	connString := fmt.Sprintf(
		"dbname=%v user=%v password=%v pool_max_conns=%v pool_max_conn_lifetime=5m",
		db,
		user,
		password,
		maxConns,
	)

	conf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		PanicErrWithMsg(err, "Failed to parse postgres config!")
	}

	conf.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		LogInfo("Postgres connected!")

		_, err := c.Exec(ctx, "SET TIME ZONE 'UTC'")
		if err != nil {
			return err
		}

		return nil
	}

	outerPool := new(utils.DBPool)
	outerPool.Limiter = make(chan struct{}, maxConns)

	innerPool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		PanicErrWithMsg(err, "Failed to connect to postgres")
	}

	outerPool.InnerPool = innerPool

	return outerPool
}

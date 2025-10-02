package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mbilarusdev/fool_auth_service/internal/models"
	"github.com/mbilarusdev/fool_auth_service/internal/repository/repoerr"
	"github.com/mbilarusdev/fool_base/v2/infra/db"
	"github.com/mbilarusdev/fool_base/v2/log"
	modelid "github.com/mbilarusdev/fool_base/v2/model_id"
	"go.uber.org/zap"
)

type PlayerDataProvider interface {
	Register(
		ctx context.Context,
		username string,
		creds string,
	) error
	Login(
		ctx context.Context,
		username string,
		creds string,
	) error
}

type PlayerRepository struct {
	pool *db.DBPool
}

func (repo *PlayerRepository) Register(
	ctx context.Context,
	username string,
	creds string,
) (*modelid.ModelId, error) {
	conn, err := repo.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer repo.pool.Release(ctx, conn)

	id := new(modelid.ModelId)

	if err := conn.QueryRow(ctx, "INSERT INTO players (username, creds) VALUES ($1, $2) RETURNING id;",
		username,
		creds,
	).Scan(&id); err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok && pgErr.Code == "23505" {
			return nil, log.Err(&repoerr.UniqueUsernameError{Username: username}, "")
		}
		return nil, log.Err(err, "Insert player failed!")
	}

	log.Info("Player inserted: ", zap.String("PlayerID", id.String()))

	return id, nil
}

func (repo *PlayerRepository) Login(
	ctx context.Context,
	username string,
	creds string,
) (*models.Player, error) {
	conn, err := repo.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer repo.pool.Release(ctx, conn)

	player := new(models.Player)

	if err := conn.QueryRow(ctx, "SELECT * FROM players p WHERE username = $1 AND creds = $2 LIMIT 1;",
		username,
		creds,
	).Scan(&player.ID, &player.Username, &player.Creds); err != nil {
		return nil, log.Err(err, "Select player failed!")
	}

	log.Info("Player logined: ", zap.String("PlayerID", player.ID.String()))

	return player, nil
}

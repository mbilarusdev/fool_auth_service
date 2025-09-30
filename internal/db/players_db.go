package source

import (
	"context"

	"github.com/mbilarusdev/fool_auth_service/internal/models"
)

type PlayersProvider interface {
	Register(ctx context.Context, player models.Player) error
	Login(ctx context.Context, player models.Player) error
}

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mbilarusdev/fool_auth_service/internal/repository"
	"github.com/mbilarusdev/fool_auth_service/internal/repository/repoerr"
	"github.com/mbilarusdev/fool_auth_service/internal/request"
	"github.com/mbilarusdev/fool_auth_service/internal/utils"
	"github.com/mbilarusdev/fool_base/v2/infra/network"
	"github.com/mbilarusdev/fool_base/v2/log"
	"go.uber.org/zap"
)

type PlayerResource interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	CheckAuthorized(w http.ResponseWriter, r *http.Request)
}

type PlayerController struct {
	repo *repository.PlayerRepository
}

func (c *PlayerController) Register(w http.ResponseWriter, r *http.Request) {
	op := "PlayerController.Register"
	var errMsg string
	var remoteErr string

	bodyBytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		errMsg = fmt.Sprintf("Error when read body at %v", op)
		remoteErr = "Incorrect request body"
		log.Err(err, errMsg)

		network.WriteError(
			w,
			http.StatusBadRequest,
			remoteErr,
		)

		return
	}

	var regReq request.RegisterRequest

	if err := json.Unmarshal(bodyBytes, &regReq); err != nil {
		errMsg = fmt.Sprintf("Error when unmarshal at %v", op)
		remoteErr = "Incorrect fields in request body"
		log.Err(err, errMsg)

		network.WriteError(
			w,
			http.StatusBadRequest,
			remoteErr,
		)

		return
	}

	username, err := utils.DecryptData(regReq.Username, utils.Conf.Secret)

	if err != nil {
		errMsg = fmt.Sprintf("Error when decrypt data at %v", op)
		remoteErr = "Request must be contains right encrypted register data"
		log.Err(err, errMsg)

		network.WriteError(
			w,
			http.StatusBadRequest,
			remoteErr,
		)

		return
	}

	ctx := context.Background()

	id, err := c.repo.Register(ctx, username, regReq.Creds)

	if err != nil {
		errMsg = fmt.Sprintf("Error when register player at %v", op)
		log.Err(err, errMsg)

		switch err.(type) {
		case *repoerr.UniqueUsernameError:
			remoteErr = "Username already exist"
			network.WriteError(
				w,
				http.StatusConflict,
				remoteErr,
			)
		default:
			remoteErr = "Failed to register player"
			network.WriteError(
				w,
				http.StatusBadRequest,
				remoteErr,
			)
		}

		return
	}

	log.Info("Registered with success: ",
		zap.String("PlayerID", id.String()),
		zap.String("Username", username))
	network.WriteResponse(w, http.StatusNoContent, nil)
}

func (c *PlayerController) Login(w http.ResponseWriter, r *http.Request) {
	op := "PlayerController.Login"
	var errMsg string
	var remoteErr string

	bodyBytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		errMsg = fmt.Sprintf("Error when read body at %v", op)
		remoteErr = "Incorrect request body"
		log.Err(err, errMsg)

		network.WriteError(
			w,
			http.StatusBadRequest,
			remoteErr,
		)

		return
	}

	var logReq request.LoginRequest

	if err := json.Unmarshal(bodyBytes, &logReq); err != nil {
		errMsg = fmt.Sprintf("Error when unmarshal at %v", op)
		remoteErr = "Incorrect fields in request body"
		log.Err(err, errMsg)

		network.WriteError(
			w,
			http.StatusBadRequest,
			remoteErr,
		)

		return
	}

	username, err := utils.DecryptData(logReq.Username, utils.Conf.Secret)

	if err != nil {
		errMsg = fmt.Sprintf("Error when decrypt data at %v", op)
		remoteErr = "Request must be contains right encrypted register data"
		log.Err(err, errMsg)

		network.WriteError(
			w,
			http.StatusBadRequest,
			remoteErr,
		)

		return
	}

	ctx := context.Background()

	player, err := c.repo.Login(ctx, username, logReq.Creds)

	if err != nil {
		errMsg = fmt.Sprintf("Error when login player at %v", op)
		remoteErr = "Failed to login player"
		log.Err(err, errMsg)

		network.WriteError(
			w,
			http.StatusUnauthorized,
			remoteErr,
		)
	}

	token, err := utils.CreateJWT(player.ID.String(), player.Username)

	if err != nil {
		errMsg = fmt.Sprintf("Error when create jwt for player at %v", op)
		remoteErr = "Failed when try to create JWT"
		log.Err(err, errMsg)
		network.WriteError(
			w,
			http.StatusBadRequest,
			remoteErr,
		)
	}

	log.Info("Logined with success: ",
		zap.String("PlayerID", player.ID.String()),
		zap.String("Username", username))
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %v", token))
	network.WriteResponse(w, http.StatusOK, player)
}

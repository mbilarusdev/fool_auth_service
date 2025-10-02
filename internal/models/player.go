package models

import (
	modelid "github.com/mbilarusdev/fool_base/v2/model_id"
)

type Player struct {
	ID       modelid.ModelId `json:"id"`
	Username string          `json:"username"`
	Creds    string          `json:"creds"`
}

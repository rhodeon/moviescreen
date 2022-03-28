package internal

import (
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/domain/repository"
)

type Application struct {
	Config       common.Config
	Repositories repository.Repositories
}

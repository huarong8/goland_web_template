package service

import (
	"gitlab.feedtoken.tech/xdance_group/xdance/internal/repository"
	"gitlab.feedtoken.tech/xdance_group/xdance/pkg/logger"
)

type ISignService interface {
	Save()
}

type SignService struct {
	usersRepos *repository.User
}

func NewSignService(usersRepos *repository.User) ISignService {
	return &SignService{
		usersRepos: usersRepos,
	}
}

func (s *SignService) Save() {
	logger.Info("execute save....")
	r := s.usersRepos.FindById()
	logger.Infow(r)

}

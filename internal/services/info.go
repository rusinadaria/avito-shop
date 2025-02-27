package services

import (
	"avito-shop/internal/repository"
	"avito-shop/models"
)

type InfoService struct {
	repo repository.Info
}

func NewInfoService(repo repository.Info) *InfoService {
	return &InfoService{repo: repo}
}

func (n *InfoService) UserInfo(userId int) (models.InfoResponse, error) {
	info, err := n.repo.GetUserInfo(userId)
	if err != nil {
		return models.InfoResponse{}, err
	}
	return info, nil

}
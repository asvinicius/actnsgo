package info

import (
	"github.com/asvinicius/actnsgo/internal/client"
	"github.com/asvinicius/actnsgo/internal/response"
)

type StatusService struct {
	statusClient *client.Client
}

func NewStatusService(statutClient *client.Client) *StatusService {
	return &StatusService{
		statusClient: statutClient,
	}
}

func (s *StatusService) GetStatus() (*response.StatusResponse, error) {
	return s.statusClient.GetStatus()
}

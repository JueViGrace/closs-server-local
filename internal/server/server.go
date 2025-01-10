package server

import (
	"github.com/JueViGrace/clo-backend/internal/api"
)

type CloServer struct {
	api api.Api
}

func New() *CloServer {
	return &CloServer{
		api: api.New(),
	}
}

func (s *CloServer) Init() (err error) {
	err = s.api.Init()
	if err != nil {
		return err
	}
	return
}

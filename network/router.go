package network

import (
	"github.com/gin-gonic/gin"
	"grpc-go-study/config"
	"grpc-go-study/service"
)

type Network struct {
	cfg     *config.Config
	service *service.Service
	engin   *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	r := &Network{cfg: cfg, service: service, engin: gin.New()}

	return r, nil
}

func (n *Network) StartServer() {
	n.engin.Run(":8080")
}

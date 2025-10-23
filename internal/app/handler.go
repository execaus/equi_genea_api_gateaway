package app

import (
	"equi_genea_api_gateaway/config"
	accountpb "equi_genea_api_gateaway/internal/pb/api/account"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	account accountpb.AccountServiceClient
}

func NewHandler(cfg *config.ServicesConfig) (*Handler, error) {
	h := Handler{}

	if err := h.connectServices(cfg); err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	fmt.Println("Connected to gRPC server successfully!")

	return &h, nil
}

func (h *Handler) connectServices(cfg *config.ServicesConfig) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Account.Host, cfg.Account.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	accountClient := accountpb.NewAccountServiceClient(conn)

	h.account = accountClient

	return nil
}

func (h *Handler) GetRouter() *gin.Engine {
	router := gin.Default()

	return router
}

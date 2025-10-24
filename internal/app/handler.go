package app

import (
	"equi_genea_api_gateaway/config"
	accountpb "equi_genea_api_gateaway/internal/pb/api/account"
	authpb "equi_genea_api_gateaway/internal/pb/api/auth"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	services Services
}

type Services struct {
	Account accountpb.AccountServiceClient
	Auth    authpb.AuthServiceClient
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

	conn, err = grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Auth.Host, cfg.Auth.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	authClient := authpb.NewAuthServiceClient(conn)

	h.services = Services{
		Account: accountClient,
		Auth:    authClient,
	}

	return nil
}

func (h *Handler) GetRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
	}

	return router
}

package app

import (
	"equi_genea_api_gateaway/config"
	accountpb "equi_genea_api_gateaway/internal/pb/api/account"
	authpb "equi_genea_api_gateaway/internal/pb/api/auth"
	herdpb "equi_genea_api_gateaway/internal/pb/api/herd"
	horsepb "equi_genea_api_gateaway/internal/pb/api/horse"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
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
	Herd    herdpb.HerdServiceClient
	Horse   horsepb.HorseServiceClient
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

	conn, err = grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Herd.Host, cfg.Herd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	herdClient := herdpb.NewHerdServiceClient(conn)

	conn, err = grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Horse.Host, cfg.Horse.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	horseClient := horsepb.NewHorseServiceClient(conn)

	h.services = Services{
		Account: accountClient,
		Auth:    authClient,
		Herd:    herdClient,
		Horse:   horseClient,
	}

	return nil
}

func (h *Handler) GetRouter(serverConfig *config.ServerConfig) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{serverConfig.AllowOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	herd := api.Group("/herd", h.authMiddleware)
	{
		herd.POST("", h.createHerd)
		herd.GET("", h.getHerdList)
		herd.GET("/:id", h.getHerdByID)
	}

	horseGender := api.Group("/horse-gender")
	{
		horseGender.GET("", h.getHorseGenderList)
	}

	horseColor := api.Group("/horse-color")
	{
		horseColor.GET("", h.getHorseColorList)
	}

	return router
}

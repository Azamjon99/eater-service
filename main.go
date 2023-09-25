package main

import (
	"context"
	"google.golang.org/grpc"
	pb "eater-service/src/application/protos/eater"
	appsvc "eater-service/src/application/services"
	eatersvc "eater-service/src/domain/eater/services"
	"eater-service/src/infrastructure/config"
	"eater-service/src/infrastructure/jwt"
	eaterrepo "eater-service/src/infrastructure/repositories/eater"
	"eater-service/src/infrastructure/sms"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	grpcserver "eater-service/src/application/grpc"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config,err := config.Load()
	if err != nil{
		panic(err)
	}

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?connect_timeout=%d&sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresDatabase,
		60,
	)
	logger, err := config.NewLogger()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	db,err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	smsClient := sms.NewClient(config.SmsProvideApiKey)
	tokenSvc := jwt.NewService(config.JWTSecret, config.JWTExpiresInSec)

	if err != nil {
		panic(err)
	}
	eaterRepo := eaterrepo.NewRepository(db)

	eaterSvc := eatersvc.NewEaterService(eaterRepo,smsClient,logger)
	eaterApp := appsvc.NewEaterApplicationService(eaterSvc,tokenSvc)
	root := gin.Default()

	root.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowCredentials: true,
	}))
	ctx,cancel := context.WithCancel(context.Background())
	g,ctx := errgroup.WithContext(ctx)

	osSignals := make(chan os.Signal,1)
	
	signal.Notify(osSignals,os.Interrupt,syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(osSignals)

	var httpServer *http.Server
	g.Go(func() error{
		httpServer = &http.Server{
			Addr: config.HttpPort,
			Handler: root,
		}

		logger.Debug("main: started http server", zap.String("port",config.HttpPort))
		if err := httpServer.ListenAndServe();err != http.ErrServerClosed{
			return err
		}
		return nil
	})

	var grpcServer *grpc.Server
	cancel()

	g.Go(func () error {
		server := grpcserver.NewServer(
			eaterApp,
		)
		grpcServer = grpc.NewServer(eaterApp)
		pb.RegisterEaterServiceServer(grpcServer,server)

		lis, err := net.Listen("tcp", config.GrpcPort)
		if err != nil {
			logger.Fatal("main: net.Listen", zap.Error(err))
		}

		defer lis.Close()
	
		logger.Debug("main: started grps server", zap.String("port",config.GrpcPort))

		return grpcServer.Serve(lis)
	})

}
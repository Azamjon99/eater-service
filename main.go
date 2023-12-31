package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpcserver "github.com/Azamjon99/eater-service/src/application/grpc"
	integrationevents "github.com/Azamjon99/eater-service/src/application/integration_events"
	pb "github.com/Azamjon99/eater-service/src/application/protos/eater"
	appsvc "github.com/Azamjon99/eater-service/src/application/services"
	addresssvc "github.com/Azamjon99/eater-service/src/domain/address/services"
	eatersvc "github.com/Azamjon99/eater-service/src/domain/eater/services"
	orderysvc "github.com/Azamjon99/eater-service/src/domain/order/services"
	ratingsvc "github.com/Azamjon99/eater-service/src/domain/rating/services"
	walletsvc "github.com/Azamjon99/eater-service/src/domain/wallet/services"
	"github.com/Azamjon99/eater-service/src/infrastructure/config"
	"github.com/Azamjon99/eater-service/src/infrastructure/jwt"
	addressrepo "github.com/Azamjon99/eater-service/src/infrastructure/repositories/address"
	eaterrepo "github.com/Azamjon99/eater-service/src/infrastructure/repositories/eater"
	orderrepo "github.com/Azamjon99/eater-service/src/infrastructure/repositories/order"
	ratingrepo "github.com/Azamjon99/eater-service/src/infrastructure/repositories/rating"
	walletrepo "github.com/Azamjon99/eater-service/src/infrastructure/repositories/wallet"
	"github.com/Azamjon99/eater-service/src/infrastructure/sms"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger, err := config.NewLogger()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?connect_timeout=%d&sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresDatabase,
		60,
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	producer := integrationevents.NewProducer(ctx, config, logger)
	smsClient := sms.NewClient(config.SmsProvideApiKey)
	tokenSvc := jwt.NewService(config.JWTSecret, config.JWTExpiresInSec)
	eaterRepo := eaterrepo.NewRepository(db)
	addresRepo := addressrepo.NewRepository(db)
	orderrepo := orderrepo.NewRepository(db)
	ratingRepo := ratingrepo.NewRepository(db)
	walletRepo := walletrepo.NewRepository(db)

	eaterSvc := eatersvc.NewEaterService(eaterRepo, smsClient, logger)
	addressSvc := addresssvc.NewAddressService(addresRepo)
	orderSvc := orderysvc.NewOrderService(orderrepo)
	ratingSvc := ratingsvc.NewRatingService(ratingRepo)

	walletSvc := walletsvc.NewWalletService(walletRepo)

	eaterApp := appsvc.NewEaterApplicationService(eaterSvc, tokenSvc)
	addressApp := appsvc.NewAddressApplicationService(addressSvc)
	ratingApp := appsvc.NewRatingApplicationService(ratingSvc, producer,logger)
	orderApp := appsvc.NewOrderApplicationService(orderSvc, producer,logger)
	walletApp := appsvc.NewWalletApplicationService(walletSvc, producer,logger)

	root := gin.Default()

	root.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	g, ctx := errgroup.WithContext(ctx)

	osSignals := make(chan os.Signal, 1)

	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(osSignals)

	var httpServer *http.Server

	g.Go(func() error {
		httpServer = &http.Server{
			Addr:    config.HttpPort,
			Handler: root,
		}

		logger.Debug("main: started http server", zap.String("port", config.HttpPort))
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	var grpcServer *grpc.Server

	g.Go(func() error {
		server := grpcserver.NewServer(
			eaterApp,
			addressApp,
			ratingApp,
			walletApp,
			orderApp,
		)
		grpcServer = grpc.NewServer()
		pb.RegisterEaterServiceServer(grpcServer, server)

		lis, err := net.Listen("tcp", config.GrpcPort)
		if err != nil {
			logger.Fatal("main: net.Listen", zap.Error(err))
		}

		defer lis.Close()

		logger.Debug("main: started grps server", zap.String("port", config.GrpcPort))

		return grpcServer.Serve(lis)
	})

	select {
	case <-osSignals:
		logger.Info("main: received os signal, shutting down")
		break
	case <-ctx.Done():
		break
	}

	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if httpServer != nil {
		httpServer.Shutdown(shutdownCtx)
	}

	if grpcServer != nil {
		grpcServer.GracefulStop()
	}

	if err := g.Wait(); err != nil {
		logger.Error("main: server returned an error", zap.Error(err))
	}
}

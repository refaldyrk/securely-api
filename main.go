package main

import (
	"context"
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"securely-api/config"
	"securely-api/handler"
	"securely-api/middleware"
	"securely-api/repository"
	"securely-api/service"
	"strings"
	"syscall"
	"time"
)

func main() {

	//Set Flag
	typeRunFlag := flag.String("type", "", "dev or docker")
	flag.Parse()

	typeRun := *typeRunFlag
	//Set Time
	startServerTime := time.Now()
	//Set Context
	ctx := context.Background()
	//Init Viper
	if typeRun == "dev" {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	} else if typeRun == "docker" {
		allEnviron := os.Environ()
		for _, v := range allEnviron {
			arrEnv := strings.Split(v, "=")
			viper.Set(strings.ToUpper(arrEnv[0]), arrEnv[1])
		}
	} else {
		panic("no such config")
	}

	//Init Config
	mongodb := config.ConnectMongo(ctx)
	err := mongodb.Ping(1000)
	if err != nil {
		panic(err)
	}

	//Init Database
	DB := mongodb.Database(viper.GetString("DATABASE_NAME"))

	//=================> Repository
	userRepo := repository.NewUserRepository(DB)
	teamRepo := repository.NewTeamRepository(DB)

	//=================> Service
	userService := service.NewUserService(userRepo)
	teamService := service.NewTeamService(teamRepo, userRepo)

	//=================> Handler
	userHandler := handler.NewUserHandler(userService)
	teamHandler := handler.NewTeamHandler(teamService)

	//Server
	app := gin.Default()

	app.Use(gin.Recovery())
	app.Use(gin.Logger())

	// cors	config
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"*"}
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"*"}
	cfg.AllowHeaders = []string{"*"}

	app.Use(cors.New(cfg))

	//======================> Route

	app.POST("/register", userHandler.Register)
	app.POST("/login", userHandler.Login)

	//======================> Myself Endpoint
	myselfEndpoint := app.Group("/api/myself")
	myselfEndpoint.Use(middleware.JWTMiddleware(DB))

	myselfEndpoint.GET("/", userHandler.MySelf)

	//=======================> Team Endpoint
	teamEndpoint := app.Group("/api/team")
	teamEndpoint.Use(middleware.JWTMiddleware(DB))

	teamEndpoint.GET("/", teamHandler.MyTeam)
	teamEndpoint.POST("/create", teamHandler.CreateTeam)
	teamEndpoint.POST("/invite/:team_id", teamHandler.InviteMember)
	teamEndpoint.DELETE("/kick/:team_id", teamHandler.KickMember)

	//Init Server
	srv := &http.Server{
		Addr:    ":" + viper.GetString("PORT"),
		Handler: app,
	}

	// graceful shutdown
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ... ", time.Since(startServerTime).Seconds(), " s")

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()

	log.Println("Server exiting")
}

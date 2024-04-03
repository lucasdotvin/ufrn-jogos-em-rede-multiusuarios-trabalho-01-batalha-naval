package main

import (
	"net/http"
	configuse "trabalho-01-batalha-naval/application/use/config"
	roomuse "trabalho-01-batalha-naval/application/use/room"
	useruse "trabalho-01-batalha-naval/application/use/user"
	"trabalho-01-batalha-naval/config"
	"trabalho-01-batalha-naval/domain/service"
	"trabalho-01-batalha-naval/infrastructure/repository"
	"trabalho-01-batalha-naval/infrastructure/repository/database/sqlite"
	"trabalho-01-batalha-naval/presentation/auth/token"
	"trabalho-01-batalha-naval/presentation/http/broadcast"
	"trabalho-01-batalha-naval/presentation/http/broadcast/websocket"
	"trabalho-01-batalha-naval/presentation/http/controller"
	"trabalho-01-batalha-naval/presentation/http/middleware"
	"trabalho-01-batalha-naval/presentation/http/middleware/auth"
)

func main() {
	cfg := config.GetConfig()

	sqliteDb, err := sqlite.NewDatabase(cfg)
	defer func() {
		_ = sqliteDb.Close()
	}()

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	bcryptHashRepository := repository.NewBcryptHashRepository(cfg)

	jwtTokenService := token.NewJwtService(cfg)

	websocketBroadcaster := websocket.NewWebSocketBroadcaster(cfg)

	roomBroadcast := broadcast.NewRoomBroadcast(websocketBroadcaster)

	sqliteRoomRepository := sqlite.NewSqliteRoomRepository(sqliteDb)
	sqliteRoomMoveRepository := sqlite.NewSqliteRoomMoveRepository(sqliteDb)
	sqliteRoomUserRepository := sqlite.NewSqliteRoomUserRepository(sqliteDb)
	sqliteUserRepository := sqlite.NewSqliteUserRepository(sqliteDb)

	roomService := service.NewRoomService(roomBroadcast, sqliteRoomRepository, sqliteRoomMoveRepository, sqliteRoomUserRepository, sqliteUserRepository)
	userService := service.NewUserService(bcryptHashRepository, sqliteUserRepository)

	signUpUseCase := useruse.NewSignUpUseCase(userService)
	signInUseCase := useruse.NewSignInUseCase(userService)
	findUserUseCase := useruse.NewFindUserUseCase(userService)
	validateUserSubscriptionUseCase := roomuse.NewValidateUserSubscriptionUseCase(roomService)
	handleUserDisconnectUseCase := roomuse.NewHandleUserDisconnectUseCase(roomService)
	listConfigUseCase := configuse.NewListConfigUseCase(cfg)
	listRoomUseCase := roomuse.NewListRoomUseCase(roomService, userService)
	findUserActiveRoomUseCase := roomuse.NewFindUserActiveRoomUseCase(roomService, userService)
	createRoomUseCase := roomuse.NewCreateRoomUseCase(cfg, roomService)
	ingressUserUseCase := roomuse.NewIngressUserUseCase(roomService, userService)
	registerShipPlacementsUseCase := roomuse.NewRegisterShipPlacementsUseCase(roomService)
	registerFireUseCase := roomuse.NewRegisterFireUseCase(roomService)

	authController := controller.NewAuthController(signUpUseCase, signInUseCase, findUserUseCase, jwtTokenService)
	broadcastController := controller.NewBroadcastController(websocketBroadcaster, validateUserSubscriptionUseCase, handleUserDisconnectUseCase)
	configController := controller.NewConfigController(listConfigUseCase)
	roomController := controller.NewRoomController(listRoomUseCase, findUserActiveRoomUseCase, createRoomUseCase, ingressUserUseCase, registerShipPlacementsUseCase, registerFireUseCase)

	cookieTokenAuthMiddleware := auth.NewCookieTokenAuthMiddleware(jwtTokenService)
	webSocketQueryTokenAuthMiddleware := auth.NewWebSocketQueryTokenAuthMiddleware(jwtTokenService)

	router := http.NewServeMux()

	apiRouter := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", apiRouter))

	v1Router := http.NewServeMux()
	apiRouter.Handle("/v1/", http.StripPrefix("/v1", v1Router))

	v1Router.HandleFunc("GET /configs", configController.Index)

	authRouter := http.NewServeMux()
	v1Router.Handle("/auth/", http.StripPrefix("/auth", authRouter))
	authRouter.HandleFunc("POST /sign-up", authController.SignUp)
	authRouter.HandleFunc("POST /sign-in", authController.SignIn)
	authRouter.HandleFunc("POST /refresh", authController.Refresh)
	authRouter.Handle("GET /me", middleware.ApplyToFunc(authController.Me, cookieTokenAuthMiddleware.Handle))

	roomRouter := http.NewServeMux()
	v1Router.Handle("/rooms/", http.StripPrefix("/rooms", middleware.Apply(roomRouter, cookieTokenAuthMiddleware.Handle)))
	roomRouter.HandleFunc("GET /", roomController.Index)
	roomRouter.HandleFunc("GET /my-active-room", roomController.FindMyActiveRoom)
	roomRouter.HandleFunc("POST /", roomController.Store)
	roomRouter.HandleFunc("POST /{room}/ingress/", roomController.IngressUser)
	roomRouter.HandleFunc("POST /{room}/placements/", roomController.RegisterShipPlacements)
	roomRouter.HandleFunc("POST /{room}/fire/", roomController.RegisterFire)

	wsRouter := http.NewServeMux()
	v1Router.Handle("/ws/", http.StripPrefix("/ws", middleware.Apply(wsRouter, webSocketQueryTokenAuthMiddleware.Handle)))
	wsRouter.HandleFunc("/rooms/events", broadcastController.SubscribeForGlobalRoomEvents)
	wsRouter.HandleFunc("/rooms/{room}", broadcastController.SubscribeForRoomEvents)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: router,
	}

	err = server.ListenAndServe()
}

package controller

import (
	"net/http"
	"strings"
	roomuse "trabalho-01-batalha-naval/application/use/room"
	"trabalho-01-batalha-naval/domain/entity"
	"trabalho-01-batalha-naval/presentation/auth"
	"trabalho-01-batalha-naval/presentation/http/broadcast"
)

type BroadcastController struct {
	broadcaster                     broadcast.Broadcaster
	validateUserSubscriptionUseCase *roomuse.ValidateUserSubscriptionUseCase
	handleUserDisconnectUseCase     *roomuse.HandleUserDisconnectUseCase
}

func NewBroadcastController(broadcaster broadcast.Broadcaster, validateUserSubscriptionUseCase *roomuse.ValidateUserSubscriptionUseCase, handleUserDisconnectUseCase *roomuse.HandleUserDisconnectUseCase) *BroadcastController {
	return &BroadcastController{
		broadcaster,
		validateUserSubscriptionUseCase,
		handleUserDisconnectUseCase,
	}
}

func (c *BroadcastController) SubscribeForGlobalRoomEvents(w http.ResponseWriter, r *http.Request) {
	userUuid, ok := auth.GetUserUuid(r.Context())

	if !ok {
		http.Error(w, entity.UserNotFoundError.Error(), http.StatusUnauthorized)
		return
	}

	_ = c.broadcaster.Subscribe(broadcast.GlobalRoomsEventChannel, userUuid, w, r)
}

func (c *BroadcastController) SubscribeForRoomEvents(w http.ResponseWriter, r *http.Request) {
	roomUuid := r.PathValue("room")

	if roomUuid == "" {
		http.Error(w, entity.RoomNotFoundError.Error(), http.StatusNotFound)
		return
	}

	userUuid, ok := auth.GetUserUuid(r.Context())

	if !ok {
		http.Error(w, entity.UserNotFoundError.Error(), http.StatusUnauthorized)
		return
	}

	err := c.validateUserSubscriptionUseCase.Execute(roomUuid, userUuid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	channel := strings.Replace(broadcast.RoomEventsChannel, "{room}", roomUuid, 1)

	_ = c.broadcaster.Subscribe(channel, userUuid, w, r, &broadcast.Callbacker{
		OnDisconnect: func(_ string, subscriberKey string) {
			_ = c.handleUserDisconnectUseCase.Execute(subscriberKey)
		},
	})
}

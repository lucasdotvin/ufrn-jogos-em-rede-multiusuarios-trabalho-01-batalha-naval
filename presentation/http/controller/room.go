package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	roomuse "trabalho-01-batalha-naval/application/use/room"
	"trabalho-01-batalha-naval/domain/entity"
	"trabalho-01-batalha-naval/presentation/auth"
	roomrequest "trabalho-01-batalha-naval/presentation/http/request/room"
	"trabalho-01-batalha-naval/presentation/http/response/room"
	"trabalho-01-batalha-naval/presentation/http/response/roommove"
)

type RoomController struct {
	listRoomUseCase               *roomuse.ListRoomUseCase
	findUserActiveRoom            *roomuse.FindUserActiveRoomUseCase
	createRoomUseCase             *roomuse.CreateRoomUseCase
	ingressUserUseCase            *roomuse.IngressUserUseCase
	registerShipPlacementsUseCase *roomuse.RegisterShipPlacementsUseCase
	registerFireUseCase           *roomuse.RegisterFireUseCase
}

func NewRoomController(
	listUseCase *roomuse.ListRoomUseCase,
	findUserActiveRoom *roomuse.FindUserActiveRoomUseCase,
	createUseCase *roomuse.CreateRoomUseCase,
	ingressUserUseCase *roomuse.IngressUserUseCase,
	registerShipPlacementsUseCase *roomuse.RegisterShipPlacementsUseCase,
	registerFireUseCase *roomuse.RegisterFireUseCase,
) *RoomController {
	return &RoomController{
		listUseCase,
		findUserActiveRoom,
		createUseCase,
		ingressUserUseCase,
		registerShipPlacementsUseCase,
		registerFireUseCase,
	}
}

func (c *RoomController) Index(w http.ResponseWriter, _ *http.Request) {
	rs, us, err := c.listRoomUseCase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	parsedResponse := room.NewRoomsResponse(rs, us)

	w.Header().Set("Content-Type", JsonContentType)
	_ = json.NewEncoder(w).Encode(parsedResponse)
}

func (c *RoomController) FindMyActiveRoom(w http.ResponseWriter, r *http.Request) {
	userUuid, ok := auth.GetUserUuid(r.Context())

	if !ok {
		http.Error(w, entity.UserNotFoundError.Error(), http.StatusUnauthorized)
		return
	}

	ro, cr, uss, rus, err := c.findUserActiveRoom.Execute(userUuid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parsedResponse := room.NewRoomResponse(ro, cr, uss, rus)

	w.Header().Set("Content-Type", JsonContentType)
	_ = json.NewEncoder(w).Encode(parsedResponse)
}

func (c *RoomController) Store(w http.ResponseWriter, r *http.Request) {
	var input roomrequest.StoreRequest
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userUuid, ok := auth.GetUserUuid(r.Context())

	if !ok {
		http.Error(w, entity.UserNotFoundError.Error(), http.StatusUnauthorized)
		return
	}

	ro, us, err := c.createRoomUseCase.Execute(input.Name, userUuid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parsedResponse := room.NewRoomResponse(ro, us, nil, nil)

	w.Header().Set("Content-Type", JsonContentType)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(parsedResponse)
}

func (c *RoomController) IngressUser(w http.ResponseWriter, r *http.Request) {
	roomUuid := r.PathValue("room")

	if roomUuid == "" {
		http.Error(w, entity.RoomNotFoundError.Error(), http.StatusBadRequest)
		return
	}

	userUuid, ok := auth.GetUserUuid(r.Context())

	if !ok {
		http.Error(w, entity.UserNotFoundError.Error(), http.StatusUnauthorized)
		return
	}

	ro, uss, rus, err := c.ingressUserUseCase.Execute(roomUuid, userUuid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parsedResponse := room.NewRoomResponse(ro, nil, uss, rus)

	w.Header().Set("Content-Type", JsonContentType)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(parsedResponse)
}

func (c *RoomController) RegisterShipPlacements(w http.ResponseWriter, r *http.Request) {
	roomUuid := r.PathValue("room")

	if roomUuid == "" {
		http.Error(w, entity.RoomNotFoundError.Error(), http.StatusBadRequest)
		return
	}

	userUuid, ok := auth.GetUserUuid(r.Context())

	if !ok {
		http.Error(w, entity.UserNotFoundError.Error(), http.StatusUnauthorized)
		return
	}

	var input roomrequest.RegisterPlacementsRequest
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	dispositions, err := input.ToDispositions()

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		fmt.Println(err)
		return
	}

	err = c.registerShipPlacementsUseCase.Execute(roomUuid, userUuid, dispositions)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *RoomController) RegisterFire(w http.ResponseWriter, r *http.Request) {
	roomUuid := r.PathValue("room")

	if roomUuid == "" {
		http.Error(w, entity.RoomNotFoundError.Error(), http.StatusBadRequest)
		return
	}

	userUuid, ok := auth.GetUserUuid(r.Context())

	if !ok {
		http.Error(w, entity.UserNotFoundError.Error(), http.StatusUnauthorized)
		return
	}

	var input roomrequest.RegisterFireRequest
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	position := input.ToPosition()
	roomMove, err := c.registerFireUseCase.Execute(roomUuid, userUuid, position)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	parsedResponse := roommove.NewRoomMoveResponse(roomMove)

	w.Header().Set("Content-Type", JsonContentType)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(parsedResponse)
}

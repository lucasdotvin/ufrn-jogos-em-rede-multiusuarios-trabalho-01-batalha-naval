package room

import "trabalho-01-batalha-naval/domain/entity"

type RegisterFireRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (r *RegisterFireRequest) ToPosition() *entity.Position {
	return &entity.Position{
		X: r.X,
		Y: r.Y,
	}
}

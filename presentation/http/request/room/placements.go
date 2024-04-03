package room

import "trabalho-01-batalha-naval/domain/entity"

type PositionField struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlacementField struct {
	Ship      string         `json:"ship"`
	Position  *PositionField `json:"position"`
	Direction string         `json:"direction"`
}

type RegisterPlacementsRequest struct {
	Placements []*PlacementField `json:"placements"`
}

func (r *RegisterPlacementsRequest) ToDispositions() ([]*entity.Disposition, error) {
	dispositions := make([]*entity.Disposition, 0, len(r.Placements))

	for _, placement := range r.Placements {
		ship, err := entity.ParseShip(placement.Ship)

		if err != nil {
			return nil, err
		}

		orientation, err := entity.ParseOrientation(placement.Direction)

		if err != nil {
			return nil, err
		}

		position := &entity.Position{
			X: placement.Position.X,
			Y: placement.Position.Y,
		}

		dispositions = append(dispositions, &entity.Disposition{
			Ship:        ship,
			Position:    position,
			Orientation: orientation,
		})
	}

	return dispositions, nil
}

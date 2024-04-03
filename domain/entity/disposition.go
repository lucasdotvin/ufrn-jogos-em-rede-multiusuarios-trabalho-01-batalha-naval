package entity

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Disposition struct {
	Ship        Ship        `json:"ship"`
	Position    *Position   `json:"position"`
	Orientation Orientation `json:"orientation"`
}

type Schema struct {
	Dispositions []*Disposition `json:"dispositions"`
	collisionMap map[string]bool
}

var (
	AlreadyOccupiedError = errors.New("position already occupied")
)

func NewSchema(dispositions []*Disposition) *Schema {
	return &Schema{
		Dispositions: dispositions,
	}
}

func (s *Schema) buildCollisionMap() error {
	if s.collisionMap != nil {
		return nil
	}

	fmt.Println("building collision map")

	p, _ := json.MarshalIndent(s.Dispositions, "", "  ")
	fmt.Println("dispositions: \n", string(p))

	s.collisionMap = make(map[string]bool)

	for _, d := range s.Dispositions {
		fromX := d.Position.X
		fromY := d.Position.Y

		toX := fromX
		toY := fromY

		if d.Orientation == OrientationHorizontal {
			toX += d.Ship.Size()
			toY += 1
		} else {
			toX += 1
			toY += d.Ship.Size()
		}

		for x := fromX; x < toX; x++ {
			for y := fromY; y < toY; y++ {
				key := s.makeMapKey(x, y)

				if s.collisionMap[key] {
					return AlreadyOccupiedError
				}

				s.collisionMap[key] = true
			}
		}
	}

	p, _ = json.MarshalIndent(s.collisionMap, "", "  ")
	fmt.Println("collision map: \n", string(p))

	return nil
}

func (s *Schema) HasCollisionError() bool {
	collisionError := s.buildCollisionMap()

	return collisionError != nil
}

func (s *Schema) HasValidShipAmount() bool {
	shipAmounts := make(map[Ship]int)

	for _, d := range s.Dispositions {
		shipAmounts[d.Ship]++
	}

	for ship, receivedAmount := range shipAmounts {
		expectedAmount := AmountByShip[ship]

		if receivedAmount != expectedAmount {
			return false
		}
	}

	return true
}

func (s *Schema) FitsLimits(height int, width int) bool {
	for _, d := range s.Dispositions {
		if d.Position.X < 0 || d.Position.Y < 0 {
			return false
		}

		if d.Orientation == OrientationHorizontal {
			if d.Position.X+d.Ship.Size() > width {
				return false
			}
		} else {
			if d.Position.Y+d.Ship.Size() > height {
				return false
			}
		}
	}

	return true
}

func (s *Schema) IsValid(heightLimit int, widthLimit int) bool {
	return !s.HasCollisionError() && s.HasValidShipAmount() && s.FitsLimits(heightLimit, widthLimit)
}

func (s *Schema) Hits(p *Position) bool {
	if s.collisionMap == nil {
		err := s.buildCollisionMap()

		if err != nil {
			fmt.Println("error building collision map: ", err)
			return false
		}
	}

	return s.collisionMap[s.makeMapKey(p.X, p.Y)]
}

func (s *Schema) makeMapKey(x int, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

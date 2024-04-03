package entity

import "errors"

type Ship string

const (
	ShipDestroyer  Ship = "destroyer"
	ShipBattleship Ship = "battleship"
	ShipCruiser    Ship = "cruiser"
	ShipSubmarine  Ship = "submarine"
)

var (
	UnknownShipError = errors.New("unknown ship")
)

var (
	AmountByShip = map[Ship]int{
		ShipDestroyer:  2,
		ShipBattleship: 1,
		ShipCruiser:    1,
		ShipSubmarine:  3,
	}
	SizeByShip = map[Ship]int{
		ShipDestroyer:  2,
		ShipBattleship: 4,
		ShipCruiser:    3,
		ShipSubmarine:  1,
	}
	TotalShipsSize = 2*2 + 4 + 3 + 3*1
)

func ParseShip(s string) (Ship, error) {
	switch s {
	case "destroyer":
		return ShipDestroyer, nil
	case "battleship":
		return ShipBattleship, nil
	case "cruiser":
		return ShipCruiser, nil
	case "submarine":
		return ShipSubmarine, nil
	default:
		return "", UnknownShipError
	}
}

func (s Ship) Amount() int {
	return AmountByShip[s]
}

func (s Ship) Size() int {
	return SizeByShip[s]
}

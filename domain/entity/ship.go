package entity

type Ship string

const (
	ShipCarrier    Ship = "carrier"
	ShipBattleship Ship = "battleship"
	ShipCruiser    Ship = "cruiser"
	ShipSubmarine  Ship = "submarine"
)

type ShipOrientation string

const (
	ShipOrientationHorizontal ShipOrientation = "horizontal"
	ShipOrientationVertical   ShipOrientation = "vertical"
)

type ShipPosition struct {
	X int
	Y int
}

type ShipPlacement struct {
	Ship        Ship            `json:"ship"`
	Orientation ShipOrientation `json:"orientation"`
	Position    ShipPosition    `json:"position"`
}

var (
	ShipAmounts = map[Ship]int{
		ShipCarrier:    2,
		ShipBattleship: 2,
		ShipCruiser:    1,
		ShipSubmarine:  3,
	}
	ShipSizes = map[Ship]int{
		ShipCarrier:    4,
		ShipBattleship: 3,
		ShipCruiser:    2,
		ShipSubmarine:  1,
	}
)

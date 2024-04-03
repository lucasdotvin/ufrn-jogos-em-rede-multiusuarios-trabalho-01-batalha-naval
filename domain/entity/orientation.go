package entity

import "errors"

type Orientation string

const (
	OrientationHorizontal Orientation = "horizontal"
	OrientationVertical   Orientation = "vertical"
)

var (
	UnknownOrientationError = errors.New("unknown orientation")
)

func ParseOrientation(s string) (Orientation, error) {
	switch s {
	case "horizontal":
		return OrientationHorizontal, nil
	case "vertical":
		return OrientationVertical, nil
	default:
		return "", UnknownOrientationError
	}
}

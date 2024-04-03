package config

import (
	"trabalho-01-batalha-naval/config"
	"trabalho-01-batalha-naval/domain/entity"
)

type ShipsConfig struct {
	Size   int         `json:"size"`
	Amount int         `json:"amount"`
	ID     entity.Ship `json:"id"`
}

type MapConfig struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type ListConfigResult struct {
	Ships []*ShipsConfig `json:"ships"`
	Map   *MapConfig     `json:"map"`
}

type ListConfigUseCase struct {
	cfg config.Config
}

func NewListConfigUseCase(cfg config.Config) *ListConfigUseCase {
	return &ListConfigUseCase{
		cfg,
	}
}

func (uc *ListConfigUseCase) Execute() *ListConfigResult {
	result := &ListConfigResult{}

	allShips := []entity.Ship{
		entity.ShipDestroyer,
		entity.ShipBattleship,
		entity.ShipCruiser,
		entity.ShipSubmarine,
	}

	result.Ships = make([]*ShipsConfig, 0, len(allShips))

	for _, ship := range allShips {
		entry := &ShipsConfig{
			Size:   ship.Size(),
			Amount: ship.Amount(),
			ID:     ship,
		}

		result.Ships = append(result.Ships, entry)
	}

	result.Map = &MapConfig{
		Height: uc.cfg.MapHeight,
		Width:  uc.cfg.MapWidth,
	}

	return result
}
